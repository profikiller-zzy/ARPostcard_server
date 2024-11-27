package infra

import (
	"ARPostcard_server/biz/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlDB *gorm.DB

func InitMysql() {
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(conf.Conf.Mysql.Dsn()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
}
