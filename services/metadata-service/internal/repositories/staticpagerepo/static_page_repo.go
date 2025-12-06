package staticpagerepo

import (
	"metadatasvc/gen/go/staticpagepb"
	"shared/models/staticpagemodel"

	"github.com/jmoiron/sqlx"
)

// Repository defines all database operations for StaticPage.
type Repository interface {
	// Create inserts a new static page and returns its ID.
	Create(page *staticpagemodel.StaticPage) (uint64, error)

	// Update updates fields using a field mask.
	Update(id uint64, req *staticpagemodel.StaticPage, fields []string) error

	// Delete soft-deletes one or multiple pages.
	Delete(ids []uint64, deletedBy uint64) error

	// GetByID retrieves a static page by ID.
	GetByID(id uint64) (*staticpagemodel.StaticPage, error)

	// GetBySlug retrieves a static page by slug.
	GetBySlug(slug string) (*staticpagemodel.StaticPage, error)

	// List returns pages with filters, sorting, and pagination.
	List(req *staticpagepb.ListRequest) ([]*staticpagemodel.StaticPage, error)

	// Count returns total items for the given filter (used for pagination).
	Count(req *staticpagepb.ListRequest) (uint64, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return NewMysqlRepository(db)
}
