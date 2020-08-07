package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	responsemock "github.com/andrysds/go-clean-architecture/test/response_mock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestHealthzHandlerIndex(t *testing.T) {
	handler := &HealthzHandler{}
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/healthz", nil)

	router := httprouter.New()
	router.GET("/healthz", handler.Index)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, responsemock.JSON("/healthz/GET/response.json"), recorder.Body.String())
}
