package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MateoCaicedoW/sqliteManager/manager"
	"github.com/jmoiron/sqlx"

	// remove this line
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := http.NewServeMux()

	// You should open the connection to the database before creating the manager
	db, err := sqlx.Open("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fs := manager.New(
		manager.WithPrefix("/files"),
		manager.WithConnection(db),
	)

	s.Handle("/", fs)

	fmt.Println("Server running on port 3000")
	err = http.ListenAndServe(":3000", s)
	if err != nil {
		fmt.Println(err)
	}
}
