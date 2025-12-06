package staticpagemodel

import (
	"shared/models/commonmodel"
	"time"
)

type StaticPage struct {
	Id             uint64                  `db:"id" json:"id"`
	Title          string                  `db:"title" json:"title" validate:"required,min=1,max=255"`
	Slug           string                  `db:"slug" json:"slug" validate:"required,min=1,max=255"`
	Content        string                  `db:"content" json:"content"`
	PageType       commonmodel.PageType    `db:"page_type" json:"pageType"`
	SortOrder      int32                   `db:"sort_order" json:"sortOrder"`
	Seo            *commonmodel.SEO        `db:"seo" json:"seo"`
	AdsPlatform    commonmodel.AdsPlatform `db:"ads_platform" json:"adsPlatform"`
	Status         commonmodel.Status      `db:"status" json:"status"`
	CreatedBy      uint64                  `db:"created_by" json:"createdBy"`
	UpdatedBy      uint64                  `db:"updated_by" json:"updatedBy"`
	CreatedAt      time.Time               `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time               `db:"updated_at" json:"updatedAt"`
	DeletedVersion int32                   `db:"deleted_version" json:"-"`
}

func (m *StaticPage) TableName() string {
	return "static_pages"
}

type StaticPageUpdate struct {
	StaticPage
	Fields []string `json:"fields"`
}
