package route

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/server/auth"
	"hackerstalk/server/db"
	"hackerstalk/server/util"
)

// GitHub OAuth로 Redirect.
func GithubAuth(c *gin.Context) {
	session := sessions.Default(c)
	state, url := auth.GetGithubAuthUrl()
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	c.Redirect(http.StatusFound, url)
}

// GitHub OAuth Callback 처리
func GithubAuthCallback(c *gin.Context) {
	session := sessions.Default(c)
	apiError := c.Query("error")
	apiErrorDescription := c.Query("error_description")
	code := c.Query("code")

	if apiError != "" {
		FAIL(c, 500, errors.New(apiError+" "+apiErrorDescription))
		return
	}

	savedState := util.GetDefault(session, "state", "")
	if savedState != c.Query("state") {
		FAIL(c, 500, errors.New("state mismatches. "+savedState+" != "+c.Query("state")))
		return
	}

	githubUser, err := auth.GetGithubUser(code)
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	userName := *githubUser.Name
	githubId := *githubUser.Login
	err = db.UpsertUser(userName, githubId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	session.Delete("state")

	var user *db.User
	user, err = db.GetUserByGithubId(githubId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	setLoginSession(c, session, user.Id, userName)
	err = session.Save()
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	c.Redirect(http.StatusFound, "/")
}
