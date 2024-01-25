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
	GetUsersFailed
	UpdateUserFailed
	PermissionDenied
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
	UpdateSettingFailed = 400 + iota
	// GetSetting
	GetSettingFailed
)
