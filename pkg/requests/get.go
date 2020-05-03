package requests

import (
	"log"
	"net/http"

	"github.com/dtr0x80/jowfuzz/pkg/random"
)

func Get(host string) (r *http.Response, err error) {
	randomUserAgent := random.RandomUserAgent()

	req, err := http.NewRequest("GET", host, nil)

	if err != nil {
		log.Fatal("Error reading request", err)
	}

	req.Header.Set("Host", host)
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", randomUserAgent)

	client := &http.Client{}

	resp, err := client.Do(req)
	return resp, err
}
