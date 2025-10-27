package database

import (
	_ "database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	beginDelay     = 2 * time.Second
	thresholdDelay = 10 * time.Second
)

var defaultGormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},

	TranslateError:         true,
	SkipDefaultTransaction: true,
	Logger:                 logger.Default.LogMode(logger.Silent),
}

var DebugPostgresConfig = PostgresConfig{
	Host:    "localhost",
	Port:    "5432",
	User:    "postgres",
	DBName:  "postgres",
	SSLMode: "disable",
}

type PostgresConfig struct {
	Host, Port, User, Password, DBName, SSLMode string
}

func ConnectPostgres(conf PostgresConfig) (*gorm.DB, error) {
	slog.Info("database connection process",
		slog.String("host", conf.Host),
		slog.String("port", fmt.Sprintf("%s", conf.Port)),
		slog.String("user", conf.User),
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.DBName,
		conf.SSLMode,
	)

	var (
		db  *gorm.DB
		err error
	)

	for delay := beginDelay; delay <= thresholdDelay; delay <<= 1 {
		db, err = gorm.Open(gormPostgres.Open(dsn), defaultGormConfig)
		if err == nil {
			return db, nil
		}

		slog.Info("failed to connect to database, next try...")

		<-time.After(delay)
	}

	return nil, fmt.Errorf("failed to connect to postgres: %w", err)
}

func PostgresMigrations(db *gorm.DB) error {
	slog.Info("start of loading migrations")

	conn, _ := db.DB()

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to initialize migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to load migrations: %w", err)
		}

		slog.Info("migrations are not used by force, since the versions do not differ")

		return nil
	}

	slog.Info("migrations are successfully loaded")

	return nil
}
