package utils

import "math/rand"

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"

func RandomStringWithLength(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
