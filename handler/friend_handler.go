package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andrysds/go-clean-architecture/entity"
	"github.com/andrysds/go-clean-architecture/service"
	"github.com/julienschmidt/httprouter"
)

// FriendHandler handle /friends routes
type FriendHandler struct {
	FriendService service.FriendUseCase
}

// Index is a handler function for GET /friends
func (h *FriendHandler) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	friends, err := h.FriendService.GetFriends()
	if err != nil {
		WriteInternalServerErrorResponse(w)
		return
	}

	res := DataResponse{
		Data: friends,
		Meta: ResponseMeta{
			HTTPStatus: http.StatusOK,
		},
	}
	WriteResponse(w, http.StatusOK, res)
}

// Create is a handler function for POST /friends
func (h *FriendHandler) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var friend *entity.Friend
	err := json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		WriteBadRequestResponse(w)
		return
	}

	created, err := h.FriendService.CreateFriend(friend)
	if err != nil {
		WriteInternalServerErrorResponse(w)
		return
	}

	res := DataResponse{
		Data: created,
		Meta: ResponseMeta{
			HTTPStatus: http.StatusCreated,
		},
	}
	WriteResponse(w, http.StatusCreated, res)
}

// Update is a handler function for PUT /friends/:id
func (h *FriendHandler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var friend *entity.Friend
	err := json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		WriteBadRequestResponse(w)
		return
	}

	updated, err := h.FriendService.UpdateFriend(friend)
	if err != nil {
		WriteInternalServerErrorResponse(w)
		return
	}

	res := DataResponse{
		Data: updated,
		Meta: ResponseMeta{
			HTTPStatus: http.StatusAccepted,
		},
	}
	WriteResponse(w, http.StatusAccepted, res)
}

// Delete is a handler function for DELETE /friends/:id
func (h *FriendHandler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		WriteBadRequestResponse(w)
		return
	}

	err = h.FriendService.DeleteFriend(id)
	if err != nil {
		WriteInternalServerErrorResponse(w)
		return
	}

	res := MessageResponse{
		Message: "Deleted!",
		Meta: ResponseMeta{
			HTTPStatus: http.StatusAccepted,
		},
	}
	WriteResponse(w, http.StatusAccepted, res)
}
