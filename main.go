package main

import (
	"flag"

	"github.com/gmdutra/jowfuzz/pkg/banner"
	"github.com/gmdutra/jowfuzz/pkg/dirs"
	"github.com/gmdutra/jowfuzz/pkg/initialize"
	"github.com/gmdutra/jowfuzz/pkg/parameters"
	"github.com/gmdutra/jowfuzz/pkg/subdomains"
	"github.com/gmdutra/jowfuzz/pkg/webcrawler"
)

var (
	host     = flag.String("host", " ", "Define target host")
	port     = flag.Int("port", 80, "Define target port")
	timeout  = flag.Int("timeout", 1, "Define timeout")
	wordlist = flag.String("wordlist", " ", "Define path of wordlist")
	threads  = flag.Int("threads", 1, "Define number of threads")

	subdomain = flag.Bool("subdomain", false, "Search subdomains on host")
	dir       = flag.Bool("dir", false, "Search directorys open on host")
	verbose   = flag.Bool("verbose", false, "Mode verbose")
	crawler   = flag.Bool("crawler", false, "Search URL's on host")
)

func main() {
	banner.Banner()
	flag.Parse()

	params := parameters.HostParameters{
		Host:     *host,
		Port:     *port,
		Timeout:  *timeout,
		Verbose:  *verbose,
		Wordlist: *wordlist,
	}

	initialize.Init(&params)

	switch {

	case *subdomain:
		subdomains.SearchSubdomains(&params)
		break

	case *dir:
		dirs.SearchDirectory(&params)
		break

	case *crawler:
		webcrawler.Crawler(&params)
		break
	}

}
