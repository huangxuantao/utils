package dal_util

import (
	"fmt"
	"gitea.com/huangxuantao89/xorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"veryon/utils/encrypt_util"
	"veryon/utils/log_util"

	// _ "github.com/mattn/go-sqlite3"
)

type DatabaseConfig struct {
	Driver      string
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	MaxOpenConn int
	ShowSQL     bool
}

// 数据库引擎
type Database struct {
	DB *xorm.Engine
}

func connStr(cfg *DatabaseConfig) (connStr string, err error) {
	password, err := encrypt_util.DesDecrypt(cfg.Password)
	if err != nil {
		return "", err
	}
	switch cfg.Driver {
	case "postgres":
		connStr = fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=disable",
			cfg.DBName, cfg.Host, cfg.User, password, cfg.Port)
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			cfg.User, password, cfg.Host, cfg.Port, cfg.DBName)
	case "sqlite3":
		connStr = fmt.Sprintf("%s.db", cfg.DBName)
	}
	return connStr, nil
}

func connDB(cfg *DatabaseConfig) (*Database, error) {
	var DBInstance = Database{}
	connStr, err := connStr(cfg)
	if err != nil {
		return nil, err
	}
	db, err := xorm.NewEngine(cfg.Driver, connStr)
	if err != nil || db == nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil || db == nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.ShowSQL(cfg.ShowSQL)
	log_util.Logger.Debugf("DB/Connect: Connected to %s", cfg.Host)
	DBInstance.DB = db
	return &DBInstance, nil
}

func GetDB(cfg *DatabaseConfig) (*Database, error) {
	d, err := connDB(cfg)
	if err != nil {
		log_util.Logger.Errorf("DB/Connect:%s", err.Error())
		return nil, err
	}
	return d, nil
}
