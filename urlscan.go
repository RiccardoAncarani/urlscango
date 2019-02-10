package urlscango

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SubmitUrl(URL, public, customAgent, referer, APIKey string) *ScanResponse {
	r := new(ScanRequest)
	r.APIKey = APIKey
	r.URL = URL
	r.Public = public
	r.CustomAgent = customAgent
	r.Referer = referer

	r.Validate()

	response, _ := AuthenticatedRequest(SCAN_URL, "POST", r, r.APIKey)

	resp := new(ScanResponse)
	json.Unmarshal(response, resp)
	return resp
}

func FetchResults(response *ScanResponse, APIKey string) *JobResult {

	jobUUID := response.UUID
	// jobResult := new(JobResult)

	body, code := AuthenticatedRequest(fmt.Sprintf(RESULT_URL, jobUUID), "GET", nil, APIKey)
	if code == http.StatusNotFound {
		tickerStatus := make(chan *JobResult)
		ticker := time.NewTicker(2 * time.Second)
		go func() {
			for t := range ticker.C {
				body, code := AuthenticatedRequest(fmt.Sprintf(RESULT_URL, jobUUID), "GET", nil, APIKey)
				if code == http.StatusOK {
					fmt.Println("LOL", t)

					ticker.Stop()
					jobResult := new(JobResult)
					json.Unmarshal(body, jobResult)
					tickerStatus <- jobResult

				}
			}
		}()
		jobResult := <-tickerStatus
		return jobResult
	} else {
		jobResult := new(JobResult)
		json.Unmarshal(body, jobResult)
		return jobResult
	}
}
