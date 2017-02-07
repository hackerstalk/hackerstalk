package route

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/db"
	"hackerstalk/util"
)

func DebugSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func DebugSignInPost(c *gin.Context) {
	name := c.PostForm("name")
	githubId := c.PostForm("githubId")

	if name == "" || githubId == "" {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    "name of githubId is missing",
		})
		return
	}

	err := db.UpsertUser(name, githubId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func DebugSession(c *gin.Context) {
	session := sessions.Default(c)
	name := util.GetDefault(session, "name", "")
	githubId := util.GetDefault(session, "githubId", "")
	c.JSON(200, gin.H{
		"name":     name,
		"githubId": githubId,
	})
}
