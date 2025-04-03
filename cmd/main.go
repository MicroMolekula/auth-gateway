package main

import (
	"github.com/MicroMolekula/auth-gateway/internal/config"
	"github.com/MicroMolekula/auth-gateway/internal/handlers"
	"github.com/MicroMolekula/auth-gateway/internal/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	symfonyProxy := server.NewProxy(cfg.Domains["symfony"], "/symfony/", handlers.SymfonyHandler)
	authProxy := server.NewProxy(cfg.Domains["auth"], "/auth/", handlers.AuthHandler)
	serv := server.NewServer(cfg)
	serv.AddProxy(symfonyProxy)
	serv.AddProxy(authProxy)
	serv.HandleProxies()
	serv.Start()
}
