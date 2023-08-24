package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	config "github.com/ntfargo/tir-goapi/src/internal/config"
)

const (
	apiDocsURL       = "https://documenter.getpostman.com/view/795261/2s9Xy6pUdQ#ee63743c-87e3-471f-8b24-370431aea6b4"
	requestRateLimit = 10
)

type serverRunner func(*gin.Engine, string) error

var envVariables, _ = config.LoadEnvVariables()

func setupServerRunner() serverRunner {
	sslMode, exists := envVariables["SSLMODE"]
	if !exists || sslMode == "disable" {
		return func(engine *gin.Engine, address string) error {
			return engine.Run(address)
		}
	}
	return func(engine *gin.Engine, address string) error {
		return engine.RunTLS(address, "/path/to/your/cert.pem", "/path/to/your/key.pem")
	}
}

func main() {
	port := ":" + config.GetPort()
	r := gin.Default()
	limiter := ratelimit.NewBucket(time.Second, 10)
	r.Use(RequestTimeMiddleware, RateLimiterMiddleware(limiter))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":   "Welcome to TIR Go API!",
			"message": "Explore our API and discover its powerful features. Check the documentation for more information.",
			"docsURL": apiDocsURL,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "404 Not Found",
		})
	})

	runner := setupServerRunner()
	if err := runner(r, port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func RequestTimeMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	elapsedTime := time.Since(startTime)
	log.Printf("Request processed in %s for %s", elapsedTime, c.Request.URL.Path)
}

func RateLimiterMiddleware(limiter *ratelimit.Bucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) <= 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
