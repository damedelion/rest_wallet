package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/damedelion/rest_wallet/config"
	_ "github.com/lib/pq"
)

func NewDBOpen(config *config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode, config.Timezone)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db, err: %v", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime) * time.Second)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to verify connection to db, err: %v", err)
	}

	return db, nil
}
