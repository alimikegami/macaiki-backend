package response

import "net/http"

type meta struct {
	Message string
	Code    int
}
type baseResponse struct {
	Meta meta
	Data interface{}
}

func SuccessResponse(data interface{}) (int, baseResponse) {
	meta := meta{
		Message: "OK",
		Code:    http.StatusOK,
	}
	return http.StatusOK, baseResponse{
		Meta: meta,
		Data: data,
	}
}

func ErrorResponse(err error, code int) (int, baseResponse) {
	meta := meta{
		Message: err.Error(),
		Code:    code,
	}
	return code, baseResponse{
		Meta: meta,
	}
}
