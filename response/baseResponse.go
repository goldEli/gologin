package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func ResponseServerError() *BaseResponse {
	return &BaseResponse{
		Code:    ResponseCodeInternalServerError,
		Message: ResponseMessage(ResponseCodeInternalServerError),
		Data:    nil,
	}
}
