package ims

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Create a http client based on the received configuration.
func (i Config) httpClient() (*http.Client,error) {

	if i.ProxyURL != "" {
		p, err := url.Parse(i.ProxyURL)
		if err != nil {
			return nil, fmt.Errorf("proxy provided but its URL is malformed")
		}
		t := &http.Transport{
			Proxy: http.ProxyURL(p),
		}
		if i.ProxyIgnoreTLS {
			t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
		client := &http.Client{
			Timeout:   30 * time.Second,
			Transport: t,
		}
		return client, nil
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return client, nil
}

