package initialize

import (
	"fmt"
	"time"

	"github.com/dtr0x80/rapt0r/pkg/parameters"
	"github.com/dtr0x80/rapt0r/pkg/dns"
)

func Init(p *parameters.HostParameters) {
	host := p.Host
	port := p.Port
	currentTime := time.Now().Format(time.RFC850)

	ipv4, ipv6 := dns.ResolverDns(host)

	fmt.Printf(`
[+] Host IPv4: %s
[+] Host IPv6: %s
[+] Hostname: %s
[+] Port: %d
[+] Start time: %s
----------------------------------------------------------------------------
	`, ipv4, ipv6, host, port, currentTime)

	fmt.Printf("\n")
}
