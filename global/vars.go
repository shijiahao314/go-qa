package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/shijiahao314/go-qa/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AppMode
type AppMode string

const (
	TEST ModeType = "TEST"
	DEV  ModeType = "DEV"
	PROD ModeType = "PROD"
)
const DefaultAppMode = DEV

// vars
var (
	Mode     AppMode          // TEST / DEV / PROD
	Config   *config.Config   // config.xxx.yaml
	DB       *gorm.DB         // DB: MySQL ...
	Logger   *zap.Logger      // Logger
	Redis    *redis.Client    // Redis
	Enforcer *casbin.Enforcer // casbin
	Etcd     *clientv3.Client // etcd
)

const (
	UserInfoKey     = "user_info"
	UserUserIDKey   = "userid"
	UserUsernameKey = "username"
	UserRoleKey     = "role"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
