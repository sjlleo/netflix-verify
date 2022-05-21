package util

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func RequestIP(requrl string, ip string, localAddr string, proxyUrl string) (string, error) {
	var proxy *url.URL
	var err error
	if ip == "" {
		return "", errors.New("IP is empty")
	}
	urlValue, err := url.Parse(requrl)
	if err != nil {
		return "", errors.New("URL parse error")
	}
	host := urlValue.Host
	if ip == "" {
		ip = host
	}
	newrequrl := strings.Replace(requrl, host, ip, 1)
	if proxyUrl != "" {
		if proxy, err = url.Parse(proxyUrl); err != nil {
			return "", errors.New("proxy parse error")
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{ServerName: host},
			Proxy:           http.ProxyURL(proxy),
			DialContext: (&net.Dialer{
				LocalAddr: &net.TCPAddr{
					IP: net.ParseIP(localAddr),
				},
			}).DialContext,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	req, err := http.NewRequest("GET", newrequrl, nil)
	if err != nil {
		return "", errors.New(strings.ReplaceAll(err.Error(), newrequrl, requrl))

	}
	req.Host = host
	req.Header.Set("USER-AGENT", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(strings.ReplaceAll(err.Error(), newrequrl, requrl))
	}
	defer resp.Body.Close()

	Header := resp.Header

	if Header["X-Robots-Tag"] != nil {
		if Header["X-Robots-Tag"][0] == "index" {
			return "us", nil
		}
	}

	if Header["Location"] == nil {
		return "", errors.New("Banned")
	} else {
		return strings.Split(Header["Location"][0], "/")[3], nil
	}
}

func ParseIP(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}
