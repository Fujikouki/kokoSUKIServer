package account

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"net/http"
)

type AccountRequest struct {
	Email    string
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

	err := h.accountUseCase.CreateAccount(ctx, req.Email, req.Username, req.Password, req.IconUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.cs.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Failed to get session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["key"] = true
	session.Values["value"] = "FujisueKouki"

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,  // 1時間
		HttpOnly: false, // クライアント側のJavaScriptからアクセスを防ぐ
		Secure:   false, // HTTPS接続でのみクッキーが送信される

	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("true"))
	if err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("true"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
