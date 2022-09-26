package xgorm

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGorm(dsn, prefix string) *gorm.DB {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold:             time.Second,   // 慢 SQL 阈值
            LogLevel:                  logger.Silent, // 日志级别
            IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:                  false,         // 禁用彩色打印
        },
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            TablePrefix:   prefix,
            SingularTable: true,
        },
        Logger: newLogger,
    })
    if err != nil {
        panic(err)
    }
    // db.AutoMigrate(
    //     &model.Admin{},
    //     &model.AdminLoginLog{},
    // )
    return db
}
