package storage

import (
	"github.com/timdin/wdipms/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StorageImpl struct {
	db *gorm.DB
}

func InitStorage(conn string) Storage {
	storageConfig := &gorm.Config{}
	db, err := gorm.Open(sqlite.Open(conn), storageConfig)
	if err != nil {
		panic(err)
	}
	db = migrateDB(db)

	return &StorageImpl{db: db}
}

func migrateDB(db *gorm.DB) *gorm.DB {
	migrateErr := db.AutoMigrate(model.Container{}, model.Item{})
	if migrateErr != nil {
		panic(migrateErr)
	}
	return db
}
