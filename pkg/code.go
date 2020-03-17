package pkg

// 错误码
const (
	Success          = 200
	Error            = 500
	InvalidParam     = 400
)

var MsgFlags = map[int]string{
	Success:          "ok",
	Error:            "fail",
	InvalidParam:     "请求参数错误",
}

func getMsg(code int) string {
	return MsgFlags[code]
}
