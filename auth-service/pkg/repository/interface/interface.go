package interfaces

import (
	"Auth/pkg/domain"
	"Auth/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
}

type EmployerRepository interface {
	EmployerSignUp(employerDetails models.EmployerSignUp) (models.EmployerDetailsResponse, error)
	FindEmployerByEmail(employer models.EmployerLogin) (models.EmployerSignUp, error)
	CheckEmployerExistsByEmail(email string) (*domain.Employer, error)
	GetCompanyDetails(employerID int32) (models.EmployerDetailsResponse, error)
	UpdateCompany(employerIDInt int32, employerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error)
}

type JobSeekerRepository interface {
	JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.JobSeekerDetailsResponse, error)
	CheckJobSeekerExistsByEmail(email string) (*domain.JobSeeker, error)
	FindJobSeekerByEmail(jobSeeker models.JobSeekerLogin) (models.JobSeekerSignUp, error)
}
