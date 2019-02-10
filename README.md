# urlscanGO
A super simple interface for urlscan.io, written in Go.

## Usage
This example shows the main functionalities of this package:

```go
package main

import (
	"github.com/RiccardoAncarani/urlscango"
)

func main() {
	apiKey := "<API KEY>" // the API key that you have to generate creating an account

    // URL submission
	r := urlscango.SubmitUrl(
		"https://www.google.it",
		"off",
		"",
		"",
		apiKey,
	)

    // This function will poll the server every 2 seconds to get the job result
    // This will return every info of the page such as subdomains, external requests and so on
	urlscango.FetchResults(
		r,
		apiKey,
	)

    // Once the job is completed we can extract screenshots
	urlscango.GetScreenshot(
		r.UUID,
		"google.png",
		apiKey,
	)

}

```