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
	return "data.log_logins"
}

// 回调中设置主键
func (this *LogLogin) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", ui.String())
}

// 记录用户向公众号发送消息的记录
type LogUserSendMsgToWechat struct {
	ID              string `gorm:"primary_key;"`
	OpenID          string `json:"open_id"`                       // 消息的发送者
	MsgType         string `gorm:"column:msg_type"`               // 发送消息的类型
	Content         string `gorm:"column:content;varchar:102400"` // 发送消息的内容
	CreateTimeStamp string `gorm:"column:create_time"`            // 消息的添加时间
	MsgID           string `gorm:"column:msg_id"`                 // 消息对应的id
	MediaID         string `gorm:"column:media_id"`               // 资源id
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

// 重命名表名
func (this *LogUserSendMsgToWechat) TableName() string {
	return "data.logUserSendMsgToWechat"
}

// 回调中设置主键
func (this *LogUserSendMsgToWechat) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", ui.String())
}
