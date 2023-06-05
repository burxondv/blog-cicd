package v1

import (
	about "gitlab.com/burxondvv/blog-cicd-gitlab/api/handlers/v1/about_h"
	pingpongh "gitlab.com/burxondvv/blog-cicd-gitlab/api/handlers/v1/ping_pong_h"

	"gitlab.com/burxondvv/blog-cicd-gitlab/api/models"
)

type HandlerV1I interface {
	About() about.AboutI
	Ping() pingpongh.PingI
}

type handlerV1 struct {
	ping  pingpongh.PingI
	about about.AboutI
}

// New ...
func New(c *models.HandlerV1Config) HandlerV1I {
	return &handlerV1{
		ping:  pingpongh.New(c),
		about: about.New(c),
	}
}

func (h *handlerV1) Ping() pingpongh.PingI {
	return h.ping
}

func (h *handlerV1) About() about.AboutI {
	return h.about
}
