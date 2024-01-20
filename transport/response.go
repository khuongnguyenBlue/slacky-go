package transport

type ErrorBody struct {
	Message string `json:"message"`
	Code string `json:"code"`
	HttpCode int `json:"-"`
}

func NewErrorBody(message string, code string) ErrorBody {
	return ErrorBody{
		Message: message,
		Code: code,
	}
}

func (e ErrorBody) Error() string {
	return e.Message
}

// define response struct
type BaseResponse struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
	Errors []ErrorBody `json:"errors"`
}

func SuccessResponse(data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Data: data,
		Errors: nil,
	}
}