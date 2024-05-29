package main

import (
	"fmt"
	"net/http"

	"github.com/MateoCaicedoW/sqliteManager/manager"
)

func main() {
	s := http.NewServeMux()

	fs := manager.New(
		manager.WithPrefix("/files"),
	)

	s.Handle("/", fs)

	fmt.Println("Server running on port 3000")
	err := http.ListenAndServe(":3000", s)
	if err != nil {
		fmt.Println(err)
	}
}
