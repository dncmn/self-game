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

type ConfigLevelTest struct {
	ID         string `gorm:"primary_key;"`
	Level      int    `gorm:"column:level"`
	Typ        int    `gorm:"column:type"`
	Index      string `gorm:"column:index"`
	Answer     string `gorm:"column:answer;type:varchar(1024)"`
	Text       string `gorm:"column:text;type:varchar(1024)"`
	ChoiceList string `gorm:"column:choice_list;type:varchar(1024)"`
	ImageList  string `gorm:"column:image_list;type:varchar(1024)"`
	VoiceURL   string `gorm:"column:voice_url:varchar(1024)"`
}

// 重命名表名
func (this *ConfigLevelTest) TableName() string {
	return "config_level_test"
}

// 回调中设置主键
func (this *ConfigLevelTest) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	scope.SetColumn("ID", ui.String())
}
