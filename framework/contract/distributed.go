package contract

import (
	"time"
)

const DistributedKey = "zima:distributed"

type Distributed interface {
	Select(serviceName string, appID string, lockTime time.Duration) (string, error)
}
