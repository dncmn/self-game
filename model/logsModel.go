package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// 日志表

type LogLogin struct {
	ID        string `gorm:"primary_key;"`
	UID       string `gorm:"column:uid;type:varchar(40)"`
	UserName  string `gorm:"column:user_name"`
	LoginTime int64  `gorm:"column:login_time"` // 登录时间戳
	LoginIP   string `gorm:"column:login_ip"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 重命名表名
func (this *LogLogin) TableName() string {
	return "log_logins"
}

// 回调中设置主键
func (this *LogLogin) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	scope.SetColumn("ID", ui.String())
}
