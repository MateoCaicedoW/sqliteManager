package handlers

import (
	"github.com/MateoCaicedoW/sqliteManager/connection"
)

type Handler struct {
	Queryer connection.Executer
}
