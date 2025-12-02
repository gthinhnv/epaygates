package common

import (
	"encoding/json"
	"fmt"
)

type SEO struct {
	MetaTitle    string   `db:"meta_title"`
	MetaDesc     string   `db:"meta_desc"`
	MetaKeywords []string `db:"meta_keywords"`
	Schema       string   `db:"schema"`
}

func (s *SEO) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot scan type %T into SEO", value)
	}
	return json.Unmarshal(bytes, s)
}
