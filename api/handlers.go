package api

import (
	"net/http"
	"typing-speed-test/database"
	"typing-speed-test/models"

	"github.com/gin-gonic/gin"
)

func HandlerLogin(c *gin.Context, s *database.Store) {
	var data models.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if data.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя пользователя не заполнено"})
		return
	}

	if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пароль не заполнен"})
		return
	}
	var password string
	password, err = s.GetPasswordByUsername(data.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}
	if password != data.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Вход выполнен"})
}

func HandlerRegister(c *gin.Context, s *database.Store) {
	c.JSON(http.StatusOK, gin.H{"message": "Регистрация прошла успешно"})
}

func MakeHandlerLogin(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandlerLogin(c, s)
	}
}

func MakeHandlerRegister(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		HandlerRegister(c, s)
	}
}
