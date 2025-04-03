package model

import "net/url"

type ServiceData struct {
	Url           *url.URL
	Authorization string
}

func NewServiceData(url *url.URL, authorization string) *ServiceData {
	return &ServiceData{
		Url:           url,
		Authorization: authorization,
	}
}
