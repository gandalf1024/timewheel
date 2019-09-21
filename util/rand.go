package util

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandIntn 获取一个 0 ~ n 之间的随机值
func RandIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(n)
}

// GetRandString 生成n位随机数字字符串
func GetRandString(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(RandIntn(10)))
	}

	return buffer.String()
}

//使用切片生成
func SliceRand(src []interface{}) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))
	return perm
}

//生成n位手机验证码
//n为几位
func RandIntnCode(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(r.Intn(9))
	}
	return str
}

func RandGroup(p ...uint32) int {
	if p == nil {
		panic("args not found")
	}

	r := make([]uint32, len(p))
	for i := 0; i < len(p); i++ {
		if i == 0 {
			r[0] = p[0]
		} else {
			r[i] = r[i-1] + p[i]
		}
	}

	rl := r[len(r)-1]
	if rl == 0 {
		return 0
	}

	rn := uint32(rand.Int63n(int64(rl)))
	for i := 0; i < len(r); i++ {
		if rn < r[i] {
			return i
		}
	}

	panic("bug")
}

func RandInterval(b1, b2 int32) int32 {
	if b1 == b2 {
		return b1
	}

	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	return int32(rand.Int63n(max-min+1) + min)
}

func RandIntervalN(b1, b2 int32, n uint32) []int32 {
	if b1 == b2 {
		return []int32{b1}
	}

	min, max := int64(b1), int64(b2)
	if min > max {
		min, max = max, min
	}
	l := max - min + 1
	if int64(n) > l {
		n = uint32(l)
	}

	r := make([]int32, n)
	m := make(map[int32]int32)
	for i := uint32(0); i < n; i++ {
		v := int32(rand.Int63n(l) + min)

		if mv, ok := m[v]; ok {
			r[i] = mv
		} else {
			r[i] = v
		}

		lv := int32(l - 1 + min)
		if v != lv {
			if mv, ok := m[lv]; ok {
				m[v] = mv
			} else {
				m[v] = lv
			}
		}

		l--
	}

	return r
}

// GenerateRandomAlphaOnlyString generates an alphabetical random string with length n.
func GenerateRandomAlphaOnlyString(n int) string {
	// make a really long string
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
