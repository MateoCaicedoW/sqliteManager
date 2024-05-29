package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/render"
	"github.com/MateoCaicedoW/sqliteManager/session"
)

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	user := session.GetValue(r, "user")
	if user == nil {
		if err := render.Render(w, "login.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := render.Render(w, "base.html", "files.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
