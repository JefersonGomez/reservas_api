package models

import (
	"gorm.io/gorm"
)

/*
	ID

NombreCliente  → quien reserva
Fecha          → "2026-05-20"
Hora           → "19:00"
NumPersonas    → cuántas personas
Estado         → "pendiente", "confirmada", "cancelada"
*/
type Reserva struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	NombreCliente string `json:"nombreCliente"`
	Fecha         string `json:"fecha"`
	Hora          string `json:"hora"`
	NumPersonas   uint   `json:"numPersonas"`
	Estado        string `json:"estado"`
}

func MigrarTablaReserva(db *gorm.DB) {
	db.AutoMigrate(&Reserva{})
}
