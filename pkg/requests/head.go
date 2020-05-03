package requests

import (
	"net/http"
	"time"

	"github.com/dtr0x80/rapt0r/pkg/parameters"
)

func GetHeader(parameters *parameters.HostParameters) http.Header {
	host := parameters.Host

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, _ := client.Head(host)
	header := resp.Header

	return header
}
