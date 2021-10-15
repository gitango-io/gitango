package method

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendRequest(client *http.Client, endpoint string, values map[string]string, method string, ) []byte {
	
	baseURI := "https://api.github.com/" + endpoint 
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatalf("Error Occurred in jsonData %v", err)
	}

	req, err := http.NewRequest(method, baseURI, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}
