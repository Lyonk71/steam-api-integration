package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.APIStatus())
		// prefix the user routes
		microtxn := v1.Group("/microtxn")
		{
			microtxn.POST("initTxn", s.InitTxn())
			microtxn.POST("finalizeTxn", s.FinalizeTxn())
		}
	}

	return router
}
