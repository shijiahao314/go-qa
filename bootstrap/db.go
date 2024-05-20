package bootstrap

import (
	"github.com/shijiahao314/go-qa/config"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	global.Logger.Info("start to init database", zap.String("type", string(global.Config.Database.Type)))
	var db *gorm.DB
	switch global.Config.Database.Type {
	case config.DatabaseTypeSqlite:
		db = initSqlite()
	case config.DatabaseTypeMysql:
		db = initMysql()
	default:
		global.Logger.Error("unsupported database type", zap.String("type", string(global.Config.Database.Type)))
		panic("unsupported database type")
	}

	// 自动建表
	err := db.AutoMigrate(
		&model.User{},
		&model.ChatInfo{},
		&model.ChatCard{},
		&model.UserSetting{},
	)
	if err != nil {
		global.Logger.Error("failed during automigration", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("successfully init database")
	return db
}

func initSqlite() *gorm.DB {
	global.Logger.Info("start to init sqlite")
	cfg := global.Config.Database.Sqlite
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		global.Logger.Error("failed to init sqlite", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("successfully connect to db", zap.String("path", cfg.Path))
	return db
}

func initMysql() *gorm.DB {
	global.Logger.Info("start to init mysql")
	m := global.Config.Database.Mysql
	dsn := m.Dsn()
	cfg := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(cfg))
	if err != nil {
		global.Logger.Error("failed to init mysql", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("successfully connect to db", zap.String("dsn", dsn))
	return db
}
