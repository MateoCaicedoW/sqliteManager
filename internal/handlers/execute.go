package handlers

import (
	"fmt"
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/internal/system/render"
)

func (h Handler) Execute(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	render.SetData("error", nil)

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
	render.SetData("error", nil)
	if err := render.Render(w, "results.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) ShowTables(w http.ResponseWriter, r *http.Request) {
	all, c, err := h.Queryer.ShowTables()
	if err != nil {
		render.SetData("error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("results", all)
	render.SetData("columns", c)
	render.SetData("error", nil)
	if err := render.Render(w, "tables.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h Handler) SelectTable(w http.ResponseWriter, r *http.Request) {
	table := r.URL.Query().Get("table")
	all, c, err := h.Queryer.SelectTable(table)
	if err != nil {
		render.SetData("error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("all", all)

	render.SetData("results", all)
	render.SetData("columns", c)
	render.SetData("error", nil)
	if err := render.Render(w, "results.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) GetColumns(w http.ResponseWriter, r *http.Request) {
	table := r.URL.Query().Get("table")
	columns, err := h.Queryer.GetColumns(table)
	if err != nil {
		render.SetData("error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("table", table)
	render.SetData("columns", columns)
	render.SetData("empty", false)
	render.SetData("error", nil)
	if err := render.Render(w, "columns.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) ClearColumns(w http.ResponseWriter, r *http.Request) {
	render.SetData("table", "")
	render.SetData("empty", true)
	if err := render.Render(w, "columns.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
