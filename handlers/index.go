package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/file_system/render"
)

func (h Handler) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderWithLayout(w, "handlers/files.html", "base.html")
}
