package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func DeleteConversations(c *gin.Context) {
	r := resty.New()

	resp, err := r.R().
		SetAuthToken(fmt.Sprint(os.Getenv("TALKJS_SECRET_KEY"))).
		Delete(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/conversations/", c.Param("id")))
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
