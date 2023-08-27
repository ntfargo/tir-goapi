package main

import (
	"fmt"
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
		tokensAvailable := limiter.TakeAvailable(1)
		if tokensAvailable <= 0 {
			logRateLimitEvent(c)
			response := buildRateLimitResponse()
			c.JSON(http.StatusTooManyRequests, response)
			c.Abort()
			return
		}
		c.Next()
	}
}

func logRateLimitEvent(c *gin.Context) {
	fmt.Printf(
		"Rate limit hit at %s for IP: %s, route: %s\n",
		time.Now().Format(time.RFC3339),
		c.ClientIP(),
		c.FullPath(),
	)
}

func buildRateLimitResponse() gin.H {
	return gin.H{
		"error":   "Too many requests",
		"message": "You have exceeded your request rate. Please try again later.",
		"code":    http.StatusTooManyRequests,
	}
}
