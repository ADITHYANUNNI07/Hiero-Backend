package repository

import (
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(DB *gorm.DB) interfaces.JobRepository {
	return &jobRepository{
		DB: DB,
	}
}

func (jr *jobRepository) PostJob(jobDetails models.JobOpening, employerID int32) (models.JobOpeningResponse, error) {
	postedOn := time.Now()

	job := models.JobOpeningResponse{
		Title:               jobDetails.Title,
		Type:                jobDetails.Type,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          employerID,
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
		CompanyName:         jobDetails.CompanyName,
		YearOfExperience:    jobDetails.YearOfExperience,
		Opportunities:       jobDetails.Opportunities,
		CandidatesHired:     jobDetails.CandidatesHired,
		StartingDate:        jobDetails.StartingDate,
	}

	if err := jr.DB.Create(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) GetAllJobs(employerID int32) ([]models.AllJob, error) {
	var jobs []models.AllJob

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("employer_id = ?", employerID).Select("id, title, application_deadline, employer_id").Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (jr *jobRepository) GetAJob(employerID, jobId int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ? AND employer_id = ?", jobId, employerID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) IsJobExist(jobID int32) (bool, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (jr *jobRepository) GetJobIDByEmployerID(employerID int64) (int64, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("employer_id = ?", employerID).Scan(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}

	return int64(job.ID), nil
}
func (jr *jobRepository) GetApplicantsByEmployerID(jobID int64) ([]models.ApplyJobResponse, error) {
	var applicants []models.ApplyJobResponse

	fmt.Println("job id", jobID)

	query := "SELECT * FROM apply_jobs WHERE job_id = ?"
	if err := jr.DB.Raw(query, jobID).Scan(&applicants).Error; err != nil {
		return nil, err
	}

	fmt.Println("Retrieved applicants:", applicants)

	return applicants, nil
}

func (jr *jobRepository) DeleteAJob(employerID, jobID int32) error {
	if err := jr.DB.Where("id = ? AND employer_id = ?", jobID, employerID).Delete(&models.JobOpeningResponse{}).Error; err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}

	return nil
}

func (jr *jobRepository) UpdateAJob(employerID int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error) {
	postedOn := time.Now()

	updatedJob := models.JobOpeningResponse{
		ID:                  uint(jobID),
		Title:               jobDetails.Title,
		Description:         jobDetails.Description,
		Requirements:        jobDetails.Requirements,
		PostedOn:            postedOn,
		EmployerID:          int32(employerID),
		Location:            jobDetails.Location,
		EmploymentType:      jobDetails.EmploymentType,
		Salary:              jobDetails.Salary,
		SkillsRequired:      jobDetails.SkillsRequired,
		ExperienceLevel:     jobDetails.ExperienceLevel,
		EducationLevel:      jobDetails.EducationLevel,
		ApplicationDeadline: jobDetails.ApplicationDeadline,
	}

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ? AND employer_id = ?", jobID, employerID).Updates(updatedJob).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return updatedJob, nil
}

func (jr *jobRepository) JobSeekerGetAllJobs(keyword string) ([]models.JobOpeningResponse, error) {
	var jobSeekerJobs []models.JobOpeningResponse

	if err := jr.DB.Where("title ILIKE ?", "%"+keyword+"%").Find(&jobSeekerJobs).Error; err != nil {
		return nil, fmt.Errorf("failed to query jobs: %v", err)
	}

	return jobSeekerJobs, nil
}

func (jr *jobRepository) GetJobDetails(jobID int32) (models.JobOpeningResponse, error) {
	var job models.JobOpeningResponse

	if err := jr.DB.Model(&models.JobOpeningResponse{}).Where("id = ?", jobID).First(&job).Error; err != nil {
		return models.JobOpeningResponse{}, err
	}

	return job, nil
}

func (jr *jobRepository) ApplyJob(application models.ApplyJob, resumeURL string) (models.ApplyJobResponse, error) {
	var jobResponse models.ApplyJobResponse

	// Check if the job has already been applied for by this jobseeker
	var count int64
	err := jr.DB.Model(&models.ApplyJob{}).
		Where("jobseeker_id = ? AND job_id = ?", application.JobseekerID, application.JobID).
		Count(&count).Error
	if err != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("error checking if job is already applied: %w", err)
	}

	if count > 0 {
		return models.ApplyJobResponse{}, errors.New("job already applied")
	}

	// Insert new job application
	result := jr.DB.Exec("INSERT INTO apply_jobs (jobseeker_id, job_id, resume_url, cover_letter) VALUES (?, ?, ?, ?)",
		application.JobseekerID,
		application.JobID,
		resumeURL,
		application.CoverLetter)

	if result.Error != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("error inserting into database: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return models.ApplyJobResponse{}, errors.New("no rows were affected during insert")
	}

	err = jr.DB.Raw("SELECT * FROM apply_jobs WHERE jobseeker_id = ? AND job_id = ?", application.JobseekerID, application.JobID).Scan(&jobResponse).Error
	if err != nil {
		return models.ApplyJobResponse{}, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return jobResponse, nil
}

func (jr *jobRepository) SaveJobs(jobID, userID int64) (models.SavedJobsResponse, error) {

	var savedJobResponse models.SavedJobsResponse

	result := jr.DB.Exec("INSERT INTO saved_jobs (job_id, jobseeker_id) VALUES (?, ?) ", jobID, userID)
	if result.Error != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("error inserting into database: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return models.SavedJobsResponse{}, errors.New("no rows were affected during insert")
	}

	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID).Scan(&savedJobResponse).Error
	if err != nil {
		return models.SavedJobsResponse{}, fmt.Errorf("failed to retrieve saved job: %w", err)
	}

	return savedJobResponse, nil
}

func (jr *jobRepository) IsJobSaved(jobID, userID int32) (bool, error) {
	var savedJob models.SavedJobsResponse
	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID).Scan(&savedJob).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to retrieve saved job: %w", err)
	}
	return true, nil
}

