package users

import (
	"github.com/gin-gonic/gin"
	"limbic/models"
	"net/http"
)

func (h handler) GetPersonalInfo(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(http.StatusOK, user)
}
