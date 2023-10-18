package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"test-talkjs/models"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func Users(c *gin.Context) {
	r := resty.New()

	resp, err := r.R().
		SetAuthToken(fmt.Sprint(os.Getenv("TALKJS_SECRET_KEY"))).
		Get(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/users"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	type UserList struct {
		Data []models.User `json:"data"`
	}

	var response UserList

	_ = json.Unmarshal(resp.Body(), &response)

	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	r := resty.New()

	resp, err := r.R().
		SetAuthToken(fmt.Sprint(os.Getenv("TALKJS_SECRET_KEY"))).
		Get(fmt.Sprint(os.Getenv("TALKJS_HOST"), "/users/", c.Param("id")))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	var response models.User
	_ = json.Unmarshal(resp.Body(), &response)

	c.JSON(http.StatusOK, response)
}
