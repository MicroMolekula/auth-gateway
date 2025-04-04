package server

import (
	"encoding/json"
	"github.com/MicroMolekula/auth-gateway/internal/config"
	"github.com/MicroMolekula/auth-gateway/internal/model"
	"github.com/MicroMolekula/auth-gateway/internal/utils"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	host     string
	location string
	handler  func(host string, urlRequest string) (*model.ServiceData, error)
}

func NewProxy(domain string, location string, handler func(host string, urlRequest string) (*model.ServiceData, error)) *Proxy {
	return &Proxy{
		host:     domain,
		location: location,
		handler:  handler,
	}
}

func (p *Proxy) Handle() {
	http.HandleFunc(p.location, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.Cfg.CORS)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		serviceData, err := p.handler(p.host, r.URL.String())
		if err != nil {
			SendErrorJSON(err, w)
			return
		}
		r.URL, err = utils.CutLocationFromUrl(r.URL, p.location)
		if err != nil {
			SendErrorJSON(err, w)
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(serviceData.Url)
		r.Header.Set("Authorization", serviceData.Authorization)
		proxy.ServeHTTP(w, r)
	})
}

func SendErrorJSON(err error, w http.ResponseWriter) {
	errResponse, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(errResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
