package main

import (
	"gee_day2"
	"net/http"
)

func main() {
	r := gee_day2.New()
	
	r.GET("/", func(c *gee_day2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee_day2.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee_day2.Context) {
		c.JSON(http.StatusOK, gee_day2.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	
	r.Run(":9999")
}