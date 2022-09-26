package xgorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewGorm(dsn, prefix string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            TablePrefix:   prefix,
            SingularTable: true,
        },
    })
    if err != nil {
        panic(err)
    }
    return db
}
