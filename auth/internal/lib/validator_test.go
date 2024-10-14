package lib

import (
	"errors"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		username string
		wantErr  bool
	}{
		{"validUser123", false},
		{"user", false},
		{"us", true}, // too short
		{"", true},   // empty string
		{"thisusernameiswaytoolongtobevalid", true}, // too long
		{"invalid user", true},                      // contains space
		{"invalid@user", true},                      // contains special character
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			err := ValidateUsername(tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email   string
		wantErr error
	}{
		{"", nil},                        // empty string
		{"valid.email@example.com", nil}, // valid email
		{"invalid-email", errors.New("invalid email")},          // invalid email
		{"another.invalid.email@", errors.New("invalid email")}, // invalid email
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil && err.Error() != tt.wantErr.Error()) || (err == nil && tt.wantErr != nil) {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		wantErr  error
	}{
		{"short", errors.New("invalid password: must be longer than 8 characters")},                                            // too short
		{"thispasswordiswaytoolongtobevalidandshouldfail", errors.New("invalid password: must be shorter than 30 characters")}, // too long
		{"validPassword123", nil},     // valid password
		{"anotherValidPassword", nil}, // valid password
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil && err.Error() != tt.wantErr.Error()) || (err == nil && tt.wantErr != nil) {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
