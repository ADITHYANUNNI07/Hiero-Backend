package interfaces

import (
	"HireoGateWay/pkg/utils/models"
	"mime/multipart"
)

type JobClient interface {
	PostJobOpening(jobDetails models.JobOpening, EmployerID int32) (models.JobOpeningResponse, error)
	GetAllJobs(employerIDInt int32) ([]models.AllJob, error)
	GetAJob(employerIDInt, jobId int32) (models.JobOpeningResponse, error)
	DeleteAJob(employerIDInt, jobID int32) error
	JobSeekerGetAllJobs(keyword string) ([]models.JobSeekerGetAllJobs, error)
	GetJobDetails(jobID int32) (models.JobOpeningResponse, error)
	UpdateAJob(employerIDInt int32, jobID int32, jobDetails models.JobOpening) (models.JobOpeningResponse, error)
	ApplyJob(jobApplication models.ApplyJob, file *multipart.FileHeader) (models.ApplyJobResponse, error)
	GetApplicants(employerID int64) ([]models.ApplyJobResponse, error)
	SaveAJob(userIdInt, jobIdInt int32) (models.SavedJobsResponse, error)
	DeleteSavedJob(jobIdInt, userIdInt int32) error
	GetASavedJob(userIdInt int32) ([]models.SavedJobsResponse, error)
	ScheduleInterview(interview models.Interview) (models.InterviewResponse, error)
	GetInterview(jobID, employerID int32) (models.InterviewResponse, error)
	//GetAnApplicant(jobID, employerID, jobseekerIdInt int32) (models.ApplyJobResponse, error)
}
