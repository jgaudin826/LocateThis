package config


import (
	"locate-this/database/dbmodel" 
	"locate-this/database" 

	"github.com/glebarez/sqlite" 
	"gorm.io/gorm"
)

type Config struct {
	// Connexion aux repositories
}

func New() (*Config, error) {
	config := Config{}

	// initialisation de la conexion a la base de données
	databaseSession, err := gorm.Open(sqlite.Open("vet_clinic_api.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories

	return &config, nil
}