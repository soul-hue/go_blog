package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog/api/v1"
	"go_blog/middleware"
	"go_blog/utils"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	//后台

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUserList)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCate)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArticleInfo)
		auth.GET("admin/article", v1.GetArticleList)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}

	//前台
	router := r.Group("api/v1")
	{
		//用户模块
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUserList)
		router.DELETE("delete/:id",v1.DelUser)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		//	文章模块
		router.GET("addArticle",v1.AddArticle)
		router.GET("article/info/:id",v1.GetArticleInfo)
		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginFront", v1.LoginFront)
	}
	_ = r.Run(utils.HttpPort)
}
