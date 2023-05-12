package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBConn *gorm.DB
var DefaultTx Transaction

func InitDatabase() *gorm.DB {
	// 连接数据库
	var err error
	DBConn, err = gorm.Open(sqlite.Open("my_postcrossing.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数名称作为表名
		},
	})
	if err != nil {
		panic("Failed to connect database.")
	}

	// 配置默认连接
	DefaultTx = Transaction{
		Tx: DBConn,
	}

	return DBConn
}
