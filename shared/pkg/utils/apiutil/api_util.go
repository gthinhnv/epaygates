package apiutil

type RespCode int

const (
	CODE_SUCCESS RespCode = iota
	CODE_ERROR
	CODE_INVALID_PARAMS
	CODE_UNAUTHORIZED
	CODE_FORBIDDEN
	CODE_NOT_FOUND
	CODE_INTERNAL_ERROR
)

type Response struct {
	Code    RespCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
