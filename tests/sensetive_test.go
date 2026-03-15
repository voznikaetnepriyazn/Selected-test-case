package loglinter_test

import (
	"testing"

	"testcase/rules"
)

func TestContainsSensitiveData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"plain text", "server started", false},
		{"password mention", "password field required", false},
		{"api_key mention", "api_key is required", false},

		{"password value", "password: secret123", true},
		{"api_key value", "api_key=abc123", true},
		{"token value", "token: bearer123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found := rules.IsSensetiveData(tt.input)
			if found != tt.expected {
				t.Errorf("ContainsSensitiveData(%q) = %v, want %v",
					tt.input, found, tt.expected)
			}
		})
	}
}

