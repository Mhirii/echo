package database

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestToMapInterface(t *testing.T) {
	email := "john.doe@example.com"
	phone := "123-456-7890"
	uuid := uuid.New()

	tests := []struct {
		user     User
		expected map[string]interface{}
	}{
		{
			user: User{
				ID:        uuid,
				AccountID: "456",
				Username:  "johndoe",
				FirstName: "John",
				LastName:  "Doe",
				Email:     &email,
				Phone:     &phone,
			},
			expected: map[string]interface{}{
				"id":         uuid,
				"account_id": "456",
				"username":   "johndoe",
				"first_name": "John",
				"last_name":  "Doe",
				"email":      "john.doe@example.com",
				"phone":      "123-456-7890",
			},
		},
		{
			user: User{
				ID:        uuid,
				AccountID: "456",
				Username:  "johndoe",
				FirstName: "",
				LastName:  "Doe",
				Email:     &email,
				Phone:     &phone,
			},
			expected: map[string]interface{}{
				"id":         uuid,
				"account_id": "456",
				"username":   "johndoe",
				"last_name":  "Doe",
				"email":      "john.doe@example.com",
				"phone":      "123-456-7890",
			},
		},
		{
			user: User{
				ID:        uuid,
				AccountID: "456",
				Username:  "johndoe",
				FirstName: "John",
				LastName:  "",
				Email:     &email,
				Phone:     &phone,
			},
			expected: map[string]interface{}{
				"id":         uuid,
				"account_id": "456",
				"username":   "johndoe",
				"first_name": "John",
				"email":      "john.doe@example.com",
				"phone":      "123-456-7890",
			},
		},
		{
			user: User{
				ID:        uuid,
				AccountID: "456",
				Username:  "johndoe",
				FirstName: "John",
				LastName:  "Doe",
				Email:     nil,
				Phone:     &phone,
			},
			expected: map[string]interface{}{
				"id":         uuid,
				"account_id": "456",
				"username":   "johndoe",
				"first_name": "John",
				"last_name":  "Doe",
				"phone":      "123-456-7890",
			},
		},
	}

	for _, test := range tests {
		output := test.user.toMapInterface()
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("toMapInterface() = %v; expected %v", output, test.expected)
		}
	}
}
