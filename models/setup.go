package models

import (
	"log"
	"os"
	"time"

	uuid "github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connection params
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	// Logger
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		return
	}

	DB = db
}

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (m *Base) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV6()
	if err != nil {
		return err
	}

	m.ID = id
	return
}
