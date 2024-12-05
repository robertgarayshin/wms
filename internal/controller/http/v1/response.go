package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response - тип ответа от API.
type response struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	Message       any    `json:"message"`
	Error         string `json:"error" example:"message"`
}

// errorResponse - возвращает ответ при наличии ошибки.
func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{Status: code,
		StatusMessage: http.StatusText(code),
		Error:         msg,
	})
}

// successResponse - возварт успешного ответа от сервера.
func successResponse(c *gin.Context, code int, msg any) {
	c.JSON(code, response{
		Status:        code,
		StatusMessage: http.StatusText(code),
		Message:       msg,
	})
}
