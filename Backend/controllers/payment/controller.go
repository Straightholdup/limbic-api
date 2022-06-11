package payment

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

	routes := r.Group("/payment")
	routes.POST("/create", auth.IsAuthenticated(db), h.CreatePayment)
	routes.POST("/confirm", h.ConfirmPayment)
}
