package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/users")
	routes.GET("/", h.GetUsers)
	routes.GET("/login", h.Login)
	routes.GET("/g_login", h.HandleGoogleLogin)
	routes.GET("/g_callback", h.HandleGoogleCallback)
}
