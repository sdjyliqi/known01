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
			Proxy:       nil,
			DialContext: nil,
			Dial: (&net.Dialer{
				Timeout:       3 * time.Second,
				Deadline:      time.Time{},
				LocalAddr:     nil,
				DualStack:     false,
				FallbackDelay: 0,
				KeepAlive:     30 * time.Second,
				Resolver:      nil,
				Cancel:        nil,
				Control:       nil,
			}).Dial,
			DialTLSContext:         nil,
			DialTLS:                nil,
			TLSClientConfig:        nil,
			TLSHandshakeTimeout:    3 * time.Second,
			DisableKeepAlives:      false,
			DisableCompression:     false,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    0,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        0,
			ResponseHeaderTimeout:  3 * time.Second,
			ExpectContinueTimeout:  1 * time.Second,
			TLSNextProto:           nil,
			ProxyConnectHeader:     nil,
			MaxResponseHeaderBytes: 0,
			WriteBufferSize:        0,
			ReadBufferSize:         0,
			ForceAttemptHTTP2:      false,
		}

		netClient = &http.Client{
			Transport:     netTransport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second * 1,
		}

	})
	return netClient
}
