package semaphore

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Semaphore struct {
	url    *url.URL
	client *http.Client
}

func New(addr string, dns string) *Semaphore {
	addrUrl, err := url.Parse(addr)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
		},
	}
	return &Semaphore{
		url:    addrUrl,
		client: client,
	}
}

func (s *Semaphore) Login(username string, password string) error {
	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: s.url.Scheme, Host: s.url.Host, Path: "/api/auth/login"},
		Header: http.Header{"Content-Type": []string{"application/json"}},
		Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"auth": "%s", "password": "%s"}`, username, password))),
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("login failed: %s", resp.Status)
	} else {
		fmt.Println("login successful")
	}
	return nil
}

func (s *Semaphore) Backup(projectID string) error {
	return nil
}

func (s *Semaphore) Restore(projectID string) error {
	return nil
}
