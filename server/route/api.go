package route

import (
	"errors"
	"time"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/server/db"
)

type NewLinkForm struct {
	Url     string   `form:"url" json:"url" binding:"required"`
	Tags    []string `form:"tags" json:"tags"`
	Comment string   `form:"comment" json:"comment"`
}

func setLoginSession(c *gin.Context, session sessions.Session, userId int, userName string) {
	session.Set("userId", userId)
	session.Set("salt", time.Now().Unix())
	c.SetCookie("name", userName, 0, "/", "", !gin.IsDebugging(), false)
}

func getUserIdFromSession(session sessions.Session) (int, error) {
	value := session.Get("userId")
	if value == nil {
		return -1, errors.New("Not logged in")
	} else {
		return value.(int), nil
	}
}

func LoginPost(c *gin.Context) {
	githubId := c.PostForm("githubId")

	user, err := db.GetUserByGithubId(githubId)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	session := sessions.Default(c)
	setLoginSession(c, session, user.Id, user.Name)
	err = session.Save()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func GetLinks(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func NewLink(c *gin.Context) {
	session := sessions.Default(c)
	userId, err := getUserIdFromSession(session)
	if err != nil {
		c.JSON(401, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	var form NewLinkForm
	if c.Bind(&form) == nil {
		err := db.NewLink(form.Url, form.Tags, form.Comment, userId)
		if err != nil {
			c.JSON(500, gin.H{
				"status": "FAIL",
				"msg":    err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "OK",
		})
	} else {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    "Bind failed??",
		})
	}
}
