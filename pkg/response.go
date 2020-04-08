package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// RC (response context) 内置context
type RC struct {
	C *gin.Context // 内嵌 gin.context
}

// Response 统一返回结构体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *RC) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  getMsg(errCode),
		Data: data,
	})
	return
}

// InvalidParam 参数无效时返回
func (g *RC) InvalidParam(err error) {
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
