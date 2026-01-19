package user

import (
	"github.com/Adharsh112/Tweets.git/internal/config"
	"github.com/Adharsh112/Tweets.git/internal/repository/user"
)

type UserService interface {
	//	Register(cts context.Context, req *dto.RegisterRequest,)(int64, int, error){
	//
	// }
}
type userService struct {
	cfg      *config.Config
	userRepo user.UserRepository
}

func NewService(cfg *config.Config, userRepo user.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
