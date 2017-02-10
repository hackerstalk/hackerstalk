package route

import (
	"errors"
	"time"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
)

func getUserIdFromSession(session sessions.Session) (int, error) {
	value := session.Get("userId")
	if value == nil {
		return -1, errors.New("Not logged in")
	} else {
		return value.(int), nil
	}
}

func setLoginSession(c *gin.Context, session sessions.Session, userId int, userName string) {
	session.Set("userId", userId)
	session.Set("salt", time.Now().Unix())
	c.SetCookie("name", userName, 0, "/", "", false, false)
}
