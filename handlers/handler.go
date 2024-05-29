package handlers

import (
	"github.com/MateoCaicedoW/sqliteManager/connection"
)

type Handler struct {
	Prefix  string
	Queryer connection.Executer
}
