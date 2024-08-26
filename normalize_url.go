package main

import (
	"fmt"
	"net/url"
)

func NormalizeURL(urlStr string) (string, error) {
	usrStruct, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	if usrStruct.Scheme != "http" && usrStruct.Scheme != "https" {
		return "", fmt.Errorf("invalid URL: %s", urlStr)
	}

	normUrl := fmt.Sprintf("%s%s", usrStruct.Host, usrStruct.Path)

	return normUrl, nil
}
