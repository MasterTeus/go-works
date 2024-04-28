package config

import (
	"os"

	"github.com/MasterTeus/go-works.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if db file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("databse file not found, creating...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()

	}

	// DB connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(schemas.Opening{})
	if err != nil {
		logger.Errorf("migrate schema error: %v", err)
		return nil, err
	}

	return db, nil
}
