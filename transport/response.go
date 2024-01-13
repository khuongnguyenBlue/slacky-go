package transport

type ErrorBody struct {
	Message string `json:"message"`
	// Code int `json:"code"`
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