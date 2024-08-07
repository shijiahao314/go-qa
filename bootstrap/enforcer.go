package bootstrap

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/go-admin-team/gorm-adapter/v3"
	"github.com/shijiahao314/go-qa/global"
)

func initEnforcer() *casbin.Enforcer {
	var e *casbin.Enforcer

	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	// a, _ := gormadapter.NewAdapter(global.Config.Database.Type, "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	a, _ := gormadapter.NewAdapterByDB(global.DB)
	e, _ = casbin.NewEnforcer("config/casbin/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Modify the policy.
	e.AddPolicy("admin", "/api/*")
	e.AddPolicy("user", "/api/chat/*")
	e.AddPolicy("user", "/api/settings/*")
	// e.AddGroupingPolicy("alice", "admin")
	e.SavePolicy()

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()

	return e
}
