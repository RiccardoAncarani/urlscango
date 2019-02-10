package urlscango

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AuthenticatedRequest(URL string, method string, request interface{}, APIKey string) ([]byte, int) {

	jsonBody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req := new(http.Request)
	if method == "POST" {
		req, err = http.NewRequest(method, URL, bytes.NewReader(jsonBody))
	} else {
		req, err = http.NewRequest(method, URL, nil)
	}
	req.Header.Add("API-Key", APIKey)
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	/* if resp.StatusCode != http.StatusOK {
		panic("Something went wrong, received HTTP status code: " + resp.Status)
	} */

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, resp.StatusCode
}
