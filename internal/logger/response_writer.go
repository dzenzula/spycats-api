package logger

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func NewResponseWriter(w gin.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		Body:           bytes.NewBufferString(""),
	}
}
