package staticpage

import (
	"shared/models/common"
	"time"
)

type StaticPage struct {
	Id             uint64             `db:"id" json:"id"`
	Title          string             `db:"title" json:"title"`
	Slug           string             `db:"slug" json:"slug"`
	Content        string             `db:"content" json:"content"`
	PageType       common.PageType    `db:"page_type" json:"pageType"`
	SortOrder      int32              `db:"sort_order" json:"sortOrder"`
	Seo            *common.SEO        `db:"seo" json:"seo"`
	AdsPlatform    common.AdsPlatform `db:"ads_platform" json:"adsPlatform"`
	Status         common.Status      `db:"status" json:"status"`
	CreatedBy      uint64             `db:"created_by" json:"createdBy"`
	UpdatedBy      uint64             `db:"updated_by" json:"updatedBy"`
	CreatedAt      time.Time          `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time          `db:"updated_at" json:"updatedAt"`
	DeletedVersion int32              `db:"deleted_version" json:"-"`
}

func (m *StaticPage) TableName() string {
	return "static_pages"
}
