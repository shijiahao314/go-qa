package errcode

const (
	Success = 0
	Failed  = 1

	// common
	ParamParseFailed = 10
	RoleMismatch     = 11
)

// auth
const (
	SignupFailed = 100 + iota
	LoginFailed
	UsernameOrPwd
	SessionDecode
	SessionSave
	NotLogin
)

// user
const (
	GetUsersFailed = 200 + iota
	UpdateUserFailed
	DeleteUserFailed
	AddUserFailed
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
)
