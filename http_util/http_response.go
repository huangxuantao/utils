package http_util

import (
	"gitea.com/huangxuantao89/utils/errno_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseData struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"total_count,omitempty"`
}

type ResponseToken struct {
	Token interface{} `json:"token"`
}

func NormalResponseJson(c *gin.Context, err *errno_util.Errno, data interface{}) {
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Cache-Control", "no-cache, private")
	c.Writer.Header().Set("Connection", "keep-alive")

	var r ResponseJson
	switch obj := data.(type) {
	case error:
		r = ResponseJson{
			Code:    err.Code,
			Message: err.Message,
			Data:    obj.Error(),
		}
	default:
		r = ResponseJson{
			Code:    err.Code,
			Message: err.Message,
			Data:    data,
		}
	}
	c.Set("custom_code", err.Code)
	c.JSON(http.StatusOK, r)
}

func NormalResponseHtml(c *gin.Context, template string, err *errno_util.Errno, data gin.H) {
	c.Set("custom_code", err.Code)
	c.HTML(http.StatusOK, template, data)
}
