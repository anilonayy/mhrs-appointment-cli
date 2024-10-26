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
			SetRetryWaitTime(5 * time.Second)
	})

	return client
}
