package bootstrap

import (
	"fmt"
	"log/slog"

	"github.com/shijiahao314/go-qa/config"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initDB 初始化数据库
func mustInitDB() *gorm.DB {
	databaseType := global.Config.Database.Type
	slog.Info("start to init database", slog.String("type", string(databaseType)))
	defer slog.Info("success init database")

	var db *gorm.DB
	switch databaseType {
	case config.DatabaseTypeSqlite:
		db = initSqlite()
	case config.DatabaseTypeMysql:
		db = initMysql()
	default:
		slog.Error("unsupported database", slog.String("type", string(databaseType)))
		panic(fmt.Sprintf("unsupported database type: %s", databaseType))
	}

	// 自动建表
	err := db.AutoMigrate(
		&model.User{},
		&model.ChatInfo{},
		&model.ChatCard{},
		&model.UserSetting{},
	)
	if err != nil {
		slog.Error("failed during automigration", slog.String("err", err.Error()))
		panic(err)
	}

	return db
}

func initSqlite() *gorm.DB {
	slog.Info("start to init sqlite")
	defer slog.Info("success init sqlite")

	cfg := global.Config.Database.Sqlite
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		slog.Error("failed to init sqlite", slog.String("err", err.Error()))
		panic(err)
	}

	return db
}

func initMysql() *gorm.DB {
	slog.Info("start to init mysql")
	defer slog.Info("success init mysql")

	m := global.Config.Database.Mysql
	dsn := m.Dsn()
	cfg := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(cfg))
	if err != nil {
		slog.Error("failed to init mysql", slog.String("err", err.Error()))
		panic(err)
	}

	return db
}
