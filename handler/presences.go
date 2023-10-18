package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type PresencesRequest struct {
	IncludeBackgroundSessions *bool     `form:"includeBackgroundSessions" json:"includeBackgroundSessions"`
	SelectedConversationId    *string   `form:"selectedConversationId" json:"selectedConversationId"`
	HasFocus                  *bool     `form:"hasFocus" json:"hasFocus"`
	IsTyping                  *bool     `form:"isTyping" json:"isTyping"`
	UserIDs                   *[]string `form:"userIds" json:"userIds"`
}

func Presences(c *gin.Context) {
	var req PresencesRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	r := resty.New()

	resp, err := r.R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetBody(req).
		SetAuthToken(fmt.Sprint(os.Getenv("TALKJS_SECRET_KEY"))).
		Post(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/presences"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	var response interface{}
	_ = json.Unmarshal(resp.Body(), &response)

	c.JSON(http.StatusOK, response)
}
