package responses

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewBadRequestError(code int, message, field string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}
