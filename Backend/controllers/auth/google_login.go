package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
	"io/ioutil"
	"limbic/models"
	"net/http"
	"time"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random string"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google_callback",
		ClientID:     "441968322627-ar2gmsjmjaaightomde8tme72ghhme7c.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
}

func (h handler) HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

type GoogleUser struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (h handler) HandleGoogleCallback(c *gin.Context) {
	content, err := getUserInfo(c.Request.URL.Query().Get("state"), c.Request.URL.Query().Get("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	var google_user GoogleUser
	if err := json.Unmarshal(content, &google_user); err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	var user models.User
	if err := h.DB.Where(&models.User{Email: google_user.Email}).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	} else if err == gorm.ErrRecordNotFound {
		user.Email = google_user.Email
		user.Name = google_user.Name
		user.Avatar = &google_user.Picture
		if err := h.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: user.Email,
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
	c.SetCookie("token", tokenString, int((5 * time.Minute).Seconds()), "/", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "http://localhost:3000/overview")
	//c.JSON(http.StatusOK, tokenString)
	//fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
