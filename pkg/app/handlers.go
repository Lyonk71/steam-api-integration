// Package app provides the handle functions for the routes
package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"SteamPurchaseService/pkg/api"
)

func (s *Server) APIStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "steam service running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) InitTxn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var initTxn api.InitTxnRequest
		err := c.ShouldBind(&initTxn)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		// kday todo: validate initTxn params
		response, err := s.steamService.InitTxn(initTxn)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) FinalizeTxn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var finalizeTxn api.FinalizeTxnRequest
		err := c.ShouldBind(&finalizeTxn)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		// kday todo: validate finalizeTxn params
		response, err := s.steamService.FinalizeTxn(finalizeTxn)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
