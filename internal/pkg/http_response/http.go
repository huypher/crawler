package http_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrInternal     = "có lỗi hệ thống xảy ra"
	ErrUnauthorized = "chưa đăng nhập"
)

func Response(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func Success(c *gin.Context, message string, data interface{}) {
	Response(c, http.StatusOK, message, data)
}

func NotAuthorized(c *gin.Context, message string) {
	Response(c, http.StatusUnauthorized, message, nil)
}

func Error(c *gin.Context, err error) {
	var httpCode int
	var message string
	switch err.Error() {
	case ErrUnauthorized:
		httpCode = http.StatusUnauthorized
		message = ErrUnauthorized
	default:
		httpCode = http.StatusInternalServerError
		//message = ErrInternal
		message = err.Error()
	}

	Response(c, httpCode, message, nil)
}

func Abort(c *gin.Context, err error) {
	var httpCode int
	var message string
	switch err.Error() {
	case ErrUnauthorized:
		httpCode = http.StatusUnauthorized
		message = ErrUnauthorized
	default:
		httpCode = http.StatusInternalServerError
		message = ErrInternal
	}

	c.AbortWithStatusJSON(httpCode, gin.H{
		"message": message,
	})
}
