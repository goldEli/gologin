package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailWithMessage(message string, ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    ResponseCodeBadRequest,
		Message: message,
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, obj)
}

func FailWithCode(code int, ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    code,
		Message: ResponseMessage(code),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, obj)
}

func Fail(ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    ResponseCodeBadRequest,
		Message: ResponseMessage(ResponseCodeBadRequest),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, obj)
}

func FailServer(ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    ResponseCodeInternalServerError,
		Message: ResponseMessage(ResponseCodeInternalServerError),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, obj)
}

func Ok(ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    ResponseCodeOk,
		Message: ResponseMessage(ResponseCodeOk),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, obj)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	obj := &BaseResponse{
		Code:    ResponseCodeOk,
		Message: ResponseMessage(ResponseCodeOk),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, obj)
}

func ResponseOk() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeOk,
		Message: ResponseMessage(ResponseCodeOk),
		Data:    nil,
	}
}

func ResponseOkWIthData(data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeOk,
		Message: ResponseMessage(ResponseCodeOk),
		Data:    data,
	}
}

func ResponseError() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeBadRequest,
		Message: ResponseMessage(ResponseCodeBadRequest),
		Data:    nil,
	}
}

func ResponseUserPasswordError() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeUserPasswordError,
		Message: ResponseMessage(ResponseCodeUserPasswordError),
		Data:    nil,
	}
}

func ResponseUnauthorizedError() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeUnauthorized,
		Message: ResponseMessage(ResponseCodeUnauthorized),
		Data:    nil,
	}
}

func ResponseServerError() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeInternalServerError,
		Message: ResponseMessage(ResponseCodeInternalServerError),
		Data:    nil,
	}
}

func ResponseNoData(code int) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: ResponseMessage(code),
		Data:    nil,
	}
}

func ResponseWIthData(code int, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: ResponseMessage(code),
		Data:    data,
	}
}
