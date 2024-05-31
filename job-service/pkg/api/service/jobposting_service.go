package service

import (
	pb "Auth/pkg/pb/job"
	interfaces "Auth/pkg/usecase/interface"
	"Auth/pkg/utils/models"
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobServer struct {
	jobUseCase interfaces.JobUseCase
	pb.UnimplementedJobServer
}

func NewJobServer(useCase interfaces.JobUseCase) pb.JobServer {
	return &JobServer{
		jobUseCase: useCase,
	}
}
func (js *JobServer) PostJob(ctx context.Context, req *pb.JobOpeningRequest) (*pb.JobOpeningResponse, error) {

	employerID := int32(req.EmployerId)

	jobDetails := models.JobOpening{
		Title:               req.Title,
		Description:         req.Description,
		Requirements:        req.Requirements,
		Location:            req.Location,
		EmploymentType:      req.EmploymentType,
		Salary:              req.Salary,
		SkillsRequired:      req.SkillsRequired,
		ExperienceLevel:     req.ExperienceLevel,
		EducationLevel:      req.EducationLevel,
		ApplicationDeadline: req.ApplicationDeadline.AsTime(),
	}

	fmt.Println("service", jobDetails)

	res, err := js.jobUseCase.PostJob(jobDetails, employerID)
	if err != nil {
		return nil, err
	}

	jobOpening := &pb.JobOpeningResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          int32(req.EmployerId), // Set the EmployerId field
	}

	return jobOpening, nil
}

func (js *JobServer) GetAllJobs(ctx context.Context, req *pb.GetAllJobsRequest) (*pb.GetAllJobsResponse, error) {
	employerID := int32(req.EmployerIDInt)

	jobs, err := js.jobUseCase.GetAllJobs(employerID)
	if err != nil {
		return nil, err
	}

	var jobResponses []*pb.JobOpeningResponse
	for _, job := range jobs {
		jobResponse := &pb.JobOpeningResponse{
			Id:                  uint64(job.ID),
			Title:               job.Title,
			ApplicationDeadline: timestamppb.New(job.ApplicationDeadline),
			EmployerId:          job.EmployerID,
		}
		jobResponses = append(jobResponses, jobResponse)
	}

	return &pb.GetAllJobsResponse{Jobs: jobResponses}, nil
}

func (js *JobServer) GetAJob(ctx context.Context, req *pb.GetAJobRequest) (*pb.JobOpeningResponse, error) {
	employerID := req.EmployerIDInt
	jobId := req.JobId

	res, err := js.jobUseCase.GetAJob(employerID, jobId)
	if err != nil {
		return nil, err
	}

	jobOpening := &pb.JobOpeningResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          employerID,
	}

	return jobOpening, nil
}

