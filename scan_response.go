package urlscango

type ScanResponse struct {
	Message    string `json:"message"`
	UUID       string `json:"uuid"`
	Result     string `json:"result"`
	API        string `json:"api"`
	Visibility string `json:"visibility"`
	Options    struct {
		Useragent string `json:"useragent"`
	} `json:"options"`
	URL string `json:"url"`
}
