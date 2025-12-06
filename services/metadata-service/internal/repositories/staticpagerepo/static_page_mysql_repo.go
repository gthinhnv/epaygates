package staticpagerepo

import (
	"context"
	"encoding/json"
	"metadatasvc/gen/go/staticpagepb"
	"shared/models/staticpagemodel"

	"github.com/jmoiron/sqlx"
)

type MysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(db *sqlx.DB) Repository {
	if err := prepareStatements(db); err != nil {
		panic(err)
	}
	return &MysqlRepository{db: db}
}

// --------------------------------------------------
// Create inserts a new static page and returns its ID.
// --------------------------------------------------
func (r *MysqlRepository) Create(ctx context.Context, page *staticpagemodel.StaticPage) (uint64, error) {
	var seoJson []byte
	if page.Seo != nil {
		seoJson, _ = json.Marshal(page.Seo)
	}

	res, err := createStmt.ExecContext(ctx,
		page.Title,
		page.Slug,
		page.Content,
		page.PageType,
		page.SortOrder,
		seoJson,
		page.AdsPlatform,
		page.Status,
		page.CreatedBy,
		page.CreatedBy,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// --------------------------------------------------
// Update updates fields using a field mask.
// --------------------------------------------------
func (r *MysqlRepository) Update(ctx context.Context, id uint64, req *staticpagemodel.StaticPage, fields []string) error {
	return nil
}

// --------------------------------------------------
// Delete soft-deletes one or multiple pages.
// --------------------------------------------------
func (r *MysqlRepository) Delete(ctx context.Context, ids []uint64, deletedBy uint64) error {
	return nil
}

// --------------------------------------------------
// GetByID retrieves a static page by ID.
// --------------------------------------------------
func (r *MysqlRepository) GetByID(ctx context.Context, id uint64) (*staticpagemodel.StaticPage, error) {
	var page staticpagemodel.StaticPage

	err := r.db.Get(&page, "SELECT * FROM static_pages WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

// --------------------------------------------------
// GetBySlug retrieves a static page by slug.
// --------------------------------------------------
func (r *MysqlRepository) GetBySlug(ctx context.Context, slug string) (*staticpagemodel.StaticPage, error) {
	return nil, nil
}

// --------------------------------------------------
// List returns pages with filters, sorting, and pagination.
// --------------------------------------------------
func (r *MysqlRepository) List(ctx context.Context, req *staticpagepb.ListRequest) ([]*staticpagemodel.StaticPage, error) {
	query := `
		SELECT *
		FROM static_pages
		LIMIT ? OFFSET ?
	`
	var pages []*staticpagemodel.StaticPage

	err := r.db.SelectContext(ctx, &pages, query, 10, 0)
	if err != nil {
		return nil, err
	}

	return pages, nil
}

// --------------------------------------------------
// Count returns total items for the given filter (used for pagination).
// --------------------------------------------------
func (r *MysqlRepository) Count(ctx context.Context, req *staticpagepb.ListRequest) (uint64, error) {
	return 0, nil
}
