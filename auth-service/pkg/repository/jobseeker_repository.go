package repository

import (
	"Auth/pkg/domain"
	interfaces "Auth/pkg/repository/interface"
	"Auth/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type jobSeekerRepository struct {
	DB *gorm.DB
}

func NewJobSeekerRepository(DB *gorm.DB) interfaces.JobSeekerRepository {
	return &jobSeekerRepository{
		DB: DB,
	}
}

func (jr *jobSeekerRepository) JobSeekerSignUp(jobSeekerDetails models.JobSeekerSignUp) (models.JobSeekerDetailsResponse, error) {
	var model models.JobSeekerDetailsResponse

	fmt.Println("email", model.Email)

	fmt.Println("models", model)
	if err := jr.DB.Raw("INSERT INTO job_seekers (email, password, first_name, last_name, phone_number, date_of_birth, gender, address, bio) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, email, first_name, last_name, phone_number, date_of_birth, gender, address, bio", jobSeekerDetails.Email, jobSeekerDetails.Password, jobSeekerDetails.FirstName, jobSeekerDetails.LastName, jobSeekerDetails.PhoneNumber, jobSeekerDetails.DateOfBirth, jobSeekerDetails.Gender, jobSeekerDetails.Address, jobSeekerDetails.Bio).Scan(&model).Error; err != nil {
		return models.JobSeekerDetailsResponse{}, err
	}
	fmt.Println("inside", model.Email)
	return model, nil
}

func (jr *jobSeekerRepository) CheckJobSeekerExistsByEmail(email string) (*domain.JobSeeker, error) {
	var jobSeeker domain.JobSeeker
	res := jr.DB.Where(&domain.JobSeeker{Email: email}).First(&jobSeeker)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.JobSeeker{}, res.Error
	}
	return &jobSeeker, nil
}

func (jr *jobSeekerRepository) FindJobSeekerByEmail(jobSeeker models.JobSeekerLogin) (models.JobSeekerSignUp, error) {
	var user models.JobSeekerSignUp
	err := jr.DB.Raw("SELECT * FROM job_seekers WHERE email=? ", jobSeeker.Email).Scan(&user).Error
	if err != nil {
		return models.JobSeekerSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}
