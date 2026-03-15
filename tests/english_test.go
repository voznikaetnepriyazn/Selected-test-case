package loglinter_test

import (
	"testing"

	"testcase/rules"
)

func TestIsEnglishOnly(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"english text", "server started", true},
		{"with numbers", "port 8080", true},
		{"url", "https://example.com", true},
		{"empty string", "", true},

		{"cyrillic", "сервер запущен", false},
		{"emoji", "started 🚀", false},
		{"mixed", "starting запуск", false},
		{"unicode arrow", "redirect → /home", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsEnglishOnly(tt.input)
			if result != tt.expected {
				t.Errorf("IsEnglishOnly(%q) = %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}
