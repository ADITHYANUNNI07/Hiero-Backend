package helper

import (
	"HireoGateWay/pkg/config"
	"HireoGateWay/pkg/utils/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

type Helper struct {
	config *config.Config
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		config: config,
	}
}

func (r *Helper) SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	var message models.Message
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		fmt.Println("error while unmarshel ", err)
	}

	message.SenderID = userID
	recipientConn, ok := User[message.RecipientID]
	fmt.Println("recipient id", message.RecipientID)
	if ok {
		recipientConn.WriteMessage(websocket.TextMessage, msg)
	}
	err := KafkaProducer(message)
	fmt.Println("==sending succesful==", err)
}

func KafkaProducer(message models.Message) error {
	fmt.Println("from kafka ", message)

	// cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, configs)
	if err != nil {
		fmt.Println("error creating producer:", err)
		return err
	}
	fmt.Println("producer created successfully:", producer)

	result, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("Friend message"),
		Value: sarama.StringEncoder(result),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("error sending message:", err)
		return err
	}

	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	fmt.Println("==sending successful==")
	return nil
}

func (h *Helper) ValidateToken(tokenString string) (*authCustomClaimsEmployer, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsEmployer{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsEmployer); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
