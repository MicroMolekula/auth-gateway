package utils

import (
	"net/url"
	"strings"
)

func CutLocationFromUrl(url *url.URL, location string) (*url.URL, error) {
	partsUrl := strings.Split(url.String(), "/")
	newPartsUrl := make([]string, 0, len(partsUrl)-1)
	for _, partUrl := range partsUrl {
		if partUrl != strings.Trim(location, "/") {
			newPartsUrl = append(newPartsUrl, partUrl)
		}
	}
	return url.Parse(strings.Join(newPartsUrl, "/"))
}
