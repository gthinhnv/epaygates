package staticpagerepo

import (
	"fmt"
	"metadatasvc/gen/go/staticpagepb"
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
	"content": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "content = ?", p.Content
	},
	"pageType": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "page_type = ?", p.PageType
	},
	"sortOrder": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "sort_order = ?", p.SortOrder
	},
	"status": func(p *staticpagemodel.StaticPage) (string, interface{}) {
		return "status = ?", p.Status
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

func buildFilters(req *staticpagepb.ListRequest) (string, []interface{}) {
	var (
		where []string
		args  []interface{}
	)

	// ids
	if len(req.Ids) > 0 {
		ph := make([]string, len(req.Ids))
		for i, id := range req.Ids {
			ph[i] = "?"
			args = append(args, id)
		}
		where = append(where, "id IN ("+strings.Join(ph, ",")+")")
	}

	// title LIKE
	if req.Title != "" {
		where = append(where, "title LIKE ?")
		args = append(args, "%"+req.Title+"%")
	}

	// slug
	if req.Slug != "" {
		where = append(where, "slug = ?")
		args = append(args, req.Slug)
	}

	// page_types
	if len(req.PageTypes) > 0 {
		ph := make([]string, len(req.PageTypes))
		for i, v := range req.PageTypes {
			ph[i] = "?"
			args = append(args, v)
		}
		where = append(where, "page_type IN ("+strings.Join(ph, ",")+")")
	}

	// ads_platforms
	if len(req.AdsPlatforms) > 0 {
		ph := make([]string, len(req.AdsPlatforms))
		for i, v := range req.AdsPlatforms {
			ph[i] = "?"
			args = append(args, v)
		}
		where = append(where, "ads_platform IN ("+strings.Join(ph, ",")+")")
	}

	// statuses
	if len(req.Statuses) > 0 {
		ph := make([]string, len(req.Statuses))
		for i, v := range req.Statuses {
			ph[i] = "?"
			args = append(args, v)
		}
		where = append(where, "status IN ("+strings.Join(ph, ",")+")")
	}

	// deleted_version
	if req.IsDeleted != nil {
		if req.GetIsDeleted() {
			where = append(where, "deleted_version > 0")
		} else {
			where = append(where, "deleted_version = 0")
		}
	}

	if len(where) == 0 {
		return "", args
	}

	return " WHERE " + strings.Join(where, " AND "), args
}

func buildListQuery(req *staticpagepb.ListRequest) (string, []interface{}, error) {
	var b strings.Builder
	b.Grow(256)

	// SELECT fields
	fields := strings.Join(req.Fields, ",")
	if fields == "" {
		fields = "*" // default
	}

	b.WriteString("SELECT ")
	b.WriteString(fields)
	b.WriteString(" FROM static_pages")

	// WHERE filters
	whereSQL, args := buildFilters(req)
	b.WriteString(whereSQL)

	// ORDER BY
	if req.OrderBy != "" {
		b.WriteString(" ORDER BY ")
		b.WriteString(req.OrderBy)
	}

	// LIMIT
	if req.Limit > 0 {
		b.WriteString(" LIMIT ?")
		args = append(args, req.Limit)
	}

	// OFFSET
	if req.Offset > 0 {
		b.WriteString(" OFFSET ?")
		args = append(args, req.Offset)
	}

	return b.String(), args, nil
}

func buildCountQuery(req *staticpagepb.ListRequest) (string, []interface{}) {
	var b strings.Builder
	b.Grow(128)

	b.WriteString("SELECT COUNT(*) FROM static_pages")

	whereSQL, args := buildFilters(req)
	b.WriteString(whereSQL)

	return b.String(), args
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
