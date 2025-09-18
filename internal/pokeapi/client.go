package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpCLient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpCLient: http.Client{
			Timeout: timeout,
		},
	}
}
