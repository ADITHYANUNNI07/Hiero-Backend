package middleware

import (
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EmployerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("authorization")
		fmt.Println(tokenHeader, "this is the token header")
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateTokenEmployer(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		employerID := int32(tokenClaims.Id)

		c.Set("id", employerID)

		c.Next()
	}
}
