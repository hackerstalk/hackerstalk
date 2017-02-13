package route

import (
	"errors"
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
	limit := 25

	// Page 파라메터 파싱
	page64, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 32)
	if err != nil {
		FAIL(c, 400, err)
		return
	}

	page := int(page64)

	userIdStr := c.DefaultQuery("user_id", "")
	if userIdStr == "" {
		items, err := db.GetLinks((page-1)*limit, limit)
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
			"items": items,
			"total": count,
			"limit": limit,
		})
	} else {
		userId64, err := strconv.ParseInt(userIdStr, 10, 32)
		if err != nil {
			FAIL(c, 400, err)
			return
		}

		userId := int(userId64)

		items, err := db.GetLinksByUser(userId, (page-1)*limit, limit)
		if err != nil {
			FAIL(c, 500, err)
			return
		}

		count, err := db.GetLinkCountByUser(userId)
		if err != nil {
			FAIL(c, 500, err)
			return
		}

		OK(c, gin.H{
			"items": items,
			"total": count,
			"limit": limit,
		})
	}

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
	OK(c)

}

// 링크 수정
func EditLink(c *gin.Context) {
	session := sessions.Default(c)
	userId, err := getUserIdFromSession(session)
	if err != nil {
		FAIL(c, 401, err)
		return
	}

	linkIdStr := c.Param("link_id")
	linkId64, err := strconv.ParseInt(linkIdStr, 10, 32)
	if err != nil {
		FAIL(c, 401, err)
		return
	}
	linkId := int(linkId64)

	// 본인의 링크가 맞는지 확인
	// Transaction을 할 필요는 없음
	check, err := db.CheckLinkOwner(userId, linkId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	if check == false {
		FAIL(c, 401, errors.New("권한이 없습니다."))
		return
	}

	var form NewLinkForm
	err = c.Bind(&form)
	if err != nil {
		FAIL(c, 400, err)
		return
	}

	err = db.EditLink(linkId, form.Url, form.Tags, form.Comment, userId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	OK(c)
}

// 링크 삭제
func DelLink(c *gin.Context) {
	session := sessions.Default(c)
	userId, err := getUserIdFromSession(session)
	if err != nil {
		FAIL(c, 401, err)
		return
	}

	linkIdStr := c.Param("link_id")
	linkId64, err := strconv.ParseInt(linkIdStr, 10, 32)
	if err != nil {
		FAIL(c, 401, err)
		return
	}
	linkId := int(linkId64)

	// 본인의 링크가 맞는지 확인
	// Transaction을 할 필요는 없음
	check, err := db.CheckLinkOwner(userId, linkId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}
	if check == false {
		FAIL(c, 401, errors.New("권한이 없습니다."))
		return
	}

	err = db.DeleteLink(linkId)
	if err != nil {
		FAIL(c, 500, err)
		return
	}

	OK(c)
}
