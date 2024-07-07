package mysql

import (
	"bluebell/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		config.Cfg.MySQLConfig.User,
		config.Cfg.MySQLConfig.Password,
		config.Cfg.MySQLConfig.Host,
		config.Cfg.MySQLConfig.Port,
		config.Cfg.MySQLConfig.DB)
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("open mysql err", zap.Error(err))
		return err
	}
	err = db.Ping()
	if err != nil {
		zap.L().Error("ping mysql err", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(config.Cfg.MySQLConfig.MaxOpen)
	db.SetMaxIdleConns(config.Cfg.MySQLConfig.MaxIdle)
	return nil
}

func CloseMysql() {
	err := db.Close()
	if err != nil {
		fmt.Printf("close mysql err: %v\n", err)
	}
	fmt.Println("close mysql success")
}
