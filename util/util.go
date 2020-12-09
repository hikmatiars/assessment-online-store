package util

import "math/rand"

func RandomString(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Random( count int ) int {
	return rand.Int() % count
}