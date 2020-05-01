package subdomains

import (
	"testing"
	"github.com/dtr0x80/jowfuzz/pkg/parameters"

)

func TestSearchSubdomains(t *testing.T){
	host := "google.com"
	wordlist := "../../databases/subdomains/subdomains.txt"

	parms := parameters.HostParameters{}

	parms.Host = host
	parms.Wordlist = wordlist
	
	SearchSubdomains(&parms)
}