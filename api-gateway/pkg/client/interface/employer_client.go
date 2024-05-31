package interfaces

import "HireoGateWay/pkg/utils/models"

type EmployerClient interface {
	EmployerSignUp(employerDetails models.EmployerSignUp) (models.TokenEmployer, error)
	EmployerLogin(employerDetails models.EmployerLogin) (models.TokenEmployer, error)
	GetCompanyDetails(employerIDInt int32) (models.EmployerDetailsResponse, error)
	UpdateCompany(employerIDInt int32, EmployerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error)
}
