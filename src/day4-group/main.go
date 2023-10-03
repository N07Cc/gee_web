package main

import (
	"gee_day4"
	"net/http"
)

func main() {
	r := gee_day4.New()
	r.GET("/index", func(c *gee_day4.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee_day4.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		
		v1.GET("/hello", func(c *gee_day4.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee_day4.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
		
		v2.POST("/login", func(c *gee_day4.Context) {
			c.JSON(http.StatusOK, gee_day4.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	
	r.Run(":9999")
}