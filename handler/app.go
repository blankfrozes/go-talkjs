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

func App(c *gin.Context) {
	r := resty.New()

	resp, err := r.R().
		SetAuthToken(fmt.Sprint(os.Getenv("TALKJS_SECRET_KEY"))).
		Get(fmt.Sprint(os.Getenv("TALKJS_HOST")))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	var response models.App
	_ = json.Unmarshal(resp.Body(), &response)

	c.JSON(http.StatusOK, response)
}
