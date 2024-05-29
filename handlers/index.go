package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/render"
)

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderWithLayout(w, "base.html", "files.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
