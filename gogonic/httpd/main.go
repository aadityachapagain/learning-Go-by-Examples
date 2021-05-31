package main

import (
	"gogonic/httpd/handler"
	"gogonic/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	feed := newsfeed.New()

	r.GET("/", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))

	r.Run()
}
