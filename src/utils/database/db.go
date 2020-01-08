package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
		log "github.com/sirupsen/logrus"
		"github.com/owlsn/apis/src/common/config"
    "sync"
)

var db *gorm.DB

var err error

var once sync.Once

// Instance : Instance
func Instance() (*gorm.DB, error){
	once.Do(func(){
		// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// 	return "t_" + defaultTableName
		// }

		url := config.Conf.MySqlUrl
		dialect := "mysql"
		enableLog := config.Conf.ShowSql
		maxIdleConns := config.Conf.MaxIdleConns
		maxOpenConns := config.Conf.MaxOpenConns

		if db, err = gorm.Open(dialect, url); err != nil {
			log.Errorf("opens database failed: %s", err.Error())
			return
		} 
		db.LogMode(enableLog)
		db.SingularTable(true) // 禁用表名负数
		db.DB().SetMaxIdleConns(maxIdleConns)
		db.DB().SetMaxOpenConns(maxOpenConns)
	})
	return db, err
}