package manager

import (
	"net/http"
	"path"
	"strings"

	"github.com/MateoCaicedoW/sqliteManager/connection"
	_ "github.com/MateoCaicedoW/sqliteManager/envload"
	"github.com/MateoCaicedoW/sqliteManager/handlers"
)

type manager struct {
	mux     *http.ServeMux
	prefix  string
	iconURL string
}

func New(options ...option) http.Handler {
	f := &manager{
		mux: http.NewServeMux(),
	}

	for _, opt := range options {
		opt(f)
	}

	h := handlers.Handler{
		Queryer: connection.New(),
	}

	f.HandleFunc("GET /{$}", h.Index)
	f.HandleFunc("POST /{$}", h.Execute)

	return f
}

func (f *manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mux.ServeHTTP(w, r)
}

func (m *manager) HandleFunc(pattern string, handler http.HandlerFunc) {
	// prefix the pattens with the routesPrefix
	parts := strings.Split(pattern, " ")
	pattern = path.Join(m.prefix, parts[0])
	if len(parts) == 2 {
		pattern = path.Join(m.prefix, parts[1])
		pattern = parts[0] + " " + pattern
	}

	// Adding the handler to the ServeMux
	m.mux.HandleFunc(pattern, handler)
}

func (m *manager) Handle(pattern string, handler http.Handler) {
	// prefix the pattens with the routesPrefix
	parts := strings.Split(pattern, " ")
	pattern = path.Join(m.prefix, parts[0])
	if len(parts) == 2 {
		pattern = path.Join(m.prefix, parts[1])
		pattern = parts[0] + " " + pattern
	}

	// Adding the handler to the ServeMux
	m.mux.Handle(pattern, handler)
}