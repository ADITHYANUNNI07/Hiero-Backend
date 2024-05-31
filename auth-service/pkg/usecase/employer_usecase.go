// employer_usecase.go
package usecase

import (
	"Auth/pkg/domain"
	"Auth/pkg/helper"
	interfaces "Auth/pkg/repository/interface"
	services "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type employerUseCase struct {
	employerRepository interfaces.EmployerRepository
}

func NewEmployerUseCase(repository interfaces.EmployerRepository) services.EmployerUseCase {
	return &employerUseCase{
		employerRepository: repository,
	}
}

func (eu *employerUseCase) EmployerSignUp(employer models.EmployerSignUp) (*domain.TokenEmployer, error) {
	email, err := eu.employerRepository.CheckEmployerExistsByEmail(employer.ContactEmail)
	fmt.Println(email)
	if err != nil {
		return &domain.TokenEmployer{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenEmployer{}, errors.New("employer with this email already exists")
	}
	hashPassword, err := helper.PasswordHash(employer.Password)
	if err != nil {
		return &domain.TokenEmployer{}, errors.New("error in hashing password")
	}
	employer.Password = hashPassword
	employerData, err := eu.employerRepository.EmployerSignUp(employer)
	if err != nil {
		return &domain.TokenEmployer{}, errors.New("could not add the employer")
	}
	tokenString, err := helper.GenerateTokenEmployer(employerData)
	if err != nil {
		return &domain.TokenEmployer{}, err
	}

	return &domain.TokenEmployer{
		Employer: employerData,
		Token:    tokenString,
	}, nil
}

func (eu *employerUseCase) EmployerLogin(employer models.EmployerLogin) (*domain.TokenEmployer, error) {
	email, err := eu.employerRepository.CheckEmployerExistsByEmail(employer.Email)
	if err != nil {
		return &domain.TokenEmployer{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenEmployer{}, errors.New("email doesn't exist")
	}
	employerDetails, err := eu.employerRepository.FindEmployerByEmail(employer)
	if err != nil {
		return &domain.TokenEmployer{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(employerDetails.Password), []byte(employer.Password))
	if err != nil {
		return &domain.TokenEmployer{}, errors.New("password not matching")
	}
	var employerDetailsResponse models.EmployerDetailsResponse

	err = copier.Copy(&employerDetailsResponse, &employerDetails)
	if err != nil {
		return &domain.TokenEmployer{}, err
	}

	tokenString, err := helper.GenerateTokenEmployer(employerDetailsResponse)
	if err != nil {
		return &domain.TokenEmployer{}, err
	}

	return &domain.TokenEmployer{
		Employer: employerDetailsResponse,
		Token:    tokenString,
	}, nil
}

func (eh *employerUseCase) GetCompanyDetails(employerID int32) (models.EmployerDetailsResponse, error) {
	// Call the repository method to fetch employer details by ID
	employerDetails, err := eh.employerRepository.GetCompanyDetails(employerID)
	if err != nil {
		return models.EmployerDetailsResponse{}, fmt.Errorf("failed to get company details: %v", err)
	}

	return employerDetails, nil
}

func (eh *employerUseCase) UpdateCompany(employerIDInt int32, employerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error) {
	updatedEmployerDetails, err := eh.employerRepository.UpdateCompany(employerIDInt, employerDetails)
	if err != nil {
		return models.EmployerDetailsResponse{}, fmt.Errorf("failed to update company details: %v", err)
	}
	return updatedEmployerDetails, nil
}
