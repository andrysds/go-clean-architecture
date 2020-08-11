package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrysds/go-clean-architecture/test/responsemock"
	"github.com/andrysds/go-clean-architecture/test/servicemock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var errFromFriendService = errors.New("error from FriendService")

func TestFriendsHandlerIndex(t *testing.T) {
	cases := []struct {
		friendServiceErr error
		statusCode       int
		response         string
	}{
		// context: got error from FriendService
		{
			friendServiceErr: errFromFriendService,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendServiceErr: nil,
			statusCode:       http.StatusOK,
			response:         responsemock.JSON("/friends/GET/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{
			FriendService: &servicemock.FriendServiceMock{
				Err: c.friendServiceErr,
			},
		}
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/friends", nil)

		router := httprouter.New()
		router.GET("/friends", handler.Index)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}

func TestFriendsHandlerCreate(t *testing.T) {
	cases := []struct {
		friendServiceErr error
		statusCode       int
		response         string
	}{
		// context: got error from FriendService
		{
			friendServiceErr: errFromFriendService,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendServiceErr: nil,
			statusCode:       http.StatusCreated,
			response:         responsemock.JSON("/friends/POST/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{
			FriendService: &servicemock.FriendServiceMock{
				Err: c.friendServiceErr,
			},
		}
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/friends", nil)

		router := httprouter.New()
		router.POST("/friends", handler.Create)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}

func TestFriendsHandlerUpdate(t *testing.T) {
	cases := []struct {
		friendServiceErr error
		statusCode       int
		response         string
	}{
		// context: got error from FriendService
		{
			friendServiceErr: errFromFriendService,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendServiceErr: nil,
			statusCode:       http.StatusAccepted,
			response:         responsemock.JSON("/friends/PUT/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{
			FriendService: &servicemock.FriendServiceMock{
				Err: c.friendServiceErr,
			},
		}
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/friends/1", nil)

		router := httprouter.New()
		router.PUT("/friends/:id", handler.Update)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}

func TestFriendsHandlerDelete(t *testing.T) {
	cases := []struct {
		friendServiceErr error
		statusCode       int
		response         string
	}{
		// context: got error from FriendService
		{
			friendServiceErr: errFromFriendService,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendServiceErr: nil,
			statusCode:       http.StatusAccepted,
			response:         responsemock.JSON("/friends/DELETE/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{
			FriendService: &servicemock.FriendServiceMock{
				Err: c.friendServiceErr,
			},
		}
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/friends/1", nil)

		router := httprouter.New()
		router.DELETE("/friends/:id", handler.Delete)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}
