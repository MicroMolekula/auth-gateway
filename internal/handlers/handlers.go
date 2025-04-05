package handlers

import (
	"fmt"
	"github.com/MicroMolekula/auth-gateway/internal/model"
	"net/url"
)

func AuthHandler(host string, urlRequest string, authorization string) (*model.ServiceData, error) {
	switch urlRequest {
	case "/login":
		authorization = ""
	case "/register":
		authorization = ""
	case "/refresh_token":
		authorization = ""
	}
	path, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		return nil, err
	}
	return model.NewServiceData(path, authorization), nil
}

func SymfonyHandler(host string, urlRequest string, authorization string) (*model.ServiceData, error) {
	path, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		return nil, err
	}
	return model.NewServiceData(path, authorization), nil
}

func GPTHandler(host string, urlRequest string, authorization string) (*model.ServiceData, error) {
	path, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		return nil, err
	}
	return model.NewServiceData(path, authorization), nil
}
