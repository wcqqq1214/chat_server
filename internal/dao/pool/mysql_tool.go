package pool

import (
	"chat-room/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	username := config.GetConfig().MySQL.User     // 账号
	password := config.GetConfig().MySQL.Password // 密码
	host := config.GetConfig().MySQL.Host         // 数据库地址，可以是 IP 或者域名
	port := config.GetConfig().MySQL.Port         // 数据库端口
	Dbname := config.GetConfig().MySQL.Name       // 数据库名
	timeout := "10s"                              // 连接超时，10 秒

	// 拼接下 dsn 参数，dsn 格式可以参考上面的语法，这里使用 Sprintf 动态拼接 dsn 参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接 dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	var err error

	// 连接 MySQL，获得 DB 类型实例，用于后面的数据库读写操作
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	// 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  // 连接池最大允许的空闲连接数，如果没有 sql 任务，需要执行的连接数大于 20，超过的连接会被连接池关闭
}

func GetDB() *gorm.DB {
	return _db
}
