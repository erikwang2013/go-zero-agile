package errorx

const defaultCode = 1

type CodeError struct {
    Code int      `json:"code"`
    Msg  string   `json:"msg"`
    Data string `json:"data"`
}

type CodeErrorResponse struct {
    Code int      `json:"code"`
    Msg  string   `json:"msg"`
    Data string `json:"data"`
}

func NewCodeError(code int, msg string) error {
    return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
    return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
    return e.Msg
}

func (e *CodeError) Datas() *CodeErrorResponse {
    return &CodeErrorResponse{
        Code: e.Code,
        Msg:  e.Msg,
    }
}
