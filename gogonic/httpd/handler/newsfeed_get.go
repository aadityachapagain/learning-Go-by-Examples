package handler

import (
	"gogonic/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsFeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, feed.GetAll())
	}
}
