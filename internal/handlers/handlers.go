package handlers

import (
	"fmt"
	"github.com/MicroMolekula/auth-gateway/internal/model"
	"net/url"
)

func AuthHandler(host string, urlRequest string) (*model.ServiceData, error) {
	path, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		return nil, err
	}
	return model.NewServiceData(path, "Bearer "), nil
}

func SymfonyHandler(host string, urlRequest string) (*model.ServiceData, error) {
	path, err := url.Parse(fmt.Sprintf("http://%s/", host))
	if err != nil {
		return nil, err
	}
	return model.NewServiceData(path, "Bearer "), nil
}
