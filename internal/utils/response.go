package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

func NewError(err string) *Error {
	return &Error{err}
}

func NewMessage(message string) *Message {
	return &Message{message}
}

func NewErrorResponse(response http.ResponseWriter, statusCode int, errMessage string) {
	response.WriteHeader(statusCode)
	response.Header().Set("Content-Type", "application/json")
	errResponse := NewError(errMessage)
	if err := json.NewEncoder(response).Encode(errResponse); err != nil {
		http.Error(response, "failed to encode error message", http.StatusInternalServerError)
	}
}

func NewMessageResponse(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	response.Header().Set("Content-Type", "application/json")
	messageResponse := NewMessage(message)
	if err := json.NewEncoder(response).Encode(messageResponse); err != nil {
		http.Error(response, "failed to encode message", http.StatusInternalServerError)
	}
}
