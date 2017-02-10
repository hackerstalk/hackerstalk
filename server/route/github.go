package route

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/server/auth"
	"hackerstalk/server/db"
	"hackerstalk/server/util"
)

func GithubAuth(c *gin.Context) {
	session := sessions.Default(c)
	state, url := auth.GetGithubAuthUrl()
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}
	c.Redirect(http.StatusFound, url)
}

func GithubAuthCallback(c *gin.Context) {
	session := sessions.Default(c)
	apiError := c.Query("error")
	apiErrorDescription := c.Query("error_description")
	code := c.Query("code")

	if apiError != "" {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"error":  apiError,
			"msg":    apiErrorDescription,
		})
		return
	}

	savedState := util.GetDefault(session, "state", "")
	if savedState != c.Query("state") {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    "state mismatches. " + savedState + " != " + c.Query("state"),
		})
		return
	}

	githubUser, err := auth.GetGithubUser(code)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	userName := *githubUser.Name
	githubId := *githubUser.Login
	err = db.UpsertUser(userName, githubId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}
	session.Delete("state")

	var user *db.User
	user, err = db.GetUserByGithubId(githubId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}
	setLoginSession(c, session, user.Id, userName)
	err = session.Save()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
