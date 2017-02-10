package route

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"gopkg.in/gin-gonic/gin.v1"

	"hackerstalk/server/db"
)

type NewLinkForm struct {
	Url     string   `form:"url" json:"url" binding:"required"`
	Tags    []string `form:"tags" json:"tags"`
	Comment string   `form:"comment" json:"comment"`
}

func GetLinks(c *gin.Context) {
	limit := 50

	// Page 파라메터 파싱
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		c.JSON(401, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	items, err := db.GetLinks((int(page)-1)*limit, limit)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	count, err := db.GetLinkCount()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "FAIL",
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
		"items":  items,
		"total":  count,
		"limit":  limit,
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
