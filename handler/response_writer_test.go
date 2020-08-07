package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrysds/go-clean-architecture/test/responsemock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestWriteInternalServerError(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/test-write-internal-server-error", nil)

	handleFunc := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		WriteInternalServerErrorResponse(w)
	}

	router := httprouter.New()
	router.GET("/test-write-internal-server-error", handleFunc)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.JSONEq(t, responsemock.JSON("/others/internal_server_error.json"), recorder.Body.String())
}

func TestWriteResponse(t *testing.T) {
	cases := []struct {
		handleFunc func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
		statusCode int
		response   string
	}{
		// context: response body can't be marshaled to JSON
		// this case is probably never gonna happened
		{
			handleFunc: func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				WriteResponse(w, http.StatusOK, make(chan int))
			},
			statusCode: http.StatusInternalServerError,
			response:   responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			handleFunc: func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				res := MessageResponse{
					Message: "OK!",
					Meta: ResponseMeta{
						HTTPStatus: http.StatusOK,
					},
				}
				WriteResponse(w, http.StatusOK, res)
			},
			statusCode: http.StatusOK,
			response:   responsemock.JSON("/healthz/GET/response.json"),
		},
	}

	for _, c := range cases {
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/test-write-response", nil)

		router := httprouter.New()
		router.GET("/test-write-response", c.handleFunc)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}
