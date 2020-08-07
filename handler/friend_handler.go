package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// FriendHandler handle /friends routes
type FriendHandler struct {
	friendUseCase interface {
		GetFriends() ([]interface{}, error)
		CreateFriend(httprouter.Params) (interface{}, error)
		UpdateFriend(httprouter.Params) (interface{}, error)
		DeleteFriend(httprouter.Params) error
	}
}

// Index is a handler function for GET /friends
func (h *FriendHandler) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	friends, err := h.friendUseCase.GetFriends()
	if err != nil {
		WriteInternalServerError(w)
		return
	}

	res := map[string]interface{}{
		"data": friends,
		"meta": map[string]interface{}{
			"http_status": http.StatusOK,
		},
	}
	WriteResponse(w, http.StatusOK, res)
}

// Create is a handler function for POST /friends
func (h *FriendHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	friend, err := h.friendUseCase.CreateFriend(ps)
	if err != nil {
		WriteInternalServerError(w)
		return
	}

	res := map[string]interface{}{
		"data": friend,
		"meta": map[string]interface{}{
			"http_status": http.StatusCreated,
		},
	}
	WriteResponse(w, http.StatusCreated, res)
}

// Update is a handler function for PUT /friends/:id
func (h *FriendHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	friend, err := h.friendUseCase.UpdateFriend(ps)
	if err != nil {
		WriteInternalServerError(w)
		return
	}

	res := map[string]interface{}{
		"data": friend,
		"meta": map[string]interface{}{
			"http_status": http.StatusAccepted,
		},
	}
	WriteResponse(w, http.StatusAccepted, res)
}

// Delete is a handler function for DELETE /friends/:id
func (h *FriendHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := h.friendUseCase.DeleteFriend(ps)
	if err != nil {
		WriteInternalServerError(w)
		return
	}

	res := map[string]interface{}{
		"message": "Deleted!",
		"meta": map[string]interface{}{
			"http_status": http.StatusAccepted,
		},
	}
	WriteResponse(w, http.StatusAccepted, res)
}
