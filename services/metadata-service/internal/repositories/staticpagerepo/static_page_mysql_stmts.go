package staticpagerepo

import (
	"github.com/jmoiron/sqlx"
)

var createStmt *sqlx.Stmt
var updateStmt *sqlx.Stmt
var deleteStmt *sqlx.Stmt
var getByIDStmt *sqlx.Stmt

func prepareStatements(db *sqlx.DB) error {
	var err error

	createStmt, err = prepareCreateStatement(db)
	if err != nil {
		return err
	}

	updateStmt, err = prepareUpdateStatement(db)
	if err != nil {
		return err
	}

	deleteStmt, err = prepareDeleteStatement(db)
	if err != nil {
		return err
	}

	getByIDStmt, err = prepareGetByIDStatement(db)
	if err != nil {
		return err
	}

	return nil
}

// prepareCreateStatement prepares the SQL statement for inserting a static page
func prepareCreateStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
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
	return db.Preparex(query)
}

func prepareUpdateStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	const query = `
        SELECT
			id,
            title,
            slug,
            content,
            page_type,
            sort_order,
            ads_platform,
            status,
            created_by,
            updated_by,
            created_at,
            updated_at,
            deleted_version
		FROM static_pages
        WHERE id = ? AND deleted_version = 0
        LIMIT 1
    `
	return db.Preparex(query)
}

func prepareDeleteStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	const query = `
        SELECT * FROM static_pages
        WHERE id = ? AND deleted_version = 0
        LIMIT 1
    `
	return db.Preparex(query)
}

func prepareGetByIDStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	const query = `
        SELECT * FROM static_pages
        WHERE id = ? AND deleted_version = 0
        LIMIT 1
    `
	return db.Preparex(query)
}
