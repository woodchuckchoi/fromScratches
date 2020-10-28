package util

import (
	"math/rand"
	"strings"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettersLength = 46
)

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(lettersLength)]
	}

	return string(b)
}

func ToSqlTimeStamp(s ...*string) {
	for i := range s {
		if *s[i] == "" {
			continue
		}
		// original s == yyyy-mm-dd_hh:mm:ss
		// output s == "yyyy-mm-dd hh:mm:ss"
		*s[i] = strings.Replace(*s[i], "_", " ", 1)
		*s[i] = "\"" + *s[i] + "\""
	}
}
