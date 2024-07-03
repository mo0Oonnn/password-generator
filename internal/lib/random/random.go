package random

import (
	"time"

	"math/rand"
)

func CreateRandomPassword(size int, useSpecialChars bool) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]rune, size)
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if useSpecialChars {
		specialChars := []rune(`!@#$%^&*()_+{}|:<>?,.;[]`)
		for i := 0; i < size; i++ {
			if i%2 != 0 {
				password[i] = chars[random.Intn(len(chars))]
			} else {
				password[i] = specialChars[random.Intn(len(specialChars))]
			}
		}
	} else {
		for i := range password {
			password[i] = chars[random.Intn(len(chars))]
		}
	}
	return string(password)
}