func (js *JobServer) DeleteAJob(ctx context.Context, req *pb.DeleteAJobRequest) (*emptypb.Empty, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

	err := js.jobUseCase.DeleteAJob(employerID, jobID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete job: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (js *JobServer) UpdateAJob(ctx context.Context, req *pb.UpdateAJobRequest) (*pb.UpdateAJobResponse, error) {
	employerID := req.EmployerIDInt
	jobID := req.JobId

	jobDetails := models.JobOpening{
		Title:               req.Title,
		Description:         req.Description,
		Requirements:        req.Requirements,
		Location:            req.Location,
		EmploymentType:      req.EmploymentType,
		Salary:              req.Salary,
		SkillsRequired:      req.SkillsRequired,
		ExperienceLevel:     req.ExperienceLevel,
		EducationLevel:      req.EducationLevel,
		ApplicationDeadline: req.ApplicationDeadline.AsTime(),
	}

	fmt.Println("service", jobDetails)

	res, err := js.jobUseCase.UpdateAJob(employerID, jobID, jobDetails)
	if err != nil {
		return nil, err
	}

	updateResponse := &pb.UpdateAJobResponse{
		Id:                  uint64(res.ID),
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		PostedOn:            timestamppb.New(res.PostedOn),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
		EmployerId:          employerID,
	}

	return updateResponse, nil
}

func (js *JobServer) JobSeekerGetAllJobs(ctx context.Context, req *pb.JobSeekerGetAllJobsRequest) (*pb.JobSeekerGetAllJobsResponse, error) {
	keyword := req.Title

	jobSeekerJobs, err := js.jobUseCase.JobSeekerGetAllJobs(keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get jobs for job seeker: %v", err)
	}

	var jobsResponse []*pb.JSGetAllJobsRespons
	for _, job := range jobSeekerJobs {
		jobResponse := &pb.JSGetAllJobsRespons{
			Id:    uint64(job.ID),
			Title: job.Title,
		}
		jobsResponse = append(jobsResponse, jobResponse)
	}

	response := &pb.JobSeekerGetAllJobsResponse{
		Jobs: jobsResponse,
	}

	return response, nil
}

func (js *JobServer) GetJobDetails(ctx context.Context, req *pb.GetJobDetailsRequest) (*pb.GetJobDetailsResponse, error) {
	jobId := req.JobId

	res, err := js.jobUseCase.GetJobDetails(jobId)
	if err != nil {
		return nil, err
	}

	jobDetailsResponse := &pb.GetJobDetailsResponse{
		Title:               res.Title,
		Description:         res.Description,
		Requirements:        res.Requirements,
		EmployerId:          int32(res.EmployerID),
		Location:            res.Location,
		EmploymentType:      res.EmploymentType,
		Salary:              res.Salary,
		SkillsRequired:      res.SkillsRequired,
		ExperienceLevel:     res.ExperienceLevel,
		EducationLevel:      res.EducationLevel,
		ApplicationDeadline: timestamppb.New(res.ApplicationDeadline),
	}

	return jobDetailsResponse, nil
}

func (js *JobServer) ApplyJob(ctx context.Context, req *pb.ApplyJobRequest) (*pb.ApplyJobResponse, error) {
	fmt.Println("Applying for job...")

	jobApplication := models.ApplyJob{
		JobID:       req.JobId,
		JobseekerID: req.JobseekerId,
		CoverLetter: req.CoverLetter,
		Resume:      req.ResumeData,
	}

	Data, err := js.jobUseCase.ApplyJob(jobApplication, req.ResumeData)
	if err != nil {
		return nil, err
	}

	return &pb.ApplyJobResponse{
		Id:          int64(Data.ID),
		JobId:       Data.JobID,
		JobseekerId: Data.JobseekerID,
		CoverLetter: Data.CoverLetter,
		ResumeUrl:   Data.ResumeURL,
	}, nil
}

func (js *JobServer) GetJobApplications(ctx context.Context, req *pb.GetJobApplicationsRequest) (*pb.GetJobApplicationsResponse, error) {
	employerID, err := strconv.ParseInt(req.EmployerId, 10, 64)
	if err != nil {
		return nil, err
	}

	applications, err := js.jobUseCase.GetApplicants(employerID)
	if err != nil {
		return nil, err
	}

	var applicationResponses []*pb.JobApplication
	for _, application := range applications {
		applicationResponse := &pb.JobApplication{
			Id:          strconv.FormatUint(uint64(application.ID), 10),
			JobId:       strconv.FormatUint(uint64(application.JobID), 10),
			JobSeekerId: strconv.FormatUint(uint64(application.JobseekerID), 10),
			Resume:      application.ResumeURL,
			CoverLetter: application.CoverLetter,
		}
		applicationResponses = append(applicationResponses, applicationResponse)
	}

	return &pb.GetJobApplicationsResponse{JobApplications: applicationResponses}, nil
}

func (js *JobServer) SaveJobs(ctx context.Context, req *pb.SaveJobRequest) (*pb.SaveJobResponse, error) {
	JobID, err := strconv.ParseInt(req.JobId, 10, 64)
	if err != nil {
		return nil, err
	}

	UserID, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	savedJob, err := js.jobUseCase.SaveJobs(JobID, UserID)
	if err != nil {
		return nil, err
	}

	response := &pb.SaveJobResponse{
		Id:      strconv.FormatUint(uint64(savedJob.ID), 10),
		JobId:   strconv.FormatInt(savedJob.JobID, 10),
		UserId:  strconv.FormatInt(savedJob.JobseekerID, 10),
		Message: "Job saved successfully",
	}

	return response, nil
}

func (js *JobServer) DeleteSavedJob(ctx context.Context, req *pb.DeleteSavedJobRequest) (*pb.DeleteSavedJobResponse, error) {
	jobID, err := strconv.ParseInt(req.JobId, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid job ID: %w", err)
	}

	userID, err := strconv.ParseInt(req.UserId, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	err = js.jobUseCase.DeleteSavedJob(int32(jobID), int32(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to delete saved job: %w", err)
	}

	return &pb.DeleteSavedJobResponse{
		Message: "Job deleted successfully",
	}, nil
}

func (js *JobServer) GetSavedJobs(ctx context.Context, req *pb.GetSavedJobsRequest) (*pb.GetSavedJobsResponse, error) {
	userID, err := strconv.ParseInt(req.UserId, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	savedJobs, err := js.jobUseCase.GetSavedJobs(int32(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to get saved jobs: %w", err)
	}

	var savedJobsResponse []*pb.SavedJobResponse
	for _, savedJob := range savedJobs {
		savedJobsResponse = append(savedJobsResponse, &pb.SavedJobResponse{
			Id:     strconv.FormatInt(int64(savedJob.ID), 10),
			JobId:  strconv.FormatInt(savedJob.JobID, 10),
			UserId: strconv.FormatInt(savedJob.JobseekerID, 10),
		})
	}

	return &pb.GetSavedJobsResponse{SavedJobs: savedJobsResponse}, nil
}

func (js *JobServer) ScheduleInterview(ctx context.Context, req *pb.ScheduleInterviewRequest) (*pb.ScheduleInterviewResponse, error) {
	jobID, err := strconv.ParseInt(req.JobId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid job ID: %w", err)
	}
	jobseekerID, err := strconv.ParseInt(req.JobseekerId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid jobseeker ID: %w", err)
	}

	employerID := req.EmployerId
	scheduledTime, err := time.Parse(time.RFC3339, req.ScheduledTime)
	if err != nil {
		return nil, fmt.Errorf("invalid scheduled time: %w", err)
	}

	mode := req.Mode
	if mode != "ONLINE" && mode != "OFFLINE" {
		return nil, fmt.Errorf("invalid mode: must be either 'ONLINE' or 'OFFLINE'")
	}

	status := req.Status
	if status != "SCHEDULED" && status != "COMPLETED" && status != "CANCELLED" {
		return nil, fmt.Errorf("invalid status: must be 'SCHEDULED', 'COMPLETED', or 'CANCELLED'")
	}

	interview := models.Interview{
		JobID:         jobID,
		JobseekerID:   jobseekerID,
		EmployerID:    employerID,
		ScheduledTime: scheduledTime,
		Mode:          mode,
		Link:          req.Link,
		Status:        status,
	}

	savedInterview, err := js.jobUseCase.ScheduleInterview(interview)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule interview: %w", err)
	}

	response := &pb.ScheduleInterviewResponse{
		Id:            strconv.FormatInt(int64(savedInterview.ID), 10),
		JobId:         strconv.FormatInt(savedInterview.JobID, 10),
		JobseekerId:   strconv.FormatInt(savedInterview.JobseekerID, 10),
		EmployerId:    savedInterview.EmployerID,
		ScheduledTime: savedInterview.ScheduledTime.Format(time.RFC3339),
		Mode:          savedInterview.Mode,
		Link:          savedInterview.Link,
		Status:        savedInterview.Status,
	}

	return response, nil
}
func (js *JobServer) GetInterview(ctx context.Context, req *pb.GetInterviewRequest) (*pb.GetInterviewsResponse, error) {
	jobID, err := strconv.ParseInt(req.JobId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid job ID: %w", err)
	}
	employerID := req.EmployerId

	interviewDetails, err := js.jobUseCase.GetInterview(int32(jobID), employerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get interview details: %w", err)
	}

	response := &pb.GetInterviewsResponse{
		Id:            uint64(interviewDetails.ID),
		JobId:         uint64(interviewDetails.JobID),
		JobseekerId:   uint64(interviewDetails.JobseekerID),
		EmployerId:    employerID,
		ScheduledTime: timestamppb.New(interviewDetails.ScheduledTime),
		Mode:          interviewDetails.Mode,
		Link:          interviewDetails.Link,
		Status:        interviewDetails.Status,
	}

	return response, nil
}
