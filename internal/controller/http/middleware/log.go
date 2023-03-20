package middleware

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// log.Info("[GIN]", "method", method, "path", path, "raw", raw, "status", statusCode, "latency", latency, "clientIP", clientIP)
		log.Info("[GIN]", "time", timeStamp.Format("2006-01-02 15:04:05"), "method", method, "path", path, "status", statusCode, "latency", latency, "clientIP", clientIP)
	}
}
