package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	responsemock "github.com/andrysds/go-clean-architecture/test/response_mock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var errFromFriendUseCase = errors.New("error from friendUseCase")

func TestFriendsHandlerIndex(t *testing.T) {
	cases := []struct {
		friendUseCaseErr error
		statusCode       int
		response         string
	}{
		// context: got error from friendUseCase
		{
			friendUseCaseErr: errFromFriendUseCase,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendUseCaseErr: nil,
			statusCode:       http.StatusOK,
			response:         responsemock.JSON("/friends/GET/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{friendUseCase: &friendUseCaseMock{err: c.friendUseCaseErr}}
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
		friendUseCaseErr error
		statusCode       int
		response         string
	}{
		// context: got error from friendUseCase
		{
			friendUseCaseErr: errFromFriendUseCase,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendUseCaseErr: nil,
			statusCode:       http.StatusCreated,
			response:         responsemock.JSON("/friends/POST/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{friendUseCase: &friendUseCaseMock{err: c.friendUseCaseErr}}
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
		friendUseCaseErr error
		statusCode       int
		response         string
	}{
		// context: got error from friendUseCase
		{
			friendUseCaseErr: errFromFriendUseCase,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendUseCaseErr: nil,
			statusCode:       http.StatusAccepted,
			response:         responsemock.JSON("/friends/PUT/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{friendUseCase: &friendUseCaseMock{err: c.friendUseCaseErr}}
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
		friendUseCaseErr error
		statusCode       int
		response         string
	}{
		// context: got error from friendUseCase
		{
			friendUseCaseErr: errFromFriendUseCase,
			statusCode:       http.StatusInternalServerError,
			response:         responsemock.JSON("/others/internal_server_error.json"),
		},
		// context: positive case
		{
			friendUseCaseErr: nil,
			statusCode:       http.StatusAccepted,
			response:         responsemock.JSON("/friends/DELETE/response.json"),
		},
	}

	for _, c := range cases {
		handler := &FriendHandler{friendUseCase: &friendUseCaseMock{err: c.friendUseCaseErr}}
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/friends/1", nil)

		router := httprouter.New()
		router.DELETE("/friends/:id", handler.Delete)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, c.statusCode, recorder.Code)
		assert.JSONEq(t, c.response, recorder.Body.String())
	}
}

type friendUseCaseMock struct{ err error }

func (m *friendUseCaseMock) GetFriends() ([]interface{}, error) {
	if m.err != nil {
		return []interface{}{}, m.err
	}
	return []interface{}{"andrys", "budi", "cecep"}, nil
}

func (m *friendUseCaseMock) CreateFriend(httprouter.Params) (interface{}, error) {
	if m.err != nil {
		return nil, m.err
	}
	return "andrys", nil
}

func (m *friendUseCaseMock) UpdateFriend(httprouter.Params) (interface{}, error) {
	if m.err != nil {
		return nil, m.err
	}
	return "andrys", nil
}
func (m *friendUseCaseMock) DeleteFriend(httprouter.Params) error {
	return m.err
}
