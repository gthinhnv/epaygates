package commonmodel

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type SEO struct {
	MetaTitle    string   `db:"meta_title" json:"meta_title"`
	MetaDesc     string   `db:"meta_desc" json:"meta_desc"`
	MetaKeywords []string `db:"meta_keywords" json:"meta_keywords"`
	Schema       string   `db:"schema" json:"schema"`
}

func (s *SEO) Scan(value interface{}) error {
	if value == nil {
		*s = SEO{} // safe default
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot scan %T into SEO", value)
	}

	return json.Unmarshal(b, s)
}

func (s SEO) Value() (driver.Value, error) {
	return json.Marshal(s)
}
