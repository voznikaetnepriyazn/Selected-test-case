package loglinter_test

import (
	"testing"

	"testcase/rules"
)

func TestStartsWithLowercase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"lowercase start", "starting server", true},
		{"with leading spaces", "  starting server", true},
		{"starts with number", "123 server started", true},
		{"starts with special char", "/api/v1/users", true},
		{"empty string", "", true},
		{"only special chars", "!@#$%", true},

		{"uppercase start", "Starting server", false},
		{"uppercase after spaces", "  Starting server", false},
		{"all uppercase", "SERVER STARTED", false},
		{"uppercase after prefix", "[INFO] Starting", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsStartsFromLowerCase(tt.input)
			if result != tt.expected {
				t.Errorf("StartsWithLowercase(%q) = %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

