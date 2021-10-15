package client

import (
	"net/http"
	"time"
)

func Client() *http.Client {
	// This allows a client and a server to re-use the same underlying TCP connection when sending multiple HTTP Requests/Responses
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
