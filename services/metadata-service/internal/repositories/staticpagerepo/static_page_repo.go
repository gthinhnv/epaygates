package staticpagerepo

import (
	"context"
	"metadatasvc/gen/go/staticpagepb"
	"shared/models/staticpagemodel"

	"github.com/jmoiron/sqlx"
)

// Repository defines all database operations for StaticPage.
type Repository interface {
	// Create inserts a new static page and returns its ID.
	Create(ctx context.Context, page *staticpagemodel.StaticPage) (uint64, error)

	// Update updates fields using a field mask.
	Update(ctx context.Context, id uint64, req *staticpagemodel.StaticPage, fields []string) error

	// Delete soft-deletes one or multiple pages.
	Delete(ctx context.Context, ids []uint64, deletedBy uint64) error

	// GetByID retrieves a static page by ID.
	GetByID(ctx context.Context, id uint64) (*staticpagemodel.StaticPage, error)

	// GetBySlug retrieves a static page by slug.
	GetBySlug(ctx context.Context, slug string) (*staticpagemodel.StaticPage, error)

	// List returns pages with filters, sorting, and pagination.
	List(ctx context.Context, req *staticpagepb.ListRequest) ([]*staticpagemodel.StaticPage, error)

	// Count returns total items for the given filter (used for pagination).
	Count(ctx context.Context, req *staticpagepb.ListRequest) (uint64, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return NewMysqlRepository(db)
}
