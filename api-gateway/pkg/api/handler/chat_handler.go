package handler

import (
	interfaces "HireoGateWay/pkg/client/interface"
	"HireoGateWay/pkg/helper"
	"HireoGateWay/pkg/utils/models"
	"HireoGateWay/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

type ChatHandler struct {
	GRPC_Client interfaces.ChatClient
	helper      *helper.Helper
}

func NewChatHandler(chatClient interfaces.ChatClient, helper *helper.Helper) *ChatHandler {
	return &ChatHandler{
		GRPC_Client: chatClient,
		helper:      helper,
	}
}

func (ch *ChatHandler) EmployerMessage(c *gin.Context) {

	fmt.Println("message")

	tokenString := c.Request.Header.Get("Authorization")
	fmt.Println("message")
	splitToken := strings.Split(tokenString, " ")
	if tokenString == "" {
		errs := response.ClientResponse(http.StatusUnauthorized, "Missing Authorization header", nil, "")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	splitToken[1] = strings.TrimSpace(splitToken[1])
	userID, err := ch.helper.ValidateToken(splitToken[1])
	fmt.Println("validate token result ", userID, err)
	if err != nil {
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid token", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	fmt.Println("upgrading ")
	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Websocket Connection Issue", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	defer delete(User, strconv.Itoa(int(userID.Id)))
	defer conn.Close()
	user := strconv.Itoa(int(userID.Id))
	User[user] = conn

	for {
		fmt.Println("loop starts", userID, User)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		fmt.Println("message sting msg ", string(msg))
		fmt.Println("message soket conn map ", User)
		fmt.Println("message user id ", user)
		ch.helper.SendMessageToUser(User, msg, user)
	}
}

// func (ch *ChatHandler) GetChat(c *gin.Context) {
// 	var chatRequest models.ChatRequest
// 	if err := c.ShouldBindJSON(&chatRequest); err != nil {
// 		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errs)
// 		return
// 	}

// 	userIDInterface, exists := c.Get("id")
// 	if !exists {
// 		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
// 		c.JSON(http.StatusBadRequest, errs)
// 		return
// 	}
// 	userID := strconv.Itoa(userIDInterface.(int))

// 	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)
// 	if err != nil {
// 		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errs)
// 		return
// 	}

// 	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
// 	c.JSON(http.StatusOK, errs)
// }

func (ch *ChatHandler) GetChat(c *gin.Context) {
	var chatRequest models.ChatRequest
	if err := c.ShouldBindJSON(&chatRequest); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userIDInterface, exists := c.Get("id")
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID := strconv.Itoa(int(userIDInterface.(int32)))

	result, err := ch.GRPC_Client.GetChat(userID, chatRequest)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get chat details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	errs := response.ClientResponse(http.StatusOK, "Successfully retrieved chat details", result, nil)
	c.JSON(http.StatusOK, errs)
}
