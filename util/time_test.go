package util

import (
	"fmt"
	"testing"
)

func Test_GetTime(t *testing.T) {
	cc := GetTime()
	fmt.Println(cc)
}

func Test_GetTimeAfterScend(t *testing.T) {
	vv := GetTimeAfterSecond(300)
	fmt.Println(vv)
}
