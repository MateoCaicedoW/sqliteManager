package manager

import (
	"database/sql"
	"log/slog"
	"net/http"
	"path"
	"strings"
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

	f.HandleFunc("POST /sign-in/{$}", h.SignIn)
	f.HandleFunc("DELETE /logout/{$}", h.Logout)
	f.HandleFunc("GET /{$}", h.Index)
	f.HandleFunc("POST /{$}", h.Execute)
	f.HandleFunc("POST /show-tables/{$}", h.ShowTables)
	f.HandleFunc("POST /details/{$}", h.SelectTable)
	f.HandleFunc("POST /columns/{$}", h.GetColumns)
	f.HandleFunc("DELETE /clear-columns/{$}", h.ClearColumns)

	return f
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

// normalizePattern ensures the pattern correctly includes the prefix
func (m *manager) normalizePattern(pattern string) string {
	parts := strings.SplitN(pattern, " ", 2) // Ensure at most 2 parts

	if len(parts) == 2 {
		// If the pattern includes a method (e.g., "GET /path")
		return parts[0] + " " + path.Join("/", m.prefix, parts[1])
	}
	return path.Join("/", m.prefix, parts[0])
}

func (m *manager) HandleFunc(pattern string, handler http.HandlerFunc) {
	// fmt.Println(m.normalizePattern(pattern), handler)
	m.mux.HandleFunc(m.normalizePattern(pattern), handler)
}

func (m *manager) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(m.normalizePattern(pattern), handler)
}
