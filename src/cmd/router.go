package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/ntfargo/tir-goapi/src/internal/controller"
)

const (
	apiDocsURL       = "https://documenter.getpostman.com/view/795261/2s9Xy6pUdQ#ee63743c-87e3-471f-8b24-370431aea6b4"
	requestRateLimit = 10
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	limiter := ratelimit.NewBucket(time.Second, requestRateLimit)
	r.Use(RequestTimeMiddleware, RateLimiterMiddleware(limiter))

	// main routes
	r.GET("/", homeHandler)
	r.NoRoute(notFoundHandler)

	// user routes
	r.POST("/users", controller.RegisterUser)

	return r
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "Welcome to TIR Go API!",
		"message": "Explore our API and discover its powerful features. Check the documentation for more information.",
		"docsURL": apiDocsURL,
	})
}

func notFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": "404 Not Found",
		"request": c.Request.URL.Path,
	})
}
