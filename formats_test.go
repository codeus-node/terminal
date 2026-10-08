package terminal_test

import (
	"testing"

	"github.com/codeus-node/terminal"
)

func TestConstantsFormatsExist(t *testing.T) {
	testCases := []struct {
		name     string
		constant string
	}{
		{"Bold", terminal.Bold},
		{"Underline", terminal.Underline},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.constant == "" {
				t.Errorf("Constant %s should not be empty", tc.name)
			}
		})
	}
}
