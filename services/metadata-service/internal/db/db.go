package db

import (
	"context"
	"encoding/json"
	"fmt"
	"metadatasvc/internal/db/migrations"
	"os"
	"path/filepath"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host                 string
	Port                 int
	User                 string
	Password             string
	Name                 string
	MaxOpenConnections   int
	MaxIdleConnections   int
	MaxLifeTimeInSeconds int
}

type DB struct {
	conn *sqlx.DB
}

func New(config *Config) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&interpolateParams=true&timeout=5s&readTimeout=30s&writeTimeout=30s&multiStatements=false",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("db connect error: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(config.MaxLifeTimeInSeconds) * time.Second)

	return &DB{conn: db}, nil
}

func (d *DB) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return d.conn.PingContext(ctx)
}

func (d *DB) Conn() *sqlx.DB {
	return d.conn
}

func (d *DB) Migrate() error {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	dir := filepath.Join(baseDir, "./migrations/versions")

	migs, err := migrations.LoadMigrations(dir)
	if err != nil {
		return err
	}

	if len(migs) == 0 {
		return nil
	}

	migrationLogPath := filepath.Join(baseDir, "./migrations/migrations.json")
	file, err := os.OpenFile(migrationLogPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open migration log file: %w", err)
	}
	defer file.Close()

	var migrationLogs map[int]*migrations.Migration
	if err := json.NewDecoder(file).Decode(&migrationLogs); err != nil && err.Error() != "EOF" {
		// If decoding fails (and it's not just an empty file), initialize an empty map
		return fmt.Errorf("failed to decode migration logs: %w", err)
	} else if migrationLogs == nil {
		migrationLogs = make(map[int]*migrations.Migration)
	}

	for _, mig := range migs {
		if mLog, exists := migrationLogs[mig.Version]; exists && mLog.Status == migrations.STATUS_SUCCESS {
			continue // Migration already applied
		}
		content, err := os.ReadFile(mig.Path)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", mig.Path, err)
		}

		tx, err := d.conn.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction for migration %s: %w", mig.Name, err)
		}

		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			mig.Status = migrations.STATUS_FAILED
			mig.Message = err.Error()
			mig.ExecutedAt = time.Now()
			migrationLogs[mig.Version] = &mig
			saveMigrationLogs(file, migrationLogs)
			return fmt.Errorf("failed to execute migration %s: %w", mig.Name, err)
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to commit transaction for migration %s: %w", mig.Name, err)
		}

		mig.Status = migrations.STATUS_SUCCESS
		mig.ExecutedAt = time.Now()
		migrationLogs[mig.Version] = &mig
		saveMigrationLogs(file, migrationLogs)
	}

	return nil
}

func saveMigrationLogs(file *os.File, logs map[int]*migrations.Migration) error {
	file.Truncate(0)
	file.Seek(0, 0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(logs); err != nil {
		return fmt.Errorf("failed to encode migration logs: %w", err)
	}
	return nil
}
