package api

import (
	"typing-speed-test/database"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, s *database.Store) {
	r.POST("/login", MakeHandlerLogin(s))
	r.POST("/register", MakeHandlerRegister(s))
}
