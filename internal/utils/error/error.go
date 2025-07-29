package error

type ErrorResponse struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

const (
	LevelBackend  = "backend"
	LevelDatabase = "database"
)

const (
	MessageInvalidJSONRequest = "invalid json request"
	MessageBadRequest         = "bad request"
	MessageBadURL             = "bad url"
	MessageInvalidMethod      = "invalid method"
)

func NewErrorResponse(level string, message string) ErrorResponse {
	return ErrorResponse{
		Level:   level,
		Message: message,
	}
}