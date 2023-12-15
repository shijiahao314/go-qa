<center><h1>GoMarkit</h1></center>

# gin

# zap


## 问题

1. 序列化后数据不一致
Q:
用户ID经过`c.JSON`序列化后next.js前端解析不一致；
`1`=>`1`
`7138163354002001920`=>`7138163354002002000`
`7138163354002001921`=>`7138163354002002000`
A:
在序列化tag中加上对该字段类型的定义string
```Golang
type UserInfo struct {
	// **将uint64序列化json中格式标记为string，否则前端解析出现溢出导致精度问题
	UserID   uint64 `json:"userid,string"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
```


