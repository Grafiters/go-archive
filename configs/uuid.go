package configs

import "math/rand"

type UID string

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Generate(identifier string) string {
	var length int = 5

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return identifier + string(b)
}