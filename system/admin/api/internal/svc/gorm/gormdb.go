package gorm

import (
	"erik-agile/common/xgorm"
	"erik-agile/system/admin/api/internal/config"

	"gorm.io/gorm"
)

type Gormdb struct {
    Gorm *gorm.DB
}

func NewGormdb(c config.Config) *Gormdb {
    return &Gormdb{
        Gorm: xgorm.NewGorm(c.Mysql.DataSource, c.Mysql.TablePrefix),
    }
}
