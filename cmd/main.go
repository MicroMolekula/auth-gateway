package main

import (
	"github.com/MicroMolekula/auth-gateway/internal/config"
	"github.com/MicroMolekula/auth-gateway/internal/handlers"
	"github.com/MicroMolekula/auth-gateway/internal/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	config.Cfg = cfg
	if err != nil {
		log.Fatal(err)
	}
	symfonyProxy := server.NewProxy(cfg.Domains["symfony"], "/symfony/", handlers.SymfonyHandler)
	authProxy := server.NewProxy(cfg.Domains["auth"], "/auth/", handlers.AuthHandler)
	gptProxy := server.NewProxy(cfg.Domains["go"], "/fitness/", handlers.GPTHandler)
	serv := server.NewServer(cfg)
	serv.AddProxy(symfonyProxy)
	serv.AddProxy(authProxy)
	serv.AddProxy(gptProxy)
	serv.HandleProxies()
	serv.Start()
}
