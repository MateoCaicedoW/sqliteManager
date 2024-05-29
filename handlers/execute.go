package handlers

import (
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/render"
	"github.com/MateoCaicedoW/sqliteManager/session"
)

func (h Handler) Execute(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	render.SetData("error", nil)
	user := session.GetValue(r, "user")
	if user == nil {
		if err := render.Render(w, "login.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if query == "" {
		render.SetData("error", "Query cannot be empty")
		if err := render.Render(w, "results.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	all, c, err := h.Queryer.Query(query)
	if err != nil {
		render.SetData("error", err.Error())
		if err := render.Render(w, "results.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	render.SetData("results", all)
	render.SetData("columns", c)
	if err := render.Render(w, "results.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
