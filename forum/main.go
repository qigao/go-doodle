package main

//go:generate sqlboiler mysql

import (
	"forum/handler/article"
	"forum/handler/user"
	"forum/repository/mysql"
	"forum/utils"
	"github.com/labstack/echo/v4"
	"http/middleware"

	_ "forum/docs"
	webSwagger "github.com/swaggo/echo-swagger" // forum-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := echo.New()
	middleware.ConfigMiddleware(r)
	middleware.SetupZeroLog(r)
	setupRouter(r)
	r.Validator = utils.NewValidator()
	r.Logger.Fatal(r.Start("8585"))
}

func setupRouter(r *echo.Echo) {

	r.GET("/swagger/*", webSwagger.WrapHandler)

	v1 := r.Group("/api")

	d, _ := utils.NewMysqlORM()
	utils.AutoMigrate(d)

	us := mysql.NewUserRepository(d)
	as := mysql.NewArticleRepository(d)

	uh := user.NewUserHandler(us)
	ah := article.NewArticleHandler(as)

	userRouter := v1.Group("/users")
	uh.Register(userRouter)

	articleRouter := v1.Group("/articles")
	ah.Register(articleRouter)
}
