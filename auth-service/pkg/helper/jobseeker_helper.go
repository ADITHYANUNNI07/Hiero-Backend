package helper

import (
	"Auth/pkg/utils/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsJobSeeker struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenJobSeeker(jobSeeker models.JobSeekerDetailsResponse) (string, error) {
	claims := &authCustomClaimsJobSeeker{
		ID:    jobSeeker.ID,
		Email: jobSeeker.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenJobSeeker(tokenString string) (*authCustomClaimsJobSeeker, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsJobSeeker{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaimsJobSeeker); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
