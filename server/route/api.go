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

// 링크 목록 가져오는 핸들러
func GetLinks(c *gin.Context) {
	limit := 50

	// Page 파라메터 파싱
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		FAIL(c, 400, err)
		return
	}

	items, err := db.GetLinks((int(page)-1)*limit, limit)
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	count, err := db.GetLinkCount()
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	OK(c, gin.H{
		"status": "OK",
		"items":  items,
		"total":  count,
		"limit":  limit,
	})
}

// 새로운 링크 추가 핸들러
func NewLink(c *gin.Context) {
	session := sessions.Default(c)
	userId, err := getUserIdFromSession(session)
	if err != nil {
		FAIL(c, 401, err)
		return
	}

	var form NewLinkForm
	err = c.Bind(&form)
	if err != nil {
		FAIL(c, 400, err)
		return
	}

	err = db.NewLink(form.Url, form.Tags, form.Comment, userId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	OK(c, gin.H{
		"status": "OK",
	})

}
