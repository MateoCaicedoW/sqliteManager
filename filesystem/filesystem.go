package filesystem

import (
	"net/http"

	"github.com/MateoCaicedoW/file_system/handlers"
)

type fileSystem struct {
	mux     *http.ServeMux
	prefix  string
	iconURL string
}

func New(options ...option) *fileSystem {
	f := &fileSystem{
		mux: http.NewServeMux(),
	}

	for _, opt := range options {
		opt(f)
	}

	f.HandleFunc("/", handlers.Index)

	return f
}

func (f *fileSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mux.ServeHTTP(w, r)
}

func (f *fileSystem) Handle(pattern string, handler http.Handler) {
	f.mux.Handle(f.prefix+pattern, handler)
}

func (f *fileSystem) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	f.mux.HandleFunc(f.prefix+pattern, handler)
}
