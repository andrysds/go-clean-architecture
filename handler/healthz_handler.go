package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HealthzHandler handle /healthz route
type HealthzHandler struct{}

// Index is a handler function for /healthz route
func (h *HealthzHandler) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "OK!")
}
