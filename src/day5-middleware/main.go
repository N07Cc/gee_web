package main

import (
	"gee_day5"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee_day5.HandlerFunc {
	return func(c *gee_day5.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee_day5.New()
	r.Use(gee_day5.Logger())
	r.GET("/", func(ctx *gee_day5.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee_day5.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	
	r.Run(":9999")
}