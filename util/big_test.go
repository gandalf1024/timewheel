package util

import (
	"fmt"
	"testing"
)

func TestBig_ParseBig256(t *testing.T) {
	big, flag := ParseBig256("123456789123456789012345678912345678123456789123456789012345678912345678")
	fmt.Println(big, flag)
}

func TestBig_MustParseBig256(t *testing.T) {
	big := MustParseBig256("123456789123456789012345678912345678123456789123456789012345678912345678")
	fmt.Println(big)
}

func TestBig_BigPow(t *testing.T) {
	big := BigPow(2, 23)
	fmt.Println(big)
}

func Test_BigMax(t *testing.T) {
	big1 := MustParseBig256("12")
	big2 := MustParseBig256("34")

	result := BigMax(big1, big2)
	fmt.Println(result)
}
