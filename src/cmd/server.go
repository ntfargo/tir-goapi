package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	config "github.com/ntfargo/tir-goapi/src/internal/config"
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
		return engine.RunTLS(address, envVariables["CERT"], envVariables["KEY"])
	}
}

func main() {
	port := ":" + config.GetPort()
	r := setupRouter()

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
