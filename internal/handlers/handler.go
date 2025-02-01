package handlers

import (
	"github.com/MateoCaicedoW/sqliteManager/internal/connection"
)

type Handler struct {
	Prefix  string
	Queryer connection.Executer
}
