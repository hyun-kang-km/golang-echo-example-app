package handler

import (
	"github.com/hyun-kang-km/golang-echo-example-app/article"
	"github.com/hyun-kang-km/golang-echo-example-app/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
