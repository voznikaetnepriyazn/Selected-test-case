package loglinter_test

import (
	"testing"

	"testcase/rules"
)

func TestHasEmojiOrSpecialSymbol(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"plain text", "server started", false},
		{"single exclamation", "success!", false},
		{"double exclamation", "success!!", false},
		{"single dot", "done.", false},
		{"double dot", "waiting..", false},
		{"url", "https://example.com", false},

		{"triple exclamation", "success!!!", true},
		{"triple question", "what???", true},
		{"triple dot", "waiting...", true},
		{"emoji rocket", "launched 🚀", true},
		{"emoji fire", "hot 🔥", true},
		{"emoji check", "done ✅", true},
		{"mixed violations", "wow!!! 🎉", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rules.IsEmojiOrSpecialSymbol(tt.input)
			if result != tt.expected {
				t.Errorf("HasEmojiOrSpecialSymbol(%q) = %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

