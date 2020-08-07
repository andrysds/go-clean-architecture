package handler

import (
	"encoding/json"
	"net/http"
)

// DataResponse is a response with data and meta property
type DataResponse struct {
	Data interface{}  `json:"data"`
	Meta ResponseMeta `json:"meta"`
}

// MessageResponse is a response with message and meta property
type MessageResponse struct {
	Message string       `json:"message"`
	Meta    ResponseMeta `json:"meta"`
}

// MessageDataResponse is a response with message, data and meta property
type MessageDataResponse struct {
	Message string       `json:"message"`
	Meta    ResponseMeta `json:"meta"`
}

// ResponseMeta is a struct for response meta property
type ResponseMeta struct {
	HTTPStatus int `json:"http_status"`
}

// WriteInternalServerError writes internal server error response to the HTTP connection
func WriteInternalServerError(w http.ResponseWriter) {
	res := MessageResponse{
		Message: "Internal Server Error!",
		Meta: ResponseMeta{
			HTTPStatus: http.StatusInternalServerError,
		},
	}
	WriteResponse(w, http.StatusInternalServerError, res)
}

// WriteResponse writes response to the HTTP connection
func WriteResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	response, err := json.Marshal(body)
	if err != nil {
		WriteInternalServerError(w)
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
