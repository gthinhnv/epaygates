package migrations

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MigrationStatus string

const (
	STATUS_SUCCESS MigrationStatus = "success"
	STATUS_FAILED  MigrationStatus = "failed"
)

type Migration struct {
	Version    int
	Name       string
	Path       string
	Status     MigrationStatus
	Message    string
	ExecutedAt time.Time
}

// LoadMigrations loads .sql migration files sorted by version (001_x.sql → version 1).
func LoadMigrations(dir string) ([]Migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	migs := []Migration{}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()

		// Only accept .sql files
		if !strings.HasSuffix(name, ".sql") {
			continue
		}

		version, err := parseVersion(name)
		if err != nil {
			return nil, err
		}

		migs = append(migs, Migration{
			Version: version,
			Name:    name,
			Path:    filepath.Join(dir, name),
		})
	}

	// Sort by version
	sort.Slice(migs, func(i, j int) bool {
		return migs[i].Version < migs[j].Version
	})

	return migs, nil
}

// parseVersion extracts leading number prefix: e.g. "001_create_users.sql" => 1
func parseVersion(filename string) (int, error) {
	i := 0
	for i < len(filename) && filename[i] >= '0' && filename[i] <= '9' {
		i++
	}
	if i == 0 {
		return 0, nil
	}
	return strconv.Atoi(filename[:i])
}
