package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"limbic/models"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h handler) Login(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var user models.User
	if err := h.DB.Where(&models.User{Email: u.Email}).First(&user).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Email is not registered")
		return
	}
	if user.Password != u.Password {
		c.JSON(http.StatusUnprocessableEntity, "Invalid email or password")
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: &jwt.NumericDate{expirationTime},
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusUnprocessableEntity, "Invalid token provider")
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	//c.SetCookie("token", )
	//http.SetCookie(c.Writer, &http.Cookie{
	//	Name:    "token",
	//	Value:   tokenString,
	//	Expires: expirationTime,
	//})
	c.SetCookie("token", tokenString, int((5 * time.Minute).Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, user)
}
