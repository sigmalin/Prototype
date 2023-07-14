`專案不使用`

# 介紹

- Session 工具
    
# 使用方法

- 初始設定

```go
import (
	"github.com/go-redis/redis"
    "session"
    provider "session/provider/redis"
)

func main() {
    // 連線 redis
    client := redis.NewClient(&redis.Options{
		Addr:     "", // redis ip 位址
		Password: "", // password
		DB:       0,
	})

    // 登錄 provider
    session.RegisterSessionProvider("session provider key", provider.NewProvider(client))

    // 產生 session manager
    mgr, _ := session.NewSessionManager("session provider key", "session name")
    // 登錄 session manager
	session.RegisterSessionManager("session manager key", mgr)

    // 處理過期的 session (秒)
    go mgr.GC(180)
}
```

- 存取方法

```go
import (
    "github.com/gin-gonic/gin"
	"session"
)

// 產生新的 session
func NewSession(c *gin.Context) session.Session {
    sessionMgr := session.GetManager("session manager key")
	return sessionMgr.SessionStart(c.Writer, c.Request)
}

// 設定 session 值
func SetSeeeion(s session.Session, value interface{}) {
    s.Set(value)
}

// 取得 session 值
func GetSeeeion(c *gin.Context) (interface{}, error) {
    sessionMgr := session.GetManager("session manager key")
    s, err := sessionMgr.SessionRead(c.Writer, c.Request)
    if err != nil {
        // 沒有 session 資料
        return nil, err
    }
    return s.Get(), nil 
}
```

# 備記

- 因為 prototype api server 使用 Json Web Token(JWT) 作為身分認證，因此停用 session 相關功能 