package global

import (
	"github.com/casbin/casbin/v2"
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
	Mode     ModeType         // TEST / DEV / PROD
	Config   *config.Config   // config.xxx.yaml
	DB       *gorm.DB         // DB: MySQL ...
	Logger   *zap.Logger      // Logger
	Redis    *redis.Client    // Redis
	Enforcer *casbin.Enforcer // casbin
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
