package response

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// ErrorResponse represents the structure of the error response body
type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// RespondWithError sends a JSON response with the given status code and message
func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	slog.LogAttrs(
		context.Background(),
		slog.LevelError,
		"error",
		slog.Int("status_code", statusCode),
		slog.String("message", err.Error()),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}
