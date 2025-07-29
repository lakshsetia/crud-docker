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
	MessageBadRequest = "bad request"
	MessageBadURL     = "bad url"
)

func NewErrorResponse(level string, message string) ErrorResponse {
	return ErrorResponse{
		Level:   level,
		Message: message,
	}
}