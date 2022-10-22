package controllers

import (
	Api "baca-manga/helpers"
	responses "baca-manga/response"
	"github.com/gin-gonic/gin"
)

func ControllerKhoeru(c *gin.Context) {
	// get api from url
	resp := Api.Get("http://komikindo.id/")
	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"status":      "error",
			"message":     "Something went wrong",
			"status_code": resp.StatusCode,
		})
		return
	}

	// end execute
	defer resp.Body.Close()

	c.JSON(404, responses.KhoeruError{
		Status:     "error",
		Message:    "something missing",
		StatusCode: resp.StatusCode,
	})

}
