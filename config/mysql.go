package config

import (
	"log"
	"os"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
    env := os.Getenv("ENV") // ENV 可以是 "prod" 或 "test"
	var dsn string
    switch env {
    case "prod":
        dsn = os.Getenv("MYSQL_DSN_PROD")
    case "test":
        dsn = os.Getenv("MYSQL_DSN_TEST")
    default:
        dsn = "root:123456@tcp(localhost:3306)/hulio_user?charset=utf8mb4&parseTime=True&loc=Local" // 默认值
    }

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 日志级别可调
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("get sql.DB failed: %v", err)
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 每个连接最大存活时间
}
