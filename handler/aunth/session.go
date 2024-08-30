package aunth

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func setSession(w http.ResponseWriter, r *http.Request) {

	session, _ := cs.Get(r, "session-name")
	session.Values["key"] = true
	session.Values["value"] = "FujisueKouki"
	err := session.Save(r, w)
	if err != nil {
		w.Write([]byte("false"))
		return
	}
	w.Write([]byte("true"))
	return

}

func getSession(w http.ResponseWriter, r *http.Request) {

	session, _ := cs.Get(r, "session-name")
	value, _ := session.Values["key"].(bool)
	f, _ := session.Values["value"].(string)
	if value {
		w.Write([]byte("true"))
		w.Write([]byte(f))
	} else {
		w.Write([]byte("false"))
	}
	return
}
