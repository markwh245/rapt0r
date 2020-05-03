package dirs

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/dtr0x80/rapt0r/pkg/files"
	"github.com/dtr0x80/rapt0r/pkg/parameters"
	"github.com/dtr0x80/rapt0r/pkg/requests"
)

func RunSearchDirectory(domain string, verbose bool, wg *sync.WaitGroup) {
	defer wg.Done()
	request, _ := requests.Get(domain)

	if request.StatusCode == 200 {
		fmt.Printf("[%d] found: %s\n", request.StatusCode, domain)
	} else if verbose {
		fmt.Printf("[%d] error: %s\n", request.StatusCode, domain)
	}
}

func SearchDirectory(parameters *parameters.HostParameters) {
	var domain string
	var wg sync.WaitGroup
	verbose := parameters.Verbose

	checkHttpsPrefix, _ := regexp.MatchString("http.*://", parameters.Host)

	if !checkHttpsPrefix {
		domain = "https://" + parameters.Host
	}

	directorys := files.ReadFile(parameters.Wordlist)

	for _, directory := range directorys {
		domainWithDirectory := domain + "/" + directory

		wg.Add(1)
		go RunSearchDirectory(domainWithDirectory, verbose, &wg)
		wg.Wait()
	}
}
