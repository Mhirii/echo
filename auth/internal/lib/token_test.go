package lib

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenAccessToken(t *testing.T) {
	// Set up environment variable for duration
	os.Setenv("JWT_ACCESS_DURATION_MIN", "60")
	defer os.Unsetenv("JWT_ACCESS_DURATION_MIN")

	// Mock secret key
	jwt_secret = "mySecretKey"

	// Test case: Valid token generation
	t.Run("ValidTokenGeneration", func(t *testing.T) {
		payload := &AccessTokenPayload{
			Username: "testuser",
			ID:       "12345",
		}
		tokenString, err := payload.GenAccessToken()
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		// Parse the token to verify claims
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})
		assert.NoError(t, err)
		assert.True(t, token.Valid)

		claims := token.Claims.(jwt.MapClaims)
		assert.Equal(t, "testuser", claims["username"])
		assert.Equal(t, "12345", claims["id"])
	})

	// Test case: Invalid duration
	t.Run("InvalidDuration", func(t *testing.T) {
		os.Setenv("JWT_ACCESS_DURATION_MIN", "invalid")
		defer os.Setenv("JWT_ACCESS_DURATION_MIN", "60")

		payload := &AccessTokenPayload{
			Username: "testuser",
			ID:       "12345",
		}
		tokenString, err := payload.GenAccessToken()
		assert.Error(t, err)
		assert.Empty(t, tokenString)
	})

	// Test case: Empty data
	t.Run("EmptyData", func(t *testing.T) {
		payload := &AccessTokenPayload{
			Username: "testuser",
			ID:       "12345",
		}
		tokenString, err := payload.GenAccessToken()
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		// Parse the token to verify claims
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})
		assert.NoError(t, err)
		assert.True(t, token.Valid)

		claims := token.Claims.(jwt.MapClaims)
		assert.Equal(t, "testuser", claims["username"])
		assert.Equal(t, "12345", claims["id"])
		assert.Empty(t, claims["data"])
	})
}

func TestGenRefreshToken(t *testing.T) {
	// Set up environment variable for duration
	os.Setenv("JWT_REFRESH_DURATION_MIN", "60")
	defer os.Unsetenv("JWT_REFRESH_DURATION_MIN")

	// Mock secret key
	jwt_secret = "mySecretKey"

	// Test case: Valid token generation
	t.Run("ValidTokenGeneration", func(t *testing.T) {
		payload := &RefreshTokenPayload{
			Username: "testuser",
			ID:       "12345",
		}
		tokenString, err := payload.GenRefershToken()
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		// Parse the token to verify claims
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})
		assert.NoError(t, err)
		assert.True(t, token.Valid)

		claims := token.Claims.(jwt.MapClaims)
		assert.Equal(t, "testuser", claims["username"])
		assert.Equal(t, "12345", claims["id"])
	})

	// Test case: Invalid duration
	t.Run("InvalidDuration", func(t *testing.T) {
		os.Setenv("JWT_REFRESH_DURATION_MIN", "invalid")
		defer os.Setenv("JWT_REFRESH_DURATION_MIN", "60")

		payload := &RefreshTokenPayload{
			Username: "testuser",
			ID:       "12345",
		}
		tokenString, err := payload.GenRefershToken()
		assert.Error(t, err)
		assert.Empty(t, tokenString)
	})
}

func TestSetCommonClaims(t *testing.T) {
	// Test case: Valid claims setting
	t.Run("ValidClaimsSetting", func(t *testing.T) {
		claims := jwt.MapClaims{}
		username := "testuser"
		id := "12345"
		duration := 60

		setCommonClaims(claims, username, id, duration)

		assert.Equal(t, username, claims["username"])
		assert.Equal(t, id, claims["id"])

		iat := claims["iat"].(int64)
		exp := claims["exp"].(int64)
		now := time.Now().Unix()

		assert.True(t, iat <= now && iat > now-5) // Allowing a small margin for execution time
		assert.Equal(t, exp, now+int64(duration*60))
	})

	// Test case: Check expiration time
	t.Run("CheckExpirationTime", func(t *testing.T) {
		claims := jwt.MapClaims{}
		username := "testuser"
		id := "12345"
		duration := 30

		setCommonClaims(claims, username, id, duration)

		exp := claims["exp"].(int64)
		expectedExp := time.Now().Add(time.Duration(duration) * time.Minute).Unix()

		assert.Equal(t, expectedExp, exp)
	})
}

