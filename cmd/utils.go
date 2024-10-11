package cmd

import (
	"context"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

func createHttpClient(cmd *cobra.Command) *http.Client {
	dnsResolverAddr, _ := cmd.Flags().GetString("dns-resolver")
	if dnsResolverAddr == "" {
		dnsResolverAddr = "1.1.1.1"
	}
	dnsResolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", dnsResolverAddr+":53")
		},
	}
	dialer := &net.Dialer{
		Resolver: dnsResolver,
	}
	jar := http.CookieJar(nil)
	return &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}
}
