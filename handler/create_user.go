package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type CreateUserRequest struct {
	Name           string            `form:"name" json:"name"`
	Email          []string          `form:"email" json:"email"`
	WelcomeMessage string            `form:"welcomeMessage" json:"welcomeMessage"`
	PhotoURL       string            `form:"photoUrl" json:"photoUrl"`
	Role           string            `form:"role" json:"role"`
	Phone          []string          `form:"phone" json:"phone"`
	Custom         map[string]string `form:"custom" json:"custom"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest

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
		Put(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/users/", c.Param("id")))
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
