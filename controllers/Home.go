package controllers

import (
	Api "baca-manga/helpers"
	responses "baca-manga/response"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	var HomeResp responses.MenuResponses

	resp, _ := Api.Get("http://komikindo.id/")
	if resp.StatusCode != 200 {
		c.JSON(500, responses.Error{
			Status:  "error",
			Message: "Something went wrong",
		})
	}

	data, errs := goquery.NewDocumentFromReader(resp.Body)
	if errs != nil {
		c.JSON(500, responses.Error{
			Status:  "error",
			Message: errs.Error(),
		})
	}
	data.Find("#menu-second-menu").Children().Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		url := s.Find("a").AttrOr("href", "No url")
		endpoint := strings.Replace(url, "http://komikindo.id/", "", -1)
		HomeResp.Menu = append(HomeResp.Menu, responses.DetailResponses{
			Name:     name,
			Url:      url,
			Endpoint: endpoint,
		})
	})
}
