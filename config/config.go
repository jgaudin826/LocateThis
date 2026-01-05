package config

import (
	"locate-this/database"
	"locate-this/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	GroupRepository         dbmodel.GroupRepository
	UserRepository          dbmodel.UserRepository
	LocationRepository      dbmodel.LocationRepository
	LocationGroupRepository dbmodel.LocationGroupRepository
}

func New() (*Config, error) {
	config := Config{}

	// initialisation de la conexion a la base de données
	databaseSession, err := gorm.Open(sqlite.Open("LocateThis.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories
	config.GroupRepository = dbmodel.NewGroupRepository(databaseSession)
	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.LocationRepository = dbmodel.NewLocationRepository(databaseSession)
	config.LocationGroupRepository = dbmodel.NewLocationGroupRepository(databaseSession)
	
	return &config, nil
}
