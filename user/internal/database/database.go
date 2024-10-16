package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type DbService interface {
	Close() error
	Migrate(dst ...interface{}) error
	Conn() (*dbService, error)
}

type dbService struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *dbService
)

func New() (DbService, error) {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance, nil
	}
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		slog.Error(err)
	}

	dbInstance = &dbService{
		db: db,
	}
	return dbInstance, nil
}

func (s *dbService) Conn() (*dbService, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}
	_, err := New()
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

func (s *dbService) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return nil
}

func (s *dbService) Migrate(dst ...interface{}) error {
	err := s.db.Migrator().AutoMigrate(dst...)
	if err != nil {
		slog.Error(err)
		return err
	}
	return nil
}
