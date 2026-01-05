package config

import (
	"locate-this/database"
	"locate-this/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Constants struct {
}

type Config struct {
	GroupEntryRepository         dbmodel.GroupRepository
	UserEntryRepository          dbmodel.UserRepository
	LocationEntryRepository      dbmodel.LocationRepository
	LocationGroupEntryRepository dbmodel.LocationGroupRepository
	// Constants
	SecretJWT     string
	SecretRefreshJWT string
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
	config.GroupEntryRepository = dbmodel.NewGroupRepository(databaseSession)
	config.UserEntryRepository = dbmodel.NewUserRepository(databaseSession)
	config.LocationEntryRepository = dbmodel.NewLocationRepository(databaseSession)
	config.LocationGroupEntryRepository = dbmodel.NewLocationGroupRepository(databaseSession)

	config.SecretJWT = "HeXIVX"
	config.SecretRefreshJWT = "XVIXeH"

	return &config, nil
}
