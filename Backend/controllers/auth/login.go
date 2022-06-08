package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"limbic/models"
	"net/http"
	"time"
)

func (h handler) Login(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if user.Email != u.Email || user.Password != u.Password {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

var user = models.User{
	Email:    "username",
	Password: "password",
}

func CreateToken(userid uint) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("jdnfksdmfksd"))
	if err != nil {
		return "", err
	}
	return token, nil
}
