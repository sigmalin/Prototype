## Prototype - API Server 
`專案不使用`
<br><br>

# 介紹

- Session 工具

<br><br>
    
# 使用方法

- 初始設定

```go
import (
    "context"
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
    go mgr.GC(context.Background(), 180)
}
```

- 存取方法

```go
import (
    "context"
    "github.com/gin-gonic/gin"
	"session"
)

// 產生新的 session
func NewSession(c *gin.Context) session.Session {
    sessionMgr := session.GetManager("session manager key")
	return sessionMgr.SessionStart(context.Background(), c.Writer, c.Request)
}

// 設定 session 值
func SetSeeeion(s session.Session, value interface{}) {
    s.Set(context.Background(), value)
}

// 取得 session 值
func GetSeeeion(c *gin.Context) (interface{}, error) {
    sessionMgr := session.GetManager("session manager key")
    s, err := sessionMgr.SessionRead(context.Background(), c.Writer, c.Request)
    if err != nil {
        // 沒有 session 資料
        return nil, err
    }
    return s.Get(context.Background()), nil 
}
```

<br><br>

# 備記

- 因為 prototype api server 使用 Json Web Token(JWT) 作為身分認證，因此停用 session 相關功能 