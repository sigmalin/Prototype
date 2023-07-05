package rgin

import (
	"net/http"
	"response/code"

	"github.com/gin-gonic/gin"
)

type GinResponse struct {
	context *gin.Context
	data    gin.H
}

func (r GinResponse) send() {
	r.context.JSON(http.StatusOK, r.data)
}

func (r GinResponse) Add(key string, value interface{}) {
	r.data[key] = value
}

func (r GinResponse) Error(code int, msg string) error {
	r.data = gin.H{"error": gin.H{"code": code, "message": msg}}
	r.send()
	return nil
}

func (r GinResponse) Message() error {
	r.data["error"] = gin.H{"code": code.SUCCESS, "message": ""}
	r.send()
	return nil
}
