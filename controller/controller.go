package controller

import (
	"github.com/gin-gonic/gin"
	"myGPT/errHandler"
	"myGPT/logic"
	"myGPT/model"
)

func RequestHandler(c *gin.Context) {
	var request *model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		errHandler.ResponseError(c, "Invalid JSON")
		return
	}

	prompt := request.Prompt
	res, err := logic.Generate(prompt)
	if err != nil {
		errHandler.ResponseError(c, "Generate error")
		return
	}

	errHandler.ResponseSuccess(c, res)
}
