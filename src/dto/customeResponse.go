package dto

// CustomResponse represents a custom response structure
type CustomResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewCustomResponse creates a new CustomResponse instance
func NewCustomResponse(status, message string, data interface{}, err string) *CustomResponse {
	return &CustomResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}

// SuccessResponse creates a success response
func SuccessResponse(message string, data interface{}) *CustomResponse {
	return &CustomResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(message string, err error) *CustomResponse {
	return &CustomResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	}
}

// NotFoundResponse creates a not found response
func NotFoundResponse(message string) *CustomResponse {
	return &CustomResponse{
		Status:  "not_found",
		Message: message,
	}
}

// BadRequestResponse creates a bad request response
func BadRequestResponse(message string) *CustomResponse {
	return &CustomResponse{
		Status:  "bad_request",
		Message: message,
	}
}
