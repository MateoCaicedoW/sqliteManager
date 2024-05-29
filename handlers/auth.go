package handlers

import (
	"net/http"
	"os"

	"github.com/MateoCaicedoW/sqliteManager/render"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	user := render.GetData("user")
	if user != nil {
		if err := render.Render(w, "base.html", "files.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := render.Render(w, "login.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	password := r.FormValue("password")

	render.SetData("error", nil)
	if user == "" || password == "" {
		render.SetData("error", "User and password are required")
		if err := render.Render(w, "login.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	realUser := os.Getenv("MANAGER_USER")
	realPassword := os.Getenv("MANAGER_PASSWORD")

	if user != realUser || password != realPassword {
		render.SetData("error", "User or password are incorrect")
		if err := render.Render(w, "login.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	render.SetData("user", user)
	if err := render.Render(w, "base.html", "files.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Logout(w http.ResponseWriter, r *http.Request) {
	render.SetData("user", nil)
	render.SetData("error", nil)
	if err := render.Render(w, "login.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
