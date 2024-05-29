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
		render.HTML(w, "handlers/results.html")
		return
	}

	all, c, err := h.Queryer.Query(query)
	if err != nil {
		render.SetData("error", err.Error())
		render.HTML(w, "handlers/results.html")
		return
	}

	render.SetData("results", all)
	render.SetData("columns", c)
	render.HTML(w, "handlers/results.html")

}
