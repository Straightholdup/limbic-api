package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
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

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "askhat1008@gmail.com"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/users/callback",
		ClientID:     "441968322627-ar2gmsjmjaaightomde8tme72ghhme7c.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func (h handler) Refresh(c *gin.Context) {

}

func (h handler) HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h handler) HandleGoogleCallback(c *gin.Context) {
	content, err := getUserInfo(c.Request.URL.Query().Get("state"), c.Request.URL.Query().Get("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Println(content)
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
