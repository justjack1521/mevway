package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"mevway/internal/decorator"
)

type BanUser struct {
	UserID string
}

type BanUserHandler decorator.APIRouterHandler[BanUser]

type banUserHandler struct {
	client services.AccessServiceClient
}

func NewBanUserHandler(clt services.AccessServiceClient) BanUserHandler {
	return banUserHandler{
		client: clt,
	}
}

func (h banUserHandler) Handle(ctx *gin.Context, query BanUser) {
	fmt.Println("User banned mwahahaha!")
}
