package utils

type Result struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() Result {
	return Result{
		Code: 0,
		Msg:  "操作成功",
		Data: nil,
	}
}

func Error(msg string) Result {
	return Result{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
}
