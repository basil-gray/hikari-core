package activity

import (
    "encoding/json"
    "net/http"
)

type Handler struct {
    Service *Service
}

func NewHandler(s *Service) *Handler {
    return &Handler{Service: s}
}

func (h *Handler) CreateActivity(w http.ResponseWriter, r *http.Request) {
    var activity Activity
    if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createdActivity, err := h.Service.CreateActivity(activity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(createdActivity)
}

// Other handlers (GetActivity, UpdateActivity, DeleteActivity) follow a similar pattern