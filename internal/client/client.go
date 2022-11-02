package client

import (
	"errors"
	"fmt"
	"log"

	api "github.com/zededa/zedcloud-api"
)

const defaultZedcloudUrl = "https://zedcontrol.zededa.net"

type Client struct {
	api.Client
	zedcloudURL string // https://zedcontrol.zededa.net
	token       string // API Token
	username    string
	password    string
}

// token or username and password must be specified.
func NewClient(zedcloudUrl, token, username, password string) (*api.Client, error) {
	if zedcloudUrl == "" {
		zedcloudUrl = defaultZedcloudUrl
	}
	if token != "" {
		username = ""
		password = ""
		log.Println("enable authentication by token")
	} else if username == "" && password == "" {
		return nil, errors.New("require either token or username and password")
	}
	apiClient, err := api.NewClient(zedcloudUrl, token, username, password)
	if err != nil {
		return nil, fmt.Errorf("could not initialize api client: %w", err)
	}
	return apiClient, nil
}
