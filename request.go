package urlscango

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AuthenticatedRequest(URL string, request interface{}, APIKey string) ([]byte, string) {

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewReader(jsonBody))
	req.Header.Add("API-Key", APIKey)
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, string(body)
}
