package aunth

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func Middleware(cs *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// セッションを取得
			session, err := cs.Get(r, "session-name")
			if err != nil {
				http.Error(w, "Failed to get session: "+err.Error(), http.StatusInternalServerError)
				println("セッション取得失敗")
				return
			}
			f, ok := session.Values["value"].(string)
			if !ok {
				http.Error(w, "Failed to get session: ", http.StatusInternalServerError)
				println("value取得失敗")
				return
			}
			println(f)
			// 認証が成功したので次のハンドラーに渡す
			next.ServeHTTP(w, r)
		})
	}
}
