package database

import (
	"log"
	"proyecto_api_escolar/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	//habilitar claves fóraneas en SQLite
	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.Exec("PRAGMA foreign_keys = ON;")
	}

	//Auto-migrar los modelos (crea las tablas si no existen)
	err = DB.AutoMigrate(
		&models.Student{},
		&models.Subject{},
		&models.Grade{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas: %v", err)
	}

	log.Println("Base de datos inicializada y tablas creadas")
}