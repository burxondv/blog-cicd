package aboutsh

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/burxondvv/blog-cicd-gitlab/api/models"
	t "gitlab.com/burxondvv/blog-cicd-gitlab/api/tokens"
	"gitlab.com/burxondvv/blog-cicd-gitlab/config"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/logger"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage"
)

type AboutI interface {
	Create(c *gin.Context)
	FindOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type AboutHandler struct {
	log        *logger.Logger
	cfg        config.Config
	jwthandler t.JWTHandler
	postgres   storage.StorageI
}

func New(c *models.HandlerV1Config) AboutI {
	return &AboutHandler{
		log:        c.Logger,
		cfg:        c.Cfg,
		jwthandler: c.JWTHandler,
		postgres:   c.Postgres,
	}
}
