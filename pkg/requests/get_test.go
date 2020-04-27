package requests

import (
	"testing"
)

func TestHttpGet(t *testing.T){
	_, err := Get("https://google.com")

	if err != nil {
		t.Errorf("[+] Testing failed")
	}
}