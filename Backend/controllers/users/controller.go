package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"limbic/controllers/auth"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/users")
	routes.GET("/", auth.IsAuthenticated(db), h.GetUsers)
	routes.GET("/personal-info", auth.IsAuthenticated(db), h.GetPersonalInfo)
}
