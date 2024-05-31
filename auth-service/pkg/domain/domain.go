package domain

import "Auth/pkg/utils/models"

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type Employer struct {
	ID                   uint   `json:"id" gorm:"uniquekey; not null"`
	Company_name         string `json:"company_name" gorm:"validate:required"`
	Industry             string `json:"industry" gorm:"validate:required"`
	Company_size         int    `json:"company_size" gorm:"validate:required"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company" gorm:"type:text"`
	Contact_email        string `json:"contact_email" gorm:"validate:required"`
	Contact_phone_number uint   `json:"contact_phone_number" gorm:"type:numeric"`
	Password             string `json:"password" gorm:"validate:required"`
}

type TokenEmployer struct {
	Employer models.EmployerDetailsResponse
	Token    string
}

type JobSeeker struct {
	ID            uint   `json:"id" gorm:"uniquekey; not null"`
	Email         string `json:"email" gorm:"validate:required"`
	Password      string `json:"password" gorm:"validate:required"`
	First_name    string `json:"first_name" gorm:"validate:required"`
	Last_name     string `json:"last_name" gorm:"validate:required"`
	Phone_number  string `json:"phone_number" gorm:"validate:required"`
	Date_of_birth string `json:"date_of_birth" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted_at    string `json:"deleted_at"`
}

type TokenJobSeeker struct {
	JobSeeker models.JobSeekerDetailsResponse
	Token     string
}
