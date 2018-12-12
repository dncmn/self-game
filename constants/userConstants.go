package constants

// 用户性别
type UserSexType int8

const (
	UserSexTypeMale UserSexType = iota + 1
	_
	_
	UserSexTypeTotal
)

// 解锁类型
type UnlockCourseType int8

const (
	UnlockCourseByFee UnlockCourseType = iota + 1 // 付费解锁
	UnlockCourseByGM                              // 人工解锁(gm后台解锁)
)
