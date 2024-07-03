package random_test

import (
	"testing"

	"github.com/mo0Oonnn/password-generator/internal/lib/random"
)

func TestCreateRandomPassword(t *testing.T) {
	password := random.CreateRandomPassword(10, false)
	if len(password) != 10 {
		t.Errorf("invalid password length")
	}
	if password == "" {
		t.Errorf("password is empty")
	}
}
