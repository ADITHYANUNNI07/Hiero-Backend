// job_seeker_service.go
package service

import (
	pb "Auth/pkg/pb/auth"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"fmt"
)

type JobSeekerServer struct {
	jobSeekerUseCase interfaces.JobSeekerUseCase
	pb.UnimplementedJobSeekerServer
}

func NewJobSeekerServer(useCase interfaces.JobSeekerUseCase) pb.JobSeekerServer {
	return &JobSeekerServer{
		jobSeekerUseCase: useCase,
	}
}

func (js *JobSeekerServer) JobSeekerSignup(ctx context.Context, req *pb.JobSeekerSignupRequest) (*pb.JobSeekerSignupResponse, error) {
	jobSeekerSignup := models.JobSeekerSignUp{
		Email:       req.Email,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Address:     req.Address,
		Bio:         req.Bio,
	}

	fmt.Println("service", jobSeekerSignup)

	res, err := js.jobSeekerUseCase.JobSeekerSignUp(jobSeekerSignup)
	if err != nil {
		return &pb.JobSeekerSignupResponse{}, err
	}

	jobSeekerDetails := &pb.JobSeekerDetails{
		Id:          uint64(res.JobSeeker.ID),
		Email:       res.JobSeeker.Email,
		FirstName:   res.JobSeeker.FirstName,
		LastName:    res.JobSeeker.LastName,
		PhoneNumber: res.JobSeeker.PhoneNumber,
		DateOfBirth: res.JobSeeker.DateOfBirth,
		Gender:      res.JobSeeker.Gender,
		Address:     res.JobSeeker.Address,
		Bio:         res.JobSeeker.Bio,
	}

	return &pb.JobSeekerSignupResponse{
		Status:           201,
		JobSeekerDetails: jobSeekerDetails,
		Token:            res.Token,
	}, nil
}

func (js *JobSeekerServer) JobSeekerLogin(ctx context.Context, req *pb.JobSeekerLoginRequest) (*pb.JobSeekerLoginResponse, error) {
	jobSeekerLogin := models.JobSeekerLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	jobSeeker, err := js.jobSeekerUseCase.JobSeekerLogin(jobSeekerLogin)
	if err != nil {
		return &pb.JobSeekerLoginResponse{}, err
	}

	jobSeekerDetails := &pb.JobSeekerDetails{
		Id:          uint64(jobSeeker.JobSeeker.ID),
		Email:       jobSeeker.JobSeeker.Email,
		FirstName:   jobSeeker.JobSeeker.FirstName,
		LastName:    jobSeeker.JobSeeker.LastName,
		PhoneNumber: jobSeeker.JobSeeker.PhoneNumber,
		DateOfBirth: jobSeeker.JobSeeker.DateOfBirth,
		Gender:      jobSeeker.JobSeeker.Gender,
		Address:     jobSeeker.JobSeeker.Address,
		Bio:         jobSeeker.JobSeeker.Bio,
	}

	return &pb.JobSeekerLoginResponse{
		Status:           200,
		JobSeekerDetails: jobSeekerDetails,
		Token:            jobSeeker.Token,
	}, nil
}