func TestParseAccessToken(t *testing.T) {
	// Mock secret key
	jwt_secret = "mySecretKey"

	// Helper function to generate a valid token
	generateValidToken := func() string {
		claims := jwt.MapClaims{
			"id":       "12345",
			"username": "testuser",
			"data":     map[string]interface{}{"role": "admin"},
			"exp":      time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(jwt_secret))
		return tokenString
	}

	// Test case: Valid token
	t.Run("ValidToken", func(t *testing.T) {
		tokenString := generateValidToken()
		payload, err := ParseAccessToken(tokenString)
		assert.NoError(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, "12345", payload.ID)
		assert.Equal(t, "testuser", payload.Username)
		assert.Equal(t, "admin", payload.Data["role"])
	})

	// Test case: Invalid token
	t.Run("InvalidToken", func(t *testing.T) {
		tokenString := "invalidTokenString"
		payload, err := ParseAccessToken(tokenString)
		assert.Error(t, err)
		assert.Nil(t, payload)
	})

	// Test case: Malformed token
	t.Run("MalformedToken", func(t *testing.T) {
		tokenString := "malformed.token.string"
		payload, err := ParseAccessToken(tokenString)
		assert.Error(t, err)
		assert.Nil(t, payload)
	})
}

func TestParseRefreshToken(t *testing.T) {
	// Mock secret key
	jwt_secret = "mySecretKey"

	// Helper function to generate a valid token
	generateValidToken := func() string {
		claims := jwt.MapClaims{
			"id":       "12345",
			"username": "testuser",
			"exp":      time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(jwt_secret))
		return tokenString
	}

	// Test case: Valid token
	t.Run("ValidToken", func(t *testing.T) {
		tokenString := generateValidToken()
		payload, err := ParseRefreshToken(tokenString)
		assert.NoError(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, "12345", payload.ID)
		assert.Equal(t, "testuser", payload.Username)
	})

	// Test case: Invalid token
	t.Run("InvalidToken", func(t *testing.T) {
		tokenString := "invalidTokenString"
		payload, err := ParseRefreshToken(tokenString)
		assert.Error(t, err)
		assert.Nil(t, payload)
	})

	// Test case: Malformed token
	t.Run("MalformedToken", func(t *testing.T) {
		tokenString := "malformed.token.string"
		payload, err := ParseRefreshToken(tokenString)
		assert.Error(t, err)
		assert.Nil(t, payload)
	})
}

func TestValidateToken(t *testing.T) {
	// Mock secret key
	jwt_secret = "mySecretKey"

	// Helper function to generate a valid token
	generateValidToken := func() string {
		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(jwt_secret))
		return tokenString
	}

	// Helper function to generate an expired token
	generateExpiredToken := func() string {
		claims := jwt.MapClaims{
			"exp": time.Now().Add(-time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(jwt_secret))
		return tokenString
	}

	// Test case: Valid token
	t.Run("ValidToken", func(t *testing.T) {
		tokenString := generateValidToken()
		err := ValidateToken(tokenString)
		assert.NoError(t, err)
	})

	// Test case: Invalid token
	t.Run("InvalidToken", func(t *testing.T) {
		tokenString := "invalidTokenString"
		err := ValidateToken(tokenString)
		assert.Error(t, err)
	})

	// Test case: Malformed token
	t.Run("MalformedToken", func(t *testing.T) {
		tokenString := "malformed.token.string"
		err := ValidateToken(tokenString)
		assert.Error(t, err)
	})

	// Test case: Expired token
	t.Run("ExpiredToken", func(t *testing.T) {
		tokenString := generateExpiredToken()
		err := ValidateToken(tokenString)
		assert.Error(t, err)
	})
}
