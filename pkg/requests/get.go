package requests

import (
	"net/http"
)

func Get(host string) (r *http.Response, err error) {
	resp, err := http.Get(host)

	return resp, err
}
