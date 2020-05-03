package main

import (
	"flag"

	"github.com/dtr0x80/rapt0r/pkg/banner"
	"github.com/dtr0x80/rapt0r/pkg/dirs"
	"github.com/dtr0x80/rapt0r/pkg/initialize"
	"github.com/dtr0x80/rapt0r/pkg/parameters"
	"github.com/dtr0x80/rapt0r/pkg/subdomains"
	"github.com/dtr0x80/rapt0r/pkg/webcrawler"
)

var (
	host     = flag.String("host", " ", "Define target host")
	port     = flag.Int("port", 80, "Define target port")
	timeout  = flag.Int("timeout", 1, "Define timeout")
	wordlist = flag.String("wordlist", " ", "Define path of wordlist")

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

	case *dir:
		dirs.SearchDirectory(&params)

	case *crawler:
		webcrawler.Crawler(&params)
	}
}
