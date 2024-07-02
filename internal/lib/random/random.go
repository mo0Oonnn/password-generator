package random

import (
	"time"

	"math/rand"
)

func CreateRandomPassword(size int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	password := make([]rune, size)
	for i := range password {
		password[i] = chars[random.Intn(len(chars))]
	}
	return string(password)
}
