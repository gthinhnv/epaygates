package common

type SEO struct {
	MetaTitle    string   `db:"meta_title"`
	MetaDesc     string   `db:"meta_desc"`
	MetaKeywords []string `db:"meta_keywords"`
	Schema       string   `db:"schema"`
}
