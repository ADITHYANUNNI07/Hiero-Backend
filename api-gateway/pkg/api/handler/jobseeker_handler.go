package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobSeekerHandler struct {
	GRPC_Client interfaces.JobSeekerClient
}

func NewJobSeekerHandler(jobSeekerClient interfaces.JobSeekerClient) *JobSeekerHandler {
	return &JobSeekerHandler{
		GRPC_Client: jobSeekerClient,
	}
}

func (jh *JobSeekerHandler) JobSeekerLogin(c *gin.Context) {
	var jobSeekerDetails models.JobSeekerLogin
	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerLogin(jobSeekerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenticate job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Job seeker authenticated successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}

func (jh *JobSeekerHandler) JobSeekerSignUp(c *gin.Context) {
	var jobSeekerDetails models.JobSeekerSignUp

	if err := c.ShouldBindJSON(&jobSeekerDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	jobSeeker, err := jh.GRPC_Client.JobSeekerSignUp(jobSeekerDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot create job seeker", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Job seeker created successfully", jobSeeker, nil)
	c.JSON(http.StatusOK, success)
}
