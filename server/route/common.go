package route

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"
)

// 세션에서 userId를 가져온다.
func getUserIdFromSession(session sessions.Session) (int, error) {
	value := session.Get("userId")
	if value == nil {
		return -1, errors.New("Not logged in")
	} else {
		return value.(int), nil
	}
}

// OK response를 리턴
func OK(c *gin.Context, obj interface{}) {
	c.JSON(200, obj)
}

// Error response를 리턴
func FAIL(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"status": "FAIL",
		"msg":    err.Error(),
	})
}
