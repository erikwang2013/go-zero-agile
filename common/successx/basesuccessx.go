package successx

const defaultCode = 0
const defaultMsg = "成功"

type CodeSuccess struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

type CodeSuccessResponse struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func NewCodeSuccess(code int, msg string, data interface{}) *CodeSuccess {
    return &CodeSuccess{Code: code, Msg: msg, Data: data}
}

func NewDefaultSuccess(data interface{}) *CodeSuccess {
    return NewCodeSuccess(defaultCode, defaultMsg, data)
}
