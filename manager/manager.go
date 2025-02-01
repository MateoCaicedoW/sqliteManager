package manager

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/MateoCaicedoW/sqliteManager/internal/connection"
	"github.com/MateoCaicedoW/sqliteManager/internal/handlers"
	_ "github.com/MateoCaicedoW/sqliteManager/internal/system/envload"
	"github.com/MateoCaicedoW/sqliteManager/internal/system/render"
)

type manager struct {
	mux     *http.ServeMux
	prefix  string
	iconURL string
	db      *sql.DB
}

func New(options ...option) http.Handler {
	f := &manager{
		mux: http.NewServeMux(),
	}

	for _, opt := range options {
		opt(f)
	}

	h := handlers.Handler{
		Queryer: connection.New(f.db),
		Prefix:  f.prefix,
	}

	render.SetData("iconURL", f.iconURL)
	render.SetData("prefix", f.prefix)

	f.register(
		route{method: "POST", pattern: "/sign-in/{$}", handler: h.SignIn},
		route{method: "DELETE", pattern: "/logout/{$}", handler: h.Logout},
		route{method: "GET", pattern: "/{$}", handler: h.Index},
		route{method: "POST", pattern: "/{$}", handler: h.Execute},
		route{method: "POST", pattern: "/show-tables/{$}", handler: h.ShowTables},
		route{method: "POST", pattern: "/details/{$}", handler: h.SelectTable},
		route{method: "POST", pattern: "/columns/{$}", handler: h.GetColumns},
		route{method: "DELETE", pattern: "/clear-columns/{$}", handler: h.ClearColumns},
	)

	return f.mux
}

type route struct {
	method  string
	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

func (m *manager) register(routes ...route) {
	for _, r := range routes {
		m.mux.HandleFunc(fmt.Sprintf("%s %s%s", r.method, m.prefix, r.pattern), r.handler)
	}
}

func (f *manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Adding common things here, loggers and other things.
	t := time.Now()

	// Parsing form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Correcting method based on _method field
	if r.Method == "POST" && r.FormValue("_method") != "" {
		r.Method = r.FormValue("_method")
	}

	f.mux.ServeHTTP(w, r)
	slog.Info(">", "method", r.Method, "path", r.URL.Path, "duration", time.Since(t))
}
