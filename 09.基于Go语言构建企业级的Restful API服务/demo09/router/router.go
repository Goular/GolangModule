package router

import (
	"github.com/gin-gonic/gin"
	"ApiServer/demo09/router/middleware"
	"net/http"
	"ApiServer/demo09/handler/sd"
	"ApiServer/demo09/handler/user"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// 错误代码测试
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

//读取 HTTP 信息： 在 API 开发中需要读取的参数通常为：HTTP Header、路径参数、URL参数、消息体，读取这些参数可以直接使用 gin 框架自带的函数：
//
//Param()：返回 URL 的参数值，例如
// router.GET("/user/:id", func(c *gin.Context) {
//     // a GET request to /user/john
//     id := c.Param("id") // id == "john"
// })
//
//Query()：读取 URL 中的地址参数，例如
//   // GET /path?id=1234&name=Manu&value=
//   c.Query("id") == "1234"
//   c.Query("name") == "Manu"
//   c.Query("value") == ""
//   c.Query("wtf") == ""
//
//DefaultQuery()：类似 Query()，但是如果 key 不存在，会返回默认值，例如
// //GET /?name=Manu&lastname=
// c.DefaultQuery("name", "unknown") == "Manu"
// c.DefaultQuery("id", "none") == "none"
// c.DefaultQuery("lastname", "none") == ""
//
// Bind()：检查 Content-Type 类型，将消息体作为指定的格式解析到 Go struct 变量中。apiserver 采用的媒体类型是 JSON，所以 Bind() 是按 JSON 格式解析的。
//
//GetHeader()：获取 HTTP 头。
//
//返回HTTP消息： 因为要返回指定的格式，apiserver 封装了自己的返回函数，通过统一的返回函数 SendResponse 来格式化返回，小节后续部分有详细介绍。


//在 RESTful API 开发中，API 经常会变动，为了兼容老的 API，引入了版本的概念，比如上例中的 /v1/user，说明该 API 版本是 v1。
//
//很多 RESTful API 最佳实践文章中均建议使用版本控制，笔者这里也建议对 API 使用版本控制。...
//
//https://juejin.im
//掘金 — 一个帮助开发者成长的社区