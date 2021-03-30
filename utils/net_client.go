package utils

import (
	"net"
	"net/http"
	"sync"
	"time"
)

var httpOnce sync.Once
var netClient *http.Client

// GetHTTPClient returns the instance of a http.Client
func GetHTTPClient() *http.Client {
	httpOnce.Do(func() {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}

		netClient = &http.Client{
			Timeout:   time.Second * 1,
			Transport: netTransport,
		}

	})
	return netClient
}
