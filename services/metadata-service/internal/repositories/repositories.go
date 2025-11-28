package repositories

import (
	"metadatasvc/internal/db"
	"metadatasvc/internal/repositories/staticpagerepo"

	"github.com/jmoiron/sqlx"
)

// Repositories is a high-level wrapper around the database.
type Repositories struct {
	db *sqlx.DB

	StaticPageRepo staticpagerepo.Repository
}

// NewRepositories initializes all repositories.
func NewRepositories(db *db.DB) *Repositories {
	conn := db.Conn()

	return &Repositories{
		db: conn,

		StaticPageRepo: staticpagerepo.NewRepository(conn),
	}
}
