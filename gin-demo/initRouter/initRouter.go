package initRouter

// 可以理解是存放路由
import (
	"gin-demo/handler" // 引入handler包
	"gin-demo/utils"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
)

// 可以优化这块，利用分组概念来区分项目或模块的路由
// func SetupRouter() *gin.Engine {
// 	// 初始路由
// 	router := gin.Default()

// 	// 添加 Get 请求路由
// 	router.GET("/", retHelloGinAndMethod)
// 	// 添加 Post 请求路由
// 	router.POST("/", retHelloGinAndMethod)
// 	// 添加 Put 请求路由
// 	router.PUT("/", retHelloGinAndMethod)
// 	// 添加 Delete 请求路由
// 	router.DELETE("/", retHelloGinAndMethod)
// 	// 添加 Patch 请求路由
// 	router.PATCH("/", retHelloGinAndMethod)
// 	// 添加 Head 请求路由
// 	router.HEAD("/", retHelloGinAndMethod)
// 	// 添加 Options 请求路由
// 	router.OPTIONS("/", retHelloGinAndMethod)

// 	// 添加 user
// 	router.GET("/user/:name", handler.UserSave)

// 	return router
// }

func SetupRouter() *gin.Engine {
	// 初始路由
	router := gin.Default()

	// 加载模板路径
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*") //测试模式下
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	// 加载静态资源
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics/")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))

	// 如下还可以再优化
	// /
	// index := router.Group("/")
	// {
	// 	// 添加 Get 请求路由
	// 	index.GET("", retHelloGinAndMethod)
	// 	// 添加 Post 请求路由
	// 	index.POST("", retHelloGinAndMethod)
	// 	// 添加 Put 请求路由
	// 	index.PUT("", retHelloGinAndMethod)
	// 	// 添加 Delete 请求路由
	// 	index.DELETE("", retHelloGinAndMethod)
	// 	// 添加 Patch 请求路由
	// 	index.PATCH("", retHelloGinAndMethod)
	// 	// 添加 Head 请求路由
	// 	index.HEAD("", retHelloGinAndMethod)
	// 	// 添加 Options 请求路由
	// 	index.OPTIONS("", retHelloGinAndMethod)
	// }

	index := router.Group("/")
	{
		// 添加 Get 请求路由
		index.Any("", handler.Index)
	}

	// /user
	userRouter := router.Group("/user")
	{
		// userRouter.GET("/:name", handler.UserSave)
		// userRouter.GET("", handler.UserSaveByQuery)
		// 注册用户
		userRouter.POST("/register", handler.UserRegister)
		// 查询用户
		userRouter.POST("/login", handler.UserLogin)
		// 查询用户信息
		userRouter.GET("/profile/", handler.UserProfile)
		// 更新用户信息
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	return router
}
