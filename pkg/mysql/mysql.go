package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlConn any

func InitMysqlDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:icephoenix@tcp(172.17.0.1:3306)/go_learn_web1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return
	}

	SqlConn, err = db.DB()

}