package routers

import (
	"fmt"
	"fuck-go/src/main/routers/records"
	"fuck-go/src/main/routers/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
	r.Run(":8888")
}

func loadRoute(apiGroup *gin.RouterGroup) {
	// ?Users
	apiGroup.POST("/user/insert", users.InsertUser)
	apiGroup.GET("/user/list", users.GetUsers)
	apiGroup.POST("/user/delete", users.DeleteUser)
	// ?Records
	apiGroup.GET("/record/list", records.GetRecords)
	apiGroup.POST("/record/insert", records.InsertRecord)
	apiGroup.POST("/record/delete", records.DeleteRecord)
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
			//              允许跨域设置                                                                                                      可以返回其他子段
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