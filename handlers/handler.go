package handlers

import (
	"github.com/MateoCaicedoW/file_system/connection"
)

type Handler struct {
	Queryer connection.Executer
}
