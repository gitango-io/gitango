package main

import (
	"log"
	"net/http"

	"github.com/gitango-io/gitango/client"
	"github.com/gitango-io/gitango/method"
)

func main() {
	c := client.Client()
	response := method.SendRequest(c, "users/meanii", map[string]string{}, http.MethodGet)
	log.Println("Response Body:", string(response))
}
