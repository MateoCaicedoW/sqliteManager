package session

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// getSession returns a session
func getSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session")
	return session
}

// SetValue sets a session
func SetValue(w http.ResponseWriter, r *http.Request, key string, value interface{}) {
	session := getSession(r)
	session.Values[key] = value
	session.Save(r, w)
}

// GetValue returns a session value
func GetValue(r *http.Request, key string) interface{} {
	session := getSession(r)
	return session.Values[key]
}

// DeleteValue deletes a session value
func DeleteValue(w http.ResponseWriter, r *http.Request, key string) {
	session := getSession(r)
	delete(session.Values, key)
	session.Save(r, w)
}

// Clear deletes a session
func Clear(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	for key := range session.Values {
		delete(session.Values, key)
	}
	session.Save(r, w)
}
