package client

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/config"
	pb "HireoGateWay/pkg/pb/auth"
	"HireoGateWay/pkg/utils/models"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

type employerClient struct {
	Client pb.EmployerClient
}

func NewEmployerClient(cfg config.Config) interfaces.EmployerClient {
	grpcConnection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewEmployerClient(grpcConnection)

	return &employerClient{
		Client: grpcClient,
	}
}

func (ec *employerClient) EmployerSignUp(employerDetails models.EmployerSignUp) (models.TokenEmployer, error) {
	employer, err := ec.Client.EmployerSignup(context.Background(), &pb.EmployerSignupRequest{
		CompanyName:         employerDetails.Company_name,
		Industry:            employerDetails.Industry,
		CompanySize:         int32(employerDetails.Company_size),
		Website:             employerDetails.Website,
		HeadquartersAddress: employerDetails.Headquarters_address,
		AboutCompany:        employerDetails.About_company,
		ContactEmail:        employerDetails.Contact_email,
		ContactPhoneNumber:  uint64(employerDetails.Contact_phone_number),
		Password:            employerDetails.Password,
	})
	if err != nil {
		return models.TokenEmployer{}, err
	}
	return models.TokenEmployer{
		Employer: models.EmployerDetailsResponse{
			ID:                   uint(employer.EmployerDetails.Id),
			Company_name:         employer.EmployerDetails.CompanyName,
			Industry:             employer.EmployerDetails.Industry,
			Company_size:         int(employer.EmployerDetails.CompanySize),
			Website:              employer.EmployerDetails.Website,
			Headquarters_address: employer.EmployerDetails.HeadquartersAddress,
			About_company:        employer.EmployerDetails.AboutCompany,
			Contact_email:        employer.EmployerDetails.ContactEmail,
			Contact_phone_number: uint(employer.EmployerDetails.ContactPhoneNumber),
		},
		Token: employer.Token,
	}, nil
}

func (ec *employerClient) EmployerLogin(employerDetails models.EmployerLogin) (models.TokenEmployer, error) {
	employer, err := ec.Client.EmployerLogin(context.Background(), &pb.EmployerLoginInRequest{
		Email:    employerDetails.Email,
		Password: employerDetails.Password,
	})

	if err != nil {
		return models.TokenEmployer{}, err
	}
	return models.TokenEmployer{
		Employer: models.EmployerDetailsResponse{
			ID:                   uint(employer.EmployerDetails.Id),
			Company_name:         employer.EmployerDetails.CompanyName,
			Industry:             employer.EmployerDetails.Industry,
			Company_size:         int(employer.EmployerDetails.CompanySize),
			Website:              employer.EmployerDetails.Website,
			Headquarters_address: employer.EmployerDetails.HeadquartersAddress,
			About_company:        employer.EmployerDetails.AboutCompany,
			Contact_email:        employer.EmployerDetails.ContactEmail,
			Contact_phone_number: uint(employer.EmployerDetails.ContactPhoneNumber),
		},
		Token: employer.Token,
	}, nil
}
func (jc *employerClient) GetCompanyDetails(employerIDInt int32) (models.EmployerDetailsResponse, error) {
	resp, err := jc.Client.GetCompanyDetails(context.Background(), &pb.GetCompanyDetailsRequest{Id: employerIDInt})
	if err != nil {
		return models.EmployerDetailsResponse{}, fmt.Errorf("failed to get company details: %v", err)
	}

	// Check if the response is nil
	if resp == nil {
		return models.EmployerDetailsResponse{}, errors.New("empty response received from server")
	}

	// Convert the protobuf response to the expected model
	employer := resp.GetEmployerDetails()
	if employer == nil {
		return models.EmployerDetailsResponse{}, errors.New("empty employer details received from server")
	}

	// Map the protobuf response to the model
	employerDetails := models.EmployerDetailsResponse{
		ID:                   uint(employer.Id),
		Company_name:         employer.CompanyName,
		Industry:             employer.Industry,
		Company_size:         int(employer.CompanySize),
		Website:              employer.Website,
		Headquarters_address: employer.HeadquartersAddress,
		About_company:        employer.AboutCompany,
		Contact_email:        employer.ContactEmail,
		Contact_phone_number: uint(employer.ContactPhoneNumber),
	}

	return employerDetails, nil
}
func (ec *employerClient) UpdateCompany(employerIDInt int32, employerDetails models.EmployerDetails) (models.EmployerDetailsResponse, error) {

	employer, err := ec.Client.UpdateCompany(context.Background(), &pb.UpdateCompanyRequest{
		Id:                  int32(employerIDInt),
		CompanyName:         employerDetails.Company_name,
		Industry:            employerDetails.Industry,
		CompanySize:         int32(employerDetails.Company_size),
		Website:             employerDetails.Website,
		HeadquartersAddress: employerDetails.Headquarters_address,
		AboutCompany:        employerDetails.About_company,
		ContactEmail:        employerDetails.Contact_email,
		ContactPhoneNumber:  uint64(employerDetails.Contact_phone_number),
	})
	if err != nil {
		return models.EmployerDetailsResponse{}, err
	}
	return models.EmployerDetailsResponse{
		ID:                   uint(employerIDInt),
		Company_name:         employer.CompanyName,
		Industry:             employer.Industry,
		Company_size:         int(employer.CompanySize),
		Website:              employer.Website,
		Headquarters_address: employer.HeadquartersAddress,
		About_company:        employer.AboutCompany,
		Contact_email:        employer.ContactEmail,
		Contact_phone_number: uint(employer.ContactPhoneNumber),
	}, nil
}
