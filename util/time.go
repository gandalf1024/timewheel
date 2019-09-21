package util

import (
	"time"
)

func GetTime() int64 {
	return time.Now().Unix()
}

func GetTimeAfterSecond(s int) int64 {
	now := time.Now()
	now = now.Add(time.Second * time.Duration(s))
	return time.Now().Unix()
}
