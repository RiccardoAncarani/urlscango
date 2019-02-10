package urlscango

type ScanRequest struct {
	URL         string `json:"url"`
	Public      string `json:"public,omitempty"`
	CustomAgent string `json:"customagent,omitempty"`
	Referer     string `json:"referer,omitempty"`
	APIKey      string `json:"-"`
}

func (r ScanRequest) Validate() {
	switch r.Public {
	case "on":
	case "off":
	default:
		panic("Public field should be ON or OFF!")
	}
}
