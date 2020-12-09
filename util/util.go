package util

import (
	"math/rand"
	"net/http"
	"time"
)

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

func DatePassed(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	if y1 == y2 && m1 == m2 && d1 == d2 {
		if date1.Hour() < date2.Hour() && date2.Minute() < date2.Minute() {
			return true
		}
	}

	return false
}

func CodeHttp( code int ) string {
	if code == http.StatusOK {
		return "success"
	} else if code == http.StatusUnprocessableEntity {
		return "unprocessableEntity"
	} else if code == http.StatusNoContent {
		return "empty"
	} else if code == http.StatusBadRequest {
		return "badRequest"
	}

	return "error"
}