package models

import "time"

type JobSeekerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type JobSeekerSignUp struct {
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type JobSeekerDetailsResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Bio         string `json:"bio"`
}

type TokenJobSeeker struct {
	JobSeeker JobSeekerDetailsResponse
	Token     string
}

type ApplyJob struct {
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
	JobID       int64 `json:"job_id" validate:"required"`
	JobseekerID int64 `json:"jobseeker_id" validate:"required"`
}

type SavedJobsResponse struct {
	ID          uint  `json:"id"`
	JobID       int64 `json:"job_id" validate:"required"`
	JobseekerID int64 `json:"jobseeker_id" validate:"required"`
}

type Interview struct {
	ID            uint      `json:"id"`
	JobID         int64     `json:"job_id" validate:"required"`
	JobseekerID   int64     `json:"jobseeker_id" validate:"required"`
	EmployerID    int32     `json:"employer_id" validate:"required"`
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
