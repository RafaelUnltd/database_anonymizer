package repositories

import (
	"database_anonymizer/app/src/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.RepositoriesInterface {
	return Repository{
		db: db,
	}
}

func (r Repository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
