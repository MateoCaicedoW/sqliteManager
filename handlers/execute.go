package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/render"
)

func (h Handler) Execute(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	render.SetData("error", nil)

	if query == "" {
		render.SetData("error", "Query cannot be empty")
		if err := render.RenderWithLayout(w, "results.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	all, c, err := h.Queryer.Query(query)
	if err != nil {
		render.SetData("error", err.Error())
		if err := render.RenderWithLayout(w, "results.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	render.SetData("results", all)
	render.SetData("columns", c)
	if err := render.RenderWithLayout(w, "results.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
