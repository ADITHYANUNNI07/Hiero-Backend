package db

import (
	"Auth/pkg/config"
	"Auth/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", "localhost", "postgres", "hireo_job", "5432", "12345")
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.AutoMigrate(&domain.JobOpening{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.JobOpeningResponse{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.ApplyJob{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.ApplyJobResponse{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.SavedJobs{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.Interview{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.InterviewResponse{}); err != nil {
		return nil, err
	}

	return db, nil
}
