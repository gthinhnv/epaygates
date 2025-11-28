package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"sort"
)

// AdsPlatform defines types of pages
type AdsPlatform int32

// AdsPlatformItem holds a AdsPlatform with a human-readable text
type AdsPlatformItem struct {
	ID   AdsPlatform
	Text string
}

// Scan implements the sql.Scanner interface
func (p *AdsPlatform) Scan(value interface{}) error {
	if value == nil {
		return errors.New("AdsPlatform: field is NULL, cannot scan")
	}

	switch v := value.(type) {
	case int64:
		*p = AdsPlatform(v)
		return nil
	case int32:
		*p = AdsPlatform(v)
		return nil
	case int:
		*p = AdsPlatform(v)
		return nil
	default:
		return fmt.Errorf("AdsPlatform: cannot scan value %v of type %T", value, value)
	}
}

// Value implements the driver.Valuer interface
func (p AdsPlatform) Value() (driver.Value, error) {
	return int64(p), nil
}

// String returns human-readable text for the AdsPlatform
func (p AdsPlatform) String() string {
	if text, ok := AdsPlatformMap[p]; ok {
		return text
	}
	return fmt.Sprintf("Unknown AdsPlatform (%d)", p)
}

// Enum constants
const (
	DISABLED AdsPlatform = iota + 1
	AUTO
	GOOGLE_ADSENSE
	PURPLEADS
	ADSKEEPER
	MGID
	PROPELLER
	ADCASH
	PUBFUTURE
	NETPUB
	ADSTERRA
)

// Map of AdsPlatform to human-readable text
var AdsPlatformMap = map[AdsPlatform]string{
	DISABLED:       "Disabled",
	AUTO:           "Auto",
	GOOGLE_ADSENSE: "GoogleAdsense",
	PURPLEADS:      "PurpleAds",
	ADSKEEPER:      "Adskeeper",
	MGID:           "Mgid",
	PROPELLER:      "Propeller",
	ADCASH:         "AdCash",
	PUBFUTURE:      "Pubfuture",
	NETPUB:         "Netpub",
	ADSTERRA:       "AdsTerra",
}

// AdsPlatforms is a slice of all AdsPlatforms, sorted by ID
var AdsPlatforms []AdsPlatformItem

func init() {
	AdsPlatforms = make([]AdsPlatformItem, 0, len(AdsPlatformMap))
	for id, text := range AdsPlatformMap {
		AdsPlatforms = append(AdsPlatforms, AdsPlatformItem{ID: id, Text: text})
	}

	// Sort by ID for consistent order
	sort.Slice(AdsPlatforms, func(i, j int) bool {
		return AdsPlatforms[i].ID < AdsPlatforms[j].ID
	})
}
