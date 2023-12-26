package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/shijiahao314/go-qa/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ModeType
type ModeType string

const (
	TEST = "TEST"
	DEV  = "DEV"
	PROD = "PROD"
)
const DEFAULT_MODE = DEV

// vars
var (
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
	Redis  *redis.Client
	Mode   ModeType
)

const (
	USER_INFO_KEY     = "user_info"
	USER_USER_ID_KEY  = "userid"
	USER_USERNAME_KEY = "username"
	USER_ROLE_KEY     = "role"
)

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)
