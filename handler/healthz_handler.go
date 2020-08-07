package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HealthzHandler handle /healthz routes
type HealthzHandler struct{}

// Index is a handler function for GET /healthz
func (h *HealthzHandler) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res := MessageResponse{
		Message: "OK!",
		Meta: ResponseMeta{
			HTTPStatus: http.StatusOK,
		},
	}
	WriteResponse(w, http.StatusOK, res)
}
