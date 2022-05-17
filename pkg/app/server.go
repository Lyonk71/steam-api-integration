// Package app provides the necessary functionality for building a server/service
package app

import (
	"SteamPurchaseService/pkg/api"
	"SteamPurchaseService/util"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	steamService api.SteamService
}

func NewServer(router *gin.Engine, steamService api.SteamService) *Server {
	return &Server{
		router:       router,
		steamService: steamService,
	}
}

func (s *Server) Run(config *util.Config) error {
	r := s.Routes()

	// run the server through the router
	err := r.Run(config.Port)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
