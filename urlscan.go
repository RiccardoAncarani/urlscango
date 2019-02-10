package urlscango

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, code := AuthenticatedRequest(fmt.Sprintf(RESULT_URL, jobUUID), "GET", nil, APIKey)
	if code == http.StatusNotFound {
		tickerStatus := make(chan *JobResult)
		ticker := time.NewTicker(2 * time.Second)
		go func() {
			for range ticker.C {
				body, code := AuthenticatedRequest(fmt.Sprintf(RESULT_URL, jobUUID), "GET", nil, APIKey)
				if code == http.StatusOK {
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

func GetScreenshot(jobUUID string, fileName string, APIKey string) {
	png, code := AuthenticatedRequest(fmt.Sprintf(SCREENSHOT_URL, jobUUID), "GET", nil, APIKey)
	if code != http.StatusOK {
		panic("Something went wrong: " + string(code))
	} else {
		err := ioutil.WriteFile(fileName, png, 0644)
		if err != nil {
			panic(err)
		}
	}
}
