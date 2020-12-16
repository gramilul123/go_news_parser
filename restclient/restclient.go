package restclient

import (
	"github.com/go-resty/resty/v2"
)

// Request return response from rss stream
func Request(url string) ([]byte, error) {
	var err error
	client := resty.New()

	response, err := client.R().Get(url)

	bytes := []byte(response.Body())

	return bytes, err
}
