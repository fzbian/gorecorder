package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid input",
			input:    "123",
			expected: 123,
			wantErr:  false,
		},
		{
			name:     "Invalid input",
			input:    "abc",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringToInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("stringToInt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetAvaliableScreens(t *testing.T) {
	var buf bytes.Buffer

	getAvaliableScreens()

	got := buf.String()
	if !strings.Contains(got, "Active displays:") {
		t.Errorf("getAvaliableScreens() = %q, want %q", got, "Active displays:")
	}
	if !strings.Contains(got, "Id:") {
		t.Errorf("getAvaliableScreens() = %q, want %q", got, "Id:")
	}
	if !strings.Contains(got, "Bounds") {
		t.Errorf("getAvaliableScreens() = %q, want %q", got, "Bounds")
	}
}

func TestCaptureScreenshot(t *testing.T) {
	tests := []struct {
		name        string
		screen      int
		output      string
		expectedErr bool
	}{
		{
			name:        "Valid input",
			screen:      0,
			output:      "test",
			expectedErr: false,
		},
		{
			name:        "Invalid screen",
			screen:      -1,
			output:      "test",
			expectedErr: true,
		},
		{
			name:        "Invalid output",
			screen:      0,
			output:      "",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := captureScreenshot(tt.screen, tt.output)
			if (err != nil) != tt.expectedErr {
				t.Errorf("captureScreenshot() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
