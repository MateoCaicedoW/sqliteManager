package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	// remove this line
	"github.com/MateoCaicedoW/sqliteManager/internal/manager"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := http.NewServeMux()

	// You should open the connection to the database before creating the manager
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fs := manager.New(
		manager.WithConnection(db),
		manager.WithPrefix("/files"),
	)

	s.Handle("/", fs)

	fmt.Println("Server running on port 3000")
	err = http.ListenAndServe(":3000", s)
	if err != nil {
		fmt.Println(err)
	}
}
