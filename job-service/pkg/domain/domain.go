package domain

import "time"

type JobOpening struct {
	ID                  uint       `json:"id"`
	Title               string     `json:"title"`
	Type                string     `json:"type"`
	Description         string     `json:"description"`
	Requirements        string     `json:"requirements"`
	Location            string     `json:"location"`
	EmploymentType      string     `json:"employment_type"`
	Salary              string     `json:"salary"`
	SkillsRequired      string     `json:"skills_required"`
	StartingDate        time.Time  `json:"starting_date"`
	ExperienceLevel     string     `json:"experience_level"`
	EducationLevel      string     `json:"education_level"`
	ApplicationDeadline time.Time  `json:"application_deadline"`
	CompanyName         string     `json:"company_name"`
	YearOfExperience    string     `json:"year_of_experience"`
	CandidatesHired     string     `json:"candidates_hired"`
	Opportunities       string     `json:"opportunities"`
	UpdatedOn           *time.Time `json:"updated_on"`
	IsDeleted           bool       `json:"is_deleted"`
}

type JobOpeningResponse struct {
	ID                  uint       `json:"id"`
	Title               string     `json:"title"`
	Type                string     `json:"type"`
	Description         string     `json:"description"`
	Requirements        string     `json:"requirements"`
	PostedOn            time.Time  `json:"posted_on"`
	EmployerID          int32      `json:"employer_id"`
	Location            string     `json:"location"`
	EmploymentType      string     `json:"employment_type"`
	Salary              string     `json:"salary"`
	SkillsRequired      string     `json:"skills_required"`
	ExperienceLevel     string     `json:"experience_level"`
	EducationLevel      string     `json:"education_level"`
	ApplicationDeadline time.Time  `json:"application_deadline"`
	CompanyName         string     `json:"company_name"`
	YearOfExperience    string     `json:"year_of_experience"`
	Opportunities       string     `json:"opportunities"`
	CandidatesHired     string     `json:"candidates_hired"`
	StartingDate        time.Time  `json:"starting_date"`
	UpdatedOn           *time.Time `json:"updated_on"`
	IsDeleted           bool       `json:"is_deleted"`
}

type ApplyJob struct {
	ID          uint   `json:"id"`
	JobseekerID int64  `json:"jobseeker_id" validate:"required"`
	JobID       int64  `json:"job_id" validate:"required"`
	Resume      []byte `json:"resume" validate:"required"`
	ResumeURL   string `json:"resume_url" validate:"required"`
	CoverLetter string `json:"cover_letter" validate:"lte=500"`
}

type ApplyJobResponse struct {
	ID          uint   `json:"id"`
	JobseekerID int64  `json:"jobseeker_id" validate:"required"`
	JobID       int64  `json:"job_id" validate:"required"`
	ResumeURL   string `json:"resume_url" validate:"required"`
	CoverLetter string `json:"cover_letter" validate:"lte=500"`
}

type SavedJobs struct {
	ID          uint  `json:"id"`
	JobID       int64 `json:"job_id" validate:"required"`
	JobseekerID int64 `json:"jobseeker_id" validate:"required"`
}

type Interview struct {
	ID            uint      `json:"id"`
	JobID         int64     `json:"job_id" validate:"required"`
	JobseekerID   int64     `json:"jobseeker_id" validate:"required"`
	EmployerID    int64     `json:"employer_id" validate:"required"`
	ScheduledTime time.Time `json:"scheduled_time" validate:"required"`
	Mode          string    `json:"mode" validate:"oneof=ONLINE OFFLINE" default:"ONLINE"`
	Link          string    `json:"link,omitempty"`
	Status        string    `json:"status" validate:"oneof=SCHEDULED COMPLETED CANCELLED" default:"SCHEDULED"`
}

type InterviewResponse struct {
	ID            uint      `json:"id"`
	JobID         int64     `json:"job_id" validate:"required"`
	JobseekerID   int64     `json:"jobseeker_id" validate:"required"`
	EmployerID    int32     `json:"employer_id" validate:"required"`
	ScheduledTime time.Time `json:"scheduled_time" validate:"required"`
	Mode          string    `json:"mode" validate:"oneof=ONLINE OFFLINE" default:"ONLINE"`
	Link          string    `json:"link,omitempty"`
	Status        string    `json:"status" validate:"oneof=SCHEDULED COMPLETED CANCELLED" default:"SCHEDULED"`
}
