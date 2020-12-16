package restclient

import (
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	response, err := Request("https://lenta.ru/rss/top7")

	log.Println(response)

	if err != nil {
		t.Error(err)
	}
}
