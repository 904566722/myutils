package strings

import (
	"math/rand"
	"time"
)

// RandStr 生成随机字符串
func RandStr(n int) string {
	// 生成指定长度的随机字符串
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randStr(n, letterBytes)
}

// RandStrLower 生成随机小写字符串
func RandStrLower(n int) string {
	// 生成指定长度的随机字符串
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	return randStr(n, letterBytes)
}

// RandStrUpper 生成随机大写字符串
func RandStrUpper(n int) string {
	// 生成指定长度的随机字符串
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return randStr(n, letterBytes)
}

func randStr(n int, letterBytes string) string {
	// 生成指定长度的随机字符串
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