func (jr *jobRepository) DeleteSavedJob(jobID, userID int32) error {
	result := jr.DB.Exec("DELETE FROM saved_jobs WHERE job_id = ? AND jobseeker_id = ?", jobID, userID)
	if result.Error != nil {
		return fmt.Errorf("error deleting saved job: %w", result.Error)
	}
	return nil
}

func (jr *jobRepository) GetSavedJobs(userID int32) ([]models.SavedJobsResponse, error) {
	var savedJobs []models.SavedJobsResponse
	err := jr.DB.Raw("SELECT * FROM saved_jobs WHERE jobseeker_id = ?", userID).Scan(&savedJobs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve saved jobs: %w", err)
	}
	return savedJobs, nil
}

func (jr *jobRepository) ScheduleInterview(interview models.Interview) (models.InterviewResponse, error) {
	var Data models.InterviewResponse
	query := `
		INSERT INTO interviews (job_id, jobseeker_id, employer_id, scheduled_time, mode, link, status)
		VALUES (?, ?, ?, ?, ?, ?, ?) returning id, job_id, jobseeker_id, employer_id, scheduled_time, mode, link, status
	`
	err := jr.DB.Raw(query, interview.JobID, interview.JobseekerID, interview.EmployerID, interview.ScheduledTime, interview.Mode, interview.Link, interview.Status).Scan(&Data).Error
	if err != nil {
		return models.InterviewResponse{}, fmt.Errorf("failed to save interview: %v", err)
	}
	return Data, nil
}

func (jr *jobRepository) GetInterview(jobID, employerID int32) (models.InterviewResponse, error) {
	var Data models.InterviewResponse
	fmt.Println("job,emp", jobID, employerID)
	err := jr.DB.Raw("SELECT * FROM interviews WHERE job_id = ? AND employer_id = ?", jobID, employerID).Scan(&Data).Error
	if err != nil {
		return models.InterviewResponse{}, fmt.Errorf("failed to retrieve interview: %v", err)
	}
	fmt.Println("Data", Data)
	return Data, nil
}
