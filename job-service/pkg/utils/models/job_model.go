package models

import (
	"time"
)

type JobOpening struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Type                string    `json:"type"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	StartingDate        time.Time `json:"starting_date"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	CompanyName         string    `json:"company_name"`
	YearOfExperience    string    `json:"year_of_experience"`
	CandidatesHired     string    `json:"candidates_hired"`
	Opportunities       string    `json:"opportunities"`
}

type JobOpeningResponse struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Type                string    `json:"type"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerID          int32     `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	CompanyName         string    `json:"company_name"`
	YearOfExperience    string    `json:"year_of_experience"`
	Opportunities       string    `json:"opportunities"`
	CandidatesHired     string    `json:"candidates_hired"`
	StartingDate        time.Time `json:"starting_date"`
}

type AllJob struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Type                string    `json:"type"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerID          int32     `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      []string  `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	CompanyName         string    `json:"company_name"`
	YearOfExperience    string    `json:"year_of_experience"`
	Opportunities       string    `json:"opportunities"`
	CandidatesHired     string    `json:"candidates_hired"`
	StartingDate        time.Time `json:"starting_date"`
}

type JobSeekerGetAllJobs struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
