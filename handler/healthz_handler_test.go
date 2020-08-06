package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHealthzHandlerIndex(t *testing.T) {
	handler := &HealthzHandler{}
	request, _ := http.NewRequest("GET", "http://localhost:8080/healthz", nil)
	recorder := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/healthz", handler.Index)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, "OK!\n", recorder.Body.String())
}
