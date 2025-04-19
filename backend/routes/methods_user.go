package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/go-playground/validator/v10"
)

// @Summary Create user post
// @Description Creates a post for a user. Can be set to public or private with IsPublic
// @Tags User
// @Accept json
// @Produce json
// @Param request body types.UserPostCreatePayload true "User Post Creation Payload"
// @Success 201 {string} string "Post created successfully"
// @Failure 400 {string} string "Error describing failure - including duplicate attendance posts"
// @Failure 500 {string} string "Internal server error"
// @Router /userpost [post]
func (h *Handler) UserPostOnPost(w http.ResponseWriter, r *http.Request) {
	var payload types.UserPostCreatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check for duplicate ATTENDED and REVIEW posts
	if payload.Type == "ATTENDED" || payload.Type == "REVIEW" {
		exists, err := h.Store.UserPostExists(payload.AuthorID, payload.ConcertID, "ATTENDED")
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if exists {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("you have already marked this concert as attended"))
			return
		}
	}

	_, err := h.Store.CreateUserPost(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

// @Summary Get posts for user dashboard
// @Description Gets public posts from a user's followed network, sorted with most recent first
// @Tags User
// @Produce json
// @Param userID query string true "ID of logged in user"
// @Param p query string false "page number of posts (sets of 20)"
// @Success 200 {object} types.UserPostGetResponse "Activity from user's followed network"
// @Failure 400 {string} string "Error describing failure"
// @Failure 500 {string} string "Internal server error"
// @Router /userpost [get]
func (h *Handler) UserPostOnGet(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("userID")
	if userIDString == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("userID not provided"))
		return
	}

	pageNumberString := r.URL.Query().Get("p")
	pageNumber, err := strconv.ParseInt(pageNumberString, 10, 64)
	if err != nil {
		pageNumber = 0
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad UserPostID provided %v", userIDString))
		return
	}

	posts, err := h.Store.GetActivityFeed(userID, pageNumber)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}

// @Summary Like or unlike a post
// @Description Toggles whether a user likes a given post
// @Tags User
// @Accept json
// @Produce json
// @Param like body types.UserLikePostPayload true "Like Toggle Payload"
// @Success 200 {string} string "Like status toggled successfully"
// @Failure 400 {string} string "Error describing failure"
// @Failure 500 {string} string "Internal server error"
// @Router /like [post]
func (h *Handler) UserLikeOnPost(w http.ResponseWriter, r *http.Request) {
	var payload types.UserLikePostPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.Store.ToggleUserLike(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

// @Summary Get number of likes
// @Description Returns the number of likes for a specific post
// @Tags User
// @Accept json
// @Produce json
// @Param userPostID query string true "Get number of likes"
// @Success 200 {object} types.UserLikeGetResponse
// @Failure 400 {string} string "Error describing failure"
// @Failure 500 {string} string "Internal server error"
// @Router /like [get]
func (h *Handler) UserLikeOnGet(w http.ResponseWriter, r *http.Request) {
	userPostIDString := r.URL.Query().Get("userPostID")
	if userPostIDString == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("userPostID not provided"))
		return
	}

	userPostID, err := strconv.ParseInt(userPostIDString, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("bad UserPostID provided %v", userPostIDString))
		return
	}

	count, err := h.Store.GetNumberOfLikes(userPostID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.UserLikeGetResponse{Count: count})
}
