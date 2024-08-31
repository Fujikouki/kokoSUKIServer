package account

import (
	"encoding/json"
	"net/http"
)

type AccountRequest struct {
	Username string
	Password string
	IconUrl  string
}

func (h *handler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var req AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := h.accountUsecase.CreateAccount(ctx, req.Username, req.Password, req.IconUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("true"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
