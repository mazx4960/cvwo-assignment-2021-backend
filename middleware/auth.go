package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user_id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tk, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tk, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Malformed token",
			})
			c.Abort()
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRET_KEY")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("user_id", uint(claims["user_id"].(float64)))
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}