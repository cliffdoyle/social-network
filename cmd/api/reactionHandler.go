package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cliffdoyle/social-network/internal/service"
)

type ReactionHandler struct {
	Service *service.ReactionService
}

func NewReactionHandler(s *service.ReactionService) *ReactionHandler {
	return &ReactionHandler{Service: s}
}

type reactionRequest struct {
	Type string `json:"type"`
}

// POST /api/posts/{id}/reactions or /api/comments/{id}/reactions
func (h *ReactionHandler) ReactToPost(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Query().Get("post_id")
	userIDStr := r.Header.Get("X-User-ID") // Or get from context/session
	if postIDStr == "" || userIDStr == "" {
		http.Error(w, "missing post_id or user_id", http.StatusBadRequest)
		return
	}
	postID, _ := strconv.Atoi(postIDStr)
	userID, _ := strconv.Atoi(userIDStr)

	var req reactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.Service.React(userID, &postID, nil, req.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ReactionHandler) ReactToComment(w http.ResponseWriter, r *http.Request) {
	commentIDStr := r.URL.Query().Get("comment_id")
	userIDStr := r.Header.Get("X-User-ID") // Or get from context/session
	if commentIDStr == "" || userIDStr == "" {
		http.Error(w, "missing comment_id or user_id", http.StatusBadRequest)
		return
	}
	commentID, _ := strconv.Atoi(commentIDStr)
	userID, _ := strconv.Atoi(userIDStr)

	var req reactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.Service.React(userID, nil, &commentID, req.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
