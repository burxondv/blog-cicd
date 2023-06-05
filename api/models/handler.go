package models

import (
	t "gitlab.com/burxondvv/blog-cicd-gitlab/api/tokens"
	"gitlab.com/burxondvv/blog-cicd-gitlab/config"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/logger"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage"
)

type HandlerV1Config struct {
	Logger     *logger.Logger
	Cfg        config.Config
	JWTHandler t.JWTHandler
	Postgres   storage.StorageI
}
