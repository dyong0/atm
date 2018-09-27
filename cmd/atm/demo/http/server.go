package main

import (
	"github.com/dyong0/atm/internal/pkg/atm"
	"github.com/dyong0/atm/internal/pkg/atm/http"
	"github.com/gin-gonic/gin"
)

type server struct {
	atm    *atm.ATM
	router *gin.Engine
}

func newServer() *server {
	router := gin.New()

	svr := server{
		router: router,
	}

	http.RegisterAuthRoutes(router, atm)

	return &svr
}

func (s *server) serve(addr ...string) {
	s.router.Run(addr...)
}
