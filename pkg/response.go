package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  getMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) InvalidParam(err error) {
	errs := err.(validator.ValidationErrors)
	g.C.JSON(200, Response{
		Code: InvalidParam,
		Msg:  getParamErrMsg(errs),
		Data: nil,
	})
}


func getParamErrMsg(vErr validator.ValidationErrors) (errMsg string) {
	for _, errStr := range vErr.Translate(Trans) {
		errMsg += fmt.Sprintln(errStr)
	}
	return
}
