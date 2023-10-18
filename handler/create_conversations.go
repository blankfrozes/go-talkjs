package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type CreateConversationsRequest struct {
	Participants    []string           `form:"participants" json:"participants"`
	Subject         *string            `form:"subject" json:"subject"`
	WelcomeMessages *[]string          `form:"welcomeMessages" json:"welcomeMessages"`
	Custom          *map[string]string `form:"custom" json:"custom"`
	PhotoURL        *string            `form:"photoUrl" json:"photoUrl"`
}

func CreateConversations(c *gin.Context) {
	var req CreateConversationsRequest

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
		Put(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/conversations/", c.Param("id")))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
