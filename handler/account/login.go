package account

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {

	var req AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	acc, err := h.accountUsecase.Login(ctx, req.Email, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(acc.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
