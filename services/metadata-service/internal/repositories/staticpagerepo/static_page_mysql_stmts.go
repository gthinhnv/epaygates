package staticpagerepo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var createStmt *sql.Stmt

func prepareStatements(db *sqlx.DB) error {
	var err error
	createStmt, err = prepareCreateStatement(db)
	if err != nil {
		return err
	}
	return nil
}

// prepareCreateStatement prepares the SQL statement for inserting a static page
func prepareCreateStatement(db *sqlx.DB) (*sql.Stmt, error) {
	const query = `
        INSERT INTO static_pages (
            title,
            slug,
            content,
            page_type,
            sort_order,
            seo,
            ads_platform,
            status,
            created_by,
            updated_by,
            created_at,
            updated_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
    `
	return db.Prepare(query)
}
