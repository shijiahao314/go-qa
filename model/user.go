package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserID    uint64 `gorm:"primary_key;type:bigint unsigned;comment:用户ID" json:"userid"`
	Username  string `gorm:"type:varchar(32);comment:用户名" json:"username"`
	Password  string `gorm:"type:varchar(128);comment:密码" json:"password"`
	Role      string `gorm:"type:varchar(32);comment:角色" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type UserInfo struct {
	// **将uint64序列化json中格式标记为string，否则前端解析出现溢出导致精度问题
	UserID   uint64 `json:"userid,string"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
