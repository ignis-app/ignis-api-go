package util

import (
	"time"
	"os"
)

var INCR int64 = 0

func Snowflake(workerId int) int64 {
	v := time.Now().UnixMilli() << 22 + int64(workerId) << 17 + int64(os.Getpid()) << 12 + INCR
	INCR = (INCR + 1) % 4096
	return v
}
