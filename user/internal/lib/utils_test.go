package lib

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"FirstName", "first_name"},
		{"LastName", "last_name"},
		{"Email", "email"},
		{"Phone", "phone"},
		{"UserName", "user_name"},
	}

	for _, test := range tests {
		output := ToSnakeCase(test.input)
		if output != test.expected {
			t.Errorf("toSnakeCase(%s) = %s; expected %s", test.input, output, test.expected)
		}
	}
}
