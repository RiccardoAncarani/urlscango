package urlscango

import (
	"encoding/json"
	"fmt"
)

func SubmitUrl(URL, public, customAgent, referer, APIKey string) {
	r := new(Request)
	r.APIKey = APIKey
	r.URL = URL
	r.Public = public
	r.CustomAgent = customAgent
	r.Referer = referer

	r.Validate()

	response, _ := AuthenticatedRequest(SCAN_URL, r, r.APIKey)

	resp := new(Response)
	json.Unmarshal(response, resp)
	fmt.Println(resp)

}

func FetchResults(JobID string) {

}
