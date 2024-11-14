package routers

import (
	"fmt"
	"fuck-go/src/main/Websocket"
	"fuck-go/src/main/routers/records"
	"fuck-go/src/main/routers/users"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateRouter() {
	// 1.创建路由
	var r = gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	// 允许使用跨域请求  全局中间件
	r.Use(cors())
	// 设置路由组（统一前缀地址）
	apiGroup := r.Group("/hero-ranking")
	loadRoute(apiGroup)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// go watchFile()

	r.Run(":8888")
}

func loadRoute(apiGroup *gin.RouterGroup) {
	// ?Users
	apiGroup.POST("/user/insert", users.InsertUser)
	apiGroup.GET("/user/list", users.GetUsers)
	apiGroup.POST("/user/delete", users.DeleteUser)
	apiGroup.POST("/user/login", users.LoginUser)
	// ?Records
	apiGroup.GET("/record/list", records.GetRecords)
	apiGroup.POST("/record/insert", records.InsertRecord)
	apiGroup.POST("/record/delete", records.DeleteRecord)
	// ?websocket
	apiGroup.GET("/ws", func(c *gin.Context) {
		Websocket.Connect(c)
	})
}

func watchFile() {
	/*watch := FSNotify.NewNotifyFile()
	watch.WatchDir("/Users/lufangzhou/Documents/work-space/personal/hero-ranking/dist")
	select {}*/
	/*
		// 指定要监视的目录
		directoryToWatch := "/Users/lufangzhou/Documents/work-space/ep/数据支撑/质量平台/trimps-frontend-dim/dist"
		// 创建 fsnotify 监视器
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("Error creating watcher:", err)
			return
		}
		defer watcher.Close()

		// 将目录添加到监视器中
		err = filepath.Walk(directoryToWatch, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return watcher.Add(path)
		})
		if err != nil {
			fmt.Println("Error adding directory to watcher:", err)
			return
		}

		// 启动 goroutine 处理文件变化事件
		go func() {
			fmt.Println("aaaa")
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("File modified:", event.Name)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("Error:", err)
				}
			}
		}()*/
}

// ?跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
