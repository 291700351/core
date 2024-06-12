package data

type JsonResponse[T any] struct {
	Code int    `json:"code" yaml:"code"`
	Msg  string `json:"msg" yaml:"msg"`
	Data T      `json:"data" yaml:"data"`
}

func NewJsonResponse[T any](code int, msg string, data T) *JsonResponse[T] {
	return &JsonResponse[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok[T any](data T) *JsonResponse[T] {
	return NewJsonResponse(SysCode.Success.Code, SysCode.Success.Msg, data)
}

func Fail(code ErrCode) *JsonResponse[any] {
	return NewJsonResponse[any](code.Code, code.Msg, nil)

}
