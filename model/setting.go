package model

type ChatModel string

const (
	GPT_35_turbo      = "gpt-3.5-turbo"
	GPT_35_turbo_0301 = "gpt-3.5-turbo-0301"
	GPT_35_turbo_16k  = "gpt-3.5-turbo-16k"
	GPT_4             = "gpt-4"
)

type UserSetting struct {
	UserID         uint64    `gorm:"type:bigint unsigned primaryKey;comment:所属UserID" json:"userid,string"`
	OpenAI_API_Key string    `json:"openai_api_key"`
	ChatModel      ChatModel `json:"chat_model"`
	TestMode       bool      `json:"test_mode"`
}
