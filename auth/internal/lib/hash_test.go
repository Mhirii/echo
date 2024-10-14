package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	// Original test cases
	secret := "mySecretPassword"
	hashed, err := Hash(secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(secret))
	assert.NoError(t, err)

	emptySecret := ""
	hashed, err = Hash(emptySecret)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(emptySecret))
	assert.NoError(t, err)

	// Long password exceeding max length of 72 characters
	longSecret := "aVeryLongPasswordThatExceedsNormalLengthAndShouldStillBeHandledProperlyByTheHashFunction"
	hashed, err = Hash(longSecret)
	assert.Error(t, err)
	assert.Empty(t, hashed)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(longSecret))
	assert.Error(t, err)

	// Max length password of 72 characters
	maxLengthSecret := "PasswordWithSeventyTwoCharactersShouldBeHandledProperlyByTheHashFunction"
	hashed, err = Hash(maxLengthSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(maxLengthSecret))
	assert.NoError(t, err)

	// Special characters
	specialCharSecret := "!@#$%^&*()_+-=[]{}|;':,.<>/?`~"
	hashed, err = Hash(specialCharSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(specialCharSecret))
	assert.NoError(t, err)

	// Whitespace characters
	whitespaceSecret := " \t\n\r"
	hashed, err = Hash(whitespaceSecret)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(whitespaceSecret))
	assert.NoError(t, err)
}

func TestComparePassword(t *testing.T) {
	hashPassword := func(password string) string {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(hashed)
	}

	// Test case: Valid password
	t.Run("ValidPassword", func(t *testing.T) {
		plain := "mySecretPassword"
		hashed := hashPassword(plain)
		err := ComparePassword(hashed, plain)
		assert.NoError(t, err)
	})

	// Test case: Invalid password
	t.Run("InvalidPassword", func(t *testing.T) {
		plain := "mySecretPassword"
		hashed := hashPassword(plain)
		err := ComparePassword(hashed, "wrongPassword")
		assert.Error(t, err)
	})

	// Test case: Empty plain password
	t.Run("EmptyPlainPassword", func(t *testing.T) {
		plain := ""
		hashed := hashPassword("mySecretPassword")
		err := ComparePassword(hashed, plain)
		assert.Error(t, err)
	})

	// Test case: Empty hashed password
	t.Run("EmptyHashedPassword", func(t *testing.T) {
		plain := "mySecretPassword"
		hashed := ""
		err := ComparePassword(hashed, plain)
		assert.Error(t, err)
	})

	// Test case: Both passwords empty
	t.Run("BothPasswordsEmpty", func(t *testing.T) {
		plain := ""
		hashed := ""
		err := ComparePassword(hashed, plain)
		assert.Error(t, err)
	})
}
