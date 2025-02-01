package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/internal/system/render"
)

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	render.SetData("error", nil)
	if err := render.Render(w, "login.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
