package semaphore

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Semaphore struct {
	Url        *url.URL
	DnsResolv  *net.Resolver
	HttpClient *http.Client
}

func New(addr string, dns string) (*Semaphore, error) {
	addrUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	dnsResolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", dns+":53")
		},
	}
	dialer := &net.Dialer{
		Resolver: dnsResolver,
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}
	return &Semaphore{
		Url:        addrUrl,
		HttpClient: client,
	}, nil
}

func (s *Semaphore) Read(path string) (*http.Response, error) {
	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: s.Url.Scheme, Host: s.Url.Host, Path: path},
		Header: http.Header{"Content-Type": []string{"application/json"}},
	}
	return s.HttpClient.Do(req)
}

func (s *Semaphore) Write(path string, payload interface{}) (*http.Response, error) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: s.Url.Scheme, Host: s.Url.Host, Path: path},
		Header: http.Header{"Content-Type": []string{"application/json"}},
		Body:   io.NopCloser(strings.NewReader(string(payloadJson))),
	}
	return s.HttpClient.Do(req)
}
