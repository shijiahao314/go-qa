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
	UserID   uint64 `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
