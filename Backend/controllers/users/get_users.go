package users

import (
	"github.com/gin-gonic/gin"
	"limbic/models"
	"net/http"
)

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User
	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &users)
}
