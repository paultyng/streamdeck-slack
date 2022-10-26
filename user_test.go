package main

import "testing"

func TestEmojiColonSyntaxToName(t *testing.T) {
	for _, tc := range []struct {
		colonSyntax string
		expected    string
	}{
		{":smile:", "smile"},
		{":smile::skin-tone-2:", "smile"},
		{"", ""},
		{"foo", ""},
	} {
		t.Run(tc.colonSyntax, func(t *testing.T) {
			actual := emojiColonSyntaxToName(tc.colonSyntax)
			if tc.expected != actual {
				t.Fatalf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
