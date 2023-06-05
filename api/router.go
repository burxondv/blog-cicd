package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/burxondvv/blog-cicd-gitlab/api/docs" //swag
	v1 "gitlab.com/burxondvv/blog-cicd-gitlab/api/handlers/v1"
	"gitlab.com/burxondvv/blog-cicd-gitlab/api/middleware"
	"gitlab.com/burxondvv/blog-cicd-gitlab/api/models"
	token "gitlab.com/burxondvv/blog-cicd-gitlab/api/tokens"
	"gitlab.com/burxondvv/blog-cicd-gitlab/config"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/logger"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         *logger.Logger
	CasbinEnforcer *casbin.Enforcer
	Postgres       storage.StorageI
}

// New ...
// @title           Blog site project API Endpoints
// @version         1.0
// @description     Here QA can test and frontend or mobile developers can get information of API endpoints.

// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SignInKey,
		Log:       option.Logger,
	}

	h := v1.New(&models.HandlerV1Config{
		Logger:     option.Logger,
		Cfg:        option.Conf,
		JWTHandler: jwtHandler,
		Postgres:   option.Postgres,
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, option.Conf))
	router.GET("/", h.Ping().Ping)
	api := router.Group("/v1")
	api.GET("/", h.Ping().Ping)

	about := api.Group("/about")
	about.POST("", h.About().Create)
	about.GET("/:id", h.About().FindOne)
	about.PUT("", h.About().Update)
	about.DELETE(":id", h.About().Delete)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
