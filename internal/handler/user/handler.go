package user

import (
	"github.com/Adharsh112/Tweets.git/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	api         *gin.Engine
	userService user.UserService
}

func NewHandler(api *gin.Engine, userService user.UserService) *Handler {
	return &Handler{
		api:         api,
		userService: userService,
	}
}

func (h *Handler) RouteList() {
}
