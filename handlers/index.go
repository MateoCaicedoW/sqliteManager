package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/render"
)

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderWithLayout(w, "handlers/files.html", "base.html")
}
