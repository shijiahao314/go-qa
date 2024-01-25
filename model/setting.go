package model

type ChatModel string

const (
	GPT_35_turbo      = "gpt-3.5-turbo"
	GPT_35_turbo_0301 = "gpt-3.5-turbo-0301"
	GPT_35_turbo_16k  = "gpt-3.5-turbo-16k"
	GPT_4             = "gpt-4"
)

type UserSetting struct {
	UserID       uint64    `gorm:"primaryKey;type:bigint unsigned;comment:所属UserID"`
	OpenaiApiKey string    `gorm:"type:char(51);comment:OpenAI API Key"`
	ChatModel    ChatModel `gorm:"type:varchar(32);comment:对话模型"`
	TestMode     bool      `gorm:"comment:测试模式开启"`
}

type UserSettingDTO struct {
	OpenaiApiKey string    `json:"openai_api_key"`
	ChatModel    ChatModel `json:"chat_model"`
	TestMode     bool      `json:"test_mode"`
}
