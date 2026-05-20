package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MigrarTablaUsuario(db *gorm.DB) {

	db.AutoMigrate(&Usuario{})
}
