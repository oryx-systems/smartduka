package db

import (
	"log"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

// DbServiceImpl is an implementation of the database repository
// It is implementation agnostic i.e logic should be handled using
// the preferred database
type DbServiceImpl struct {
	create gorm.Create
	query  gorm.Query
	update gorm.Update
}

// NewDBService creates a new database service
func NewDBService(c gorm.Create, q gorm.Query, u gorm.Update) *DbServiceImpl {
	environment := helpers.MustGetEnvVar("REPOSITORY")

	switch environment {
	case "firebase":
		return nil

	case "postgres":
		return &DbServiceImpl{
			create: c,
			query:  q,
			update: u,
		}

	default:
		log.Panicf("unknown repository: %s", environment)

	}

	return nil
}
