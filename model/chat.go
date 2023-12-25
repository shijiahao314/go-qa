package model

import (
	"time"
)

type ChatRole string

const (
	ChatRoleSystem    = "system"
	ChatRoleUser      = "user"
	ChatRoleAssistant = "assistant"
)

type ChatInfo struct {
	ID        uint      `gorm:"primaryKey;comment:对话组ID" json:"id,string"`
	UserID    uint64    `gorm:"type:bigint unsigned;comment:所属UserID" json:"userid,string"`
	Title     string    `gorm:"type:varchar(32);comment:题目" json:"title"`
	Num       uint      `gorm:"comment:对话数量" json:"num"`
	CreatedAt time.Time `gorm:"comment:创建时间" json:"ctime"`
	UpdatedAt time.Time `gorm:"comment:修改时间" json:"utime"`
}

type ChatCard struct {
	ID         uint     `gorm:"primarykey;comment:对话卡片ID"`
	ChatInfoID uint     `gorm:"comment:所属ChatInfoID" json:"chat_info_id"` // 外键
	Content    string   `gorm:"type:text;comment:内容" json:"content"`
	Role       ChatRole `gorm:"type:varchar(8);comment:角色" json:"role"`
}
