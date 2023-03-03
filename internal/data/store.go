package data

import (
	"nagato/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var store *Store

type Store struct {
	db *gorm.DB
}

func (s *Store) Download() *downloads {
	return newDownloads(s)
}

func (s *Store) Configs() *configs {
	return newConfigs(s)
}

func GetStore() *Store {
	if store == nil {
		SetSqlite()
	}

	return store
}

func SetSqlite() {
	// 打开 SQLite 数据库
	db, err := gorm.Open(sqlite.Open(model.GetConfig().Sqlite.DB), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	store = &Store{db}

	// 自动创建数据表
	db.AutoMigrate(&model.Config{})
	db.AutoMigrate(&model.Download{})
}
