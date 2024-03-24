package errcode

const (
	Success = 0
	Failed  = 1

	// common
	InvalidRequest = 10 + iota
	RoleMismatch   = 11
	InternalServerError
	AssertError
)

// auth
const (
	SignupFailed = 100 + iota
	UsernameTooShort
	LoginFailed
	UsernameOrPwd
	SessionDecode
	SessionSaveFailed
	NotLogin
	NoRoleExist
	Unauthorized
)

// user
const (
	// add
	UserExists = 200 + iota
	AddUserFailed
	// delete
	DeleteUserFailed
	// update
	UpdateUserFailed
	// get
	GetUserFailed
	GetUsersFailed
)

// chat
const (
	// ChatInfo
	AddChatInfoFailed = 300 + iota
	DeleteChatInfoFailed
	UpdateChatInfoFailed
	GetChatInfosFailed
	// ChatCard
	AddChatCardFailed
	DeleteChatCardFailed
	UpdateChatCardFailed
	GetChatCardsFailed
	GetChatCardFailed
)

// setting
const (
	// UpdateSetting
	ChatModelNotExists = 400 + iota
	UpdateSettingFailed
	// GetSetting
	GetSettingFailed
)
