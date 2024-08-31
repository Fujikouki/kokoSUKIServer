package account

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"log"
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

	acc, err := h.accountUseCase.Login(ctx, req.Email, req.Password)

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
		SameSite: http.SameSiteLaxMode,
	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for key, value := range session.Values {
		log.Printf("Session key: %s, value: %v", key, value)
	}

	_, err = w.Write([]byte("true"))
	if err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(acc.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
