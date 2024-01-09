package validate

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserLoginCheck struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateHeader(c *gin.Context) (int, error) {
	if c.Request.Header["Authorization"] == nil || len(c.Request.Header["Authorization"][0]) <= 0 {
		return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	}
	// else {
	// 	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	// 	if token != "$2y$10$uysIMpMRXfx..DG2CHaDZ.9wuuxwbJEIoqlppUQXTcKya4z6dcLja" {
	// 		return http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized))
	// 	}
	// }
	return 0, nil
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("Echo31"), nil
	})

	return err
}

func ParameterValicate(c *gin.Context) (int, error) {
	var user_login_check UserLoginCheck
	if err := c.ShouldBindJSON(&user_login_check); err != nil {
		return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	}
	return 0, nil
}
