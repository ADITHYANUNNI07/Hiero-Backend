package usecase

import (
	"Auth/pkg/domain"
	"Auth/pkg/helper"
	interfaces "Auth/pkg/repository/interface"
	services "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"errors"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type jobSeekerUseCase struct {
	jobSeekerRepository interfaces.JobSeekerRepository
}

func NewJobSeekerUseCase(repository interfaces.JobSeekerRepository) services.JobSeekerUseCase {
	return &jobSeekerUseCase{
		jobSeekerRepository: repository,
	}
}

func (jsu *jobSeekerUseCase) JobSeekerSignUp(jobSeeker models.JobSeekerSignUp) (*domain.TokenJobSeeker, error) {
	email, err := jsu.jobSeekerRepository.CheckJobSeekerExistsByEmail(jobSeeker.Email)
	if err != nil {
		return &domain.TokenJobSeeker{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenJobSeeker{}, errors.New("job seeker with this email already exists")
	}
	hashPassword, err := helper.PasswordHash(jobSeeker.Password)
	if err != nil {
		return &domain.TokenJobSeeker{}, errors.New("error in hashing password")
	}
	jobSeeker.Password = hashPassword
	jobSeekerData, err := jsu.jobSeekerRepository.JobSeekerSignUp(jobSeeker)
	if err != nil {
		return &domain.TokenJobSeeker{}, errors.New("could not add the job seeker")
	}
	tokenString, err := helper.GenerateTokenJobSeeker(jobSeekerData)
	if err != nil {
		return &domain.TokenJobSeeker{}, err
	}

	return &domain.TokenJobSeeker{
		JobSeeker: jobSeekerData,
		Token:     tokenString,
	}, nil
}

func (jsu *jobSeekerUseCase) JobSeekerLogin(jobSeeker models.JobSeekerLogin) (*domain.TokenJobSeeker, error) {
	email, err := jsu.jobSeekerRepository.CheckJobSeekerExistsByEmail(jobSeeker.Email)
	if err != nil {
		return &domain.TokenJobSeeker{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenJobSeeker{}, errors.New("email doesn't exist")
	}
	jobSeekerDetails, err := jsu.jobSeekerRepository.FindJobSeekerByEmail(jobSeeker)
	if err != nil {
		return &domain.TokenJobSeeker{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(jobSeekerDetails.Password), []byte(jobSeeker.Password))
	if err != nil {
		return &domain.TokenJobSeeker{}, errors.New("password not matching")
	}
	var jobSeekerDetailsResponse models.JobSeekerDetailsResponse

	err = copier.Copy(&jobSeekerDetailsResponse, &jobSeekerDetails)
	if err != nil {
		return &domain.TokenJobSeeker{}, err
	}

	tokenString, err := helper.GenerateTokenJobSeeker(jobSeekerDetailsResponse)
	if err != nil {
		return &domain.TokenJobSeeker{}, err
	}

	return &domain.TokenJobSeeker{
		JobSeeker: jobSeekerDetailsResponse,
		Token:     tokenString,
	}, nil
}
