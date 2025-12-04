package translator

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Translator struct {
	locales  map[string]map[string]string // lang → key → value
	fallback string
}

type LocaleEntry struct {
	Name string // "en"
	Path string // "locales/en.json"
}

// ----------------------------------------------------------------------
// INIT
// ----------------------------------------------------------------------

func New(entries []*LocaleEntry, fallback string) *Translator {
	// Pre-allocate map based on number of locale files
	locales := make(map[string]map[string]string, len(entries))

	for _, e := range entries {
		if e == nil {
			continue
		}

		raw, err := os.ReadFile(e.Path)
		if err != nil {
			log.Printf("[i18n] missing file: %s", e.Path)
			continue
		}

		// Pre-allocate each language dictionary with estimated capacity
		msgs := make(map[string]string, 64)

		if err := json.Unmarshal(raw, &msgs); err != nil {
			log.Printf("[i18n] invalid json: %s (%v)", e.Path, err)
			continue
		}

		locales[e.Name] = msgs
	}

	return &Translator{
		locales:  locales,
		fallback: fallback,
	}
}

// ----------------------------------------------------------------------
// TRANSLATE
// ----------------------------------------------------------------------

func (t *Translator) T(lang, key string, params map[string]string) string {
	msgs, ok := t.locales[lang]
	if !ok {
		msgs = t.locales[t.fallback]
	}

	val := msgs[key]
	if val == "" {
		return key
	}

	// No params → return directly (zero allocations)
	if len(params) == 0 {
		return val
	}

	// Param replacement: fast linear replace
	for k, v := range params {
		val = strings.ReplaceAll(val, "{{"+k+"}}", v)
	}

	return val
}
