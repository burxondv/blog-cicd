package pingpongh

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/burxondvv/blog-cicd-gitlab/api/models"
	t "gitlab.com/burxondvv/blog-cicd-gitlab/api/tokens"
	"gitlab.com/burxondvv/blog-cicd-gitlab/config"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/logger"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage"
)

type PingI interface {
	Ping(c *gin.Context)
}

type PingHandler struct {
	log        *logger.Logger
	cfg        config.Config
	jwthandler t.JWTHandler
	postgres   storage.StorageI
}

func New(c *models.HandlerV1Config) PingI {
	return &PingHandler{
		log:        c.Logger,
		cfg:        c.Cfg,
		jwthandler: c.JWTHandler,
		postgres:   c.Postgres,
	}
}
