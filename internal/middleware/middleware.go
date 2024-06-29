package middleware

import (
	"bytes"
	"cmd/catapi/internal/logger"
	"io"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetLogger()

		rw := logger.NewResponseWriter(c.Writer)
		c.Writer = rw

		requestBody, _ := io.ReadAll(c.Request.Body)
		log.Printf("Incoming request: %s %s from %s\nRequest Body: %s",
			c.Request.Method, c.Request.URL.Path, c.ClientIP(), string(requestBody))

		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		c.Next()

		statusCode := c.Writer.Status()
		log.Printf("Response status: %d\nResponse Body: %s", statusCode, rw.Body.String())
	}
}
