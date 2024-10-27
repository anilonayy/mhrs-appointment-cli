package resty

import (
	"github.com/go-resty/resty/v2"
	"sync"
	"time"
)

var (
	client *resty.Client
	once   sync.Once
)

func GetClient() *resty.Client {
	once.Do(func() {
		client = resty.New().
			SetTimeout(5 * time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(5 * time.Second).
			SetHeaders(map[string]string{
				"accept":                           "application/json, text/plain, */*",
				"accept-language":                  "tr-TR",
				"access-control-allow-credentials": "true",
				"access-control-allow-headers":     "Authorization,Content-Type, Accept, X-Requested-With, remember-me",
				"access-control-allow-methods":     "DELETE, POST, GET, OPTIONS",
				"access-control-allow-origin":      "*",
				"access-control-max-age":           "3600",
				"content-type":                     "application/json",
				"sec-ch-ua":                        "\"Chromium\";v=\"130\", \"Google Chrome\";v=\"130\", \"Not?A_Brand\";v=\"99\"",
				"sec-ch-ua-mobile":                 "?0",
				"sec-ch-ua-platform":               "\"macOS\"",
			})
	})

	return client
}
