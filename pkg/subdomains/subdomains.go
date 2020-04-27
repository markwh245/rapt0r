package subdomains

import (
	"fmt"
	"os"
	"log"
	"regexp"
	"sync"

	"github.com/gmdutra/jowfuzz/pkg/files"
	"github.com/gmdutra/jowfuzz/pkg/dns"
	"github.com/gmdutra/jowfuzz/pkg/parameters"
)

func RunSearchdomains(domain string, verbose bool, wg *sync.WaitGroup){
	defer wg.Done()

	ipv4, _ := dns.ResolverDns(domain)

	if ipv4 != nil {
		fmt.Printf("[+] found: %s \t\t %s\n", domain, ipv4)
	} else if verbose {
		fmt.Printf("[+] error: %s \t\t %s\n", domain, ipv4)
	}
}

func SearchSubdomains(parameters *parameters.HostParameters) {
	checkHttpsPrefix, _ := regexp.MatchString("http.*://", parameters.Host)
	verbose := parameters.Verbose

	var wg sync.WaitGroup

	if checkHttpsPrefix {
		log.Fatalf("[+] Error: bad format url, please, remove http[s]")
		os.Exit(1)
	}

	subdomains := files.ReadFile(parameters.Wordlist)

	for _, subdomain := range subdomains {
		domain := subdomain + "." + parameters.Host

		wg.Add(1)
		go RunSearchdomains(domain, verbose, &wg)
		wg.Wait()
	}
}
