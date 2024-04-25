package errHandler

import (
	"github.com/gin-gonic/gin"
	"myGPT/model"
	"net/http"
)

func ResponseError(c *gin.Context, msg string) {
	re := &model.ResponseContent{
		Code:    400,
		Msg:     msg,
		Content: nil,
	}
	c.JSON(http.StatusOK, re)
}

func ResponseSuccess(c *gin.Context, content interface{}) {
	re := &model.ResponseContent{
		Code:    200,
		Msg:     "成功",
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}
