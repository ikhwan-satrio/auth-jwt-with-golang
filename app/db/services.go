package db

import (
	"context"
	"fmt"

	"github.com/ikhwan-satrio/auth-golang/app/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBService struct {
	Config config.DBConfig
}

func (s *DBService) CreateDB() (*gorm.DB, context.Context, error) {
	var dialector gorm.Dialector

	switch s.Config.Driver {
	case "sqlite":
		dialector = sqlite.Open(s.Config.DSN)
	case "postgres":
		// dialector = postgres.Open(s.Config.DSN)
		return nil, nil, fmt.Errorf("postgres driver not implemented yet")
	case "mysql":
		// dialector = mysql.Open(s.Config.DSN)
		return nil, nil, fmt.Errorf("mysql driver not implemented yet")
	default:
		return nil, nil, fmt.Errorf("unsupported driver: %s", s.Config.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})

	ctx := context.Background()

	if err != nil {
		return nil, nil, err
	}

	db.AutoMigrate(&User{})

	return db, ctx, nil
}
