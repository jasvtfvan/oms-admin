package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		defer func() {
			clientIP := c.ClientIP()
			proto := c.Request.Proto
			method := c.Request.Method
			path := c.Request.URL.Path
			statusCode := c.Writer.Status()
			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
			latency := time.Since(start)
			userAgent := c.Request.UserAgent()

			logStr := fmt.Sprintf("[GIN] %s [%s] | \t %s %s [%s] | \t %d %s [%s] \n[GIN:UserAgent] %s",
				start.Format(time.DateTime),
				clientIP,
				proto,
				method,
				path,
				statusCode,
				errorMessage,
				latency.String(),
				userAgent,
			)

			logger.Debug(logStr)
		}()

		c.Next()
	}
}
