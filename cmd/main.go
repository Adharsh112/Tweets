package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adharsh112/Tweets.git/internal/config"
	userHandler "github.com/Adharsh112/Tweets.git/internal/handler/user"
	userRepo "github.com/Adharsh112/Tweets.git/internal/repository/user"
	userService "github.com/Adharsh112/Tweets.git/internal/service/user"
	"github.com/Adharsh112/Tweets.git/pkg/internalSQL"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "server is healthy",
		})
	})

	userRepo := userRepo.NewRepository(db)
	userService := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
