package staticpage

import (
	"shared/models/common"
	"time"
)

type StaticPage struct {
	Id             uint64             `db:"id"`
	Title          string             `db:"title"`
	Slug           string             `db:"slug"`
	Content        string             `db:"content"`
	PageType       common.PageType    `db:"page_type"`
	SortOrder      int32              `db:"sort_order"`
	Seo            *common.SEO        `db:"seo"`
	AdsPlatform    common.AdsPlatform `db:"ads_platform"`
	Status         common.Status      `db:"status"`
	CreatedBy      uint64             `db:"created_by"`
	UpdatedBy      uint64             `db:"updated_by"`
	CreatedAt      time.Time          `db:"created_at"`
	UpdatedAt      time.Time          `db:"updated_at"`
	DeletedVersion int32              `db:"deleted_version"`
}

func (m *StaticPage) TableName() string {
	return "static_pages"
}
