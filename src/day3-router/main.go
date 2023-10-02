package main

import (
	"gee_day3"
	"net/http"
)

func main() {
	r := gee_day3.New()
	r.GET("/", func(c *gee_day3.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	
	r.GET("/hello", func(c *gee_day3.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	
	r.GET("/hello/:name", func(c *gee_day3.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	
	r.GET("/assets/*filepath", func(ctx *gee_day3.Context) {
		ctx.JSON(http.StatusOK, gee_day3.H{"filepath": ctx.Param("filepath")})
	})
	
	r.Run(":9999")
}