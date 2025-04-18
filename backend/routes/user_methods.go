package routes

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/go-playground/validator/v10"
)

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

	userPostTypes := []string{"ATTENDED", "WISHLIST", "REVIEW", "LISTCREATED"}
	if !slices.Contains(userPostTypes, payload.Type) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid UserPost type"))
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

func (h *Handler) UserLikeOnGet(w http.ResponseWriter, r *http.Request) {

}
