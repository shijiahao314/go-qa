package model

import (
	"database/sql"
	"time"
)

// UserRole
type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type User struct {
	UserID      uint64   `gorm:"primaryKey;type:bigint unsigned;comment:用户ID"`
	Username    string   `gorm:"type:varchar(32);comment:用户名"`
	Password    string   `gorm:"type:varchar(128);comment:密码"`
	Role        UserRole `gorm:"type:varchar(16);comment:角色"`
	Avatar      string   `gorm:"type:varchar(128);comment:头像"`
	Nickname    string   `gorm:"type:varchar(32);comment:昵称"`
	Email       string   `gorm:"type:varchar(32);comment:邮箱"`
	PhoneNumber string   `gorm:"type:varchar(11);comment:手机号"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
}

type UserDTO struct {
	UserID      uint64   `json:"userid,string"`
	Username    string   `json:"username"`
	Role        UserRole `json:"role,string"`
	Avatar      string   `json:"avatar"`
	Nickname    string   `json:"nickname"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
}
