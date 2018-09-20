package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"self-game/constants"
	"time"
)

type User struct {
	ID           string                `gorm:"column:id;type:varchar(40);primary_key;" `
	UserName     string                `gorm:"column:user_name;type:varchar(40);"`
	Sex          constants.UserSexType `gorm:"column:sex"`
	Country      string                `gorm:"column:country"`
	City         string                `gorm:"column:city"`
	Password     string                `gorm:"column:password;"`
	Mobile       string                `gorm:"column:mobile"`
	RegisterIP   string                `gorm:"column:register_ip;"`
	RegosterTime int64                 `gorm:"column:register_time"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

// 重命名表名
func (this *User) TableName() string {
	return "users"
}

// 回调中设置主键
func (this *User) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	scope.SetColumn("ID", ui.String())
}
