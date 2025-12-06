package staticpagerepo

import (
	"fmt"
	"shared/models/staticpagemodel"
	"strings"

	"github.com/jmoiron/sqlx"
)

var updateFieldMap = map[string]func(*staticpagemodel.StaticPage) (string, interface{}){
	"title": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "title = ?", p.Title
	},
	"slug": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "slug = ?", p.Slug
	},
}

func buildUpdateQuery(pageModel *staticpagemodel.StaticPage, fields []string) (string, []interface{}, error) {
	n := len(fields)
	if n == 0 {
		return "", nil, fmt.Errorf("update fields cannot be empty")
	}

	setClauses := make([]string, 0, n+1)
	args := make([]interface{}, 0, n+1)

	for _, f := range fields {
		if build, ok := updateFieldMap[f]; ok {
			clause, val := build(pageModel)
			setClauses = append(setClauses, clause)
			args = append(args, val)
		}
	}

	setClauses = append(setClauses, "updated_at = NOW()")

	var b strings.Builder
	b.Grow(64 + len(setClauses)*16)

	b.WriteString("UPDATE static_pages SET ")
	b.WriteString(strings.Join(setClauses, ", "))
	b.WriteString(" WHERE id = ?")

	args = append(args, pageModel.Id)

	return b.String(), args, nil
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
