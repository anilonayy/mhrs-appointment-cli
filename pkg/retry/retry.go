package retry

import (
	"github.com/avast/retry-go"
	"time"
)

func Do(callback func() error) error {
	return retry.Do(callback,
		retry.Attempts(3),
		retry.Delay(5*time.Second),
		retry.LastErrorOnly(true))
}
