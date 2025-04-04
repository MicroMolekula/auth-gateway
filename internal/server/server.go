package server

import (
	"github.com/MicroMolekula/auth-gateway/internal/config"
	"log"
	"net/http"
)

type Server struct {
	Host    string
	Port    string
	proxies []*Proxy
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Host:    cfg.Server.Host,
		Port:    cfg.Server.Port,
		proxies: make([]*Proxy, 0, 3),
	}
}

func (s *Server) AddProxy(proxy *Proxy) {
	s.proxies = append(s.proxies, proxy)
}

func (s *Server) HandleProxies() {
	for _, proxy := range s.proxies {
		proxy.Handle()
	}
}

func (s *Server) Start() {
	if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
		log.Fatal(err)
	}
}
