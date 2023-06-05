package storage

import (
	"gitlab.com/burxondvv/blog-cicd-gitlab/config"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/db"
	"gitlab.com/burxondvv/blog-cicd-gitlab/pkg/logger"
	"gitlab.com/burxondvv/blog-cicd-gitlab/storage/postgres/aboutrepo"
)

type StorageI interface {
	About() aboutrepo.AboutI
}

type StoragePg struct {
	aboutRepo aboutrepo.AboutI
}

// NewStoragePg
func New(db *db.Postgres, log *logger.Logger, cfg *config.Config) StorageI {
	return &StoragePg{
		aboutRepo: aboutrepo.New(db, log, cfg),
	}
}

func (s *StoragePg) About() aboutrepo.AboutI {
	return s.aboutRepo
}
