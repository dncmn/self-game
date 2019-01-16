package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"self-game/constants"
	"time"
)

type User struct {
	ID        string                `gorm:"column:id;type:varchar(40);primary_key;" `
	UID       string                `gorm:"column:uid"`
	UserName  string                `gorm:"column:user_name;type:varchar(40);"`
	Sex       constants.UserSexType `gorm:"column:sex"`
	Country   string                `gorm:"column:country"`
	City      string                `gorm:"column:city"`
	Mobile    string                `gorm:"column:mobile"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type UserRegister struct {
	ID           string `gorm:"column:id;type:varchar(40);primary_key;" `
	UserName     string `gorm:"column:user_name;comment:'用户名'"`          // 用户名
	Password     string `gorm:"column:password;"`                        //  用户注册的密码
	RegisterIP   string `gorm:"column:register_ip;"`                     // 用户注册时的ip
	RegisterTime int64  `gorm:"column:register_time"`                    // 注册的时间戳
	User         User   `gorm:"ForeignKey:ID;AssociationForeignKey:UID"` // ForeignKey:指的是User表中的关联的key,AssociationForeignKey:UserRegister的关联的id
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

type UserCourse struct {
	ID            string                     `gorm:"column:id;type:varchar(40);primary_key;" `
	UID           string                     `gorm:"column:uid;not null"`
	CourseID      int                        `gorm:"column:course_id;not null;"`    // 课程id
	IsPay         bool                       `gorm:"column:is_pay;"`                // 课程是否付费
	UnlockType    constants.UnlockCourseType `gorm:"column:unlock_type;default:-1"` // 默认没有解锁
	Process       int                        `gorm:"column:process"`                // 课程进度
	UnLockProcess int                        `gorm:"column:unlock_process"`         // 未付费时的课程进度
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

// 重命名表名
func (this *UserRegister) TableName() string {
	return "user_register"
}

// 回调中设置主键
func (this *UserRegister) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", ui.String())
}

// 重命名表名
func (this *User) TableName() string {
	return "users"
}

// 回调中设置主键
func (this *User) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", ui.String())
}

// 重命名表名
func (this *UserCourse) TableName() string {
	return "user_course"
}

// 回调中设置主键
func (this *UserCourse) BeforeCreate(scope *gorm.Scope) {
	ui, _ := uuid.NewV4()
	_ = scope.SetColumn("ID", ui.String())
}
