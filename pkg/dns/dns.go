package dns

import (
	"net"
)

const defaultPort = "80"

func ResolverDns(domain string) (ipv4 *net.TCPAddr, ipv6 *net.TCPAddr) {
	host := domain + ":" + defaultPort
	ipv4, _ = net.ResolveTCPAddr("tcp4", host)
	ipv6, _ = net.ResolveTCPAddr("tcp6", host)

	return ipv4, ipv6
}
