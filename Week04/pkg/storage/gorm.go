package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewGormDB() *gorm.DB {
	//读配置文件
	var user, password, address, port, dataBaseName string
	user = "root"
	password = "root"
	address = "localhost"
	port = "3306"
	dataBaseName = "gorm"
	dsn := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + dataBaseName + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // storage source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库链接错误")
	}
	return gormDB
}
