package interfaces

import (
	"Auth/pkg/domain"
	"Auth/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}

type EmployerUseCase interface {
	EmployerSignUp(employerDetails models.EmployerSignUp) (*domain.TokenEmployer, error)
	EmployerLogin(employerDetails models.EmployerLogin) (*domain.TokenEmployer, error)
	GetCompanyDetails(employerID int32) (models.EmployerDetailsResponse, error)
	UpdateCompany(employerIDInt int32, employerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error)
}

type JobSeekerUseCase interface {
	JobSeekerLogin(jobSeeker models.JobSeekerLogin) (*domain.TokenJobSeeker, error)
	JobSeekerSignUp(jobSeeker models.JobSeekerSignUp) (*domain.TokenJobSeeker, error)
}
