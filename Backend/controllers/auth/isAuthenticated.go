package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"limbic/models"
	"net/http"
)

func IsAuthenticated(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// We can obtain the session token from the requests cookies, which come with every request
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := cookie

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := DB.Where(&models.User{Email: claims.Email}).First(&user).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Some unexpected error")
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
