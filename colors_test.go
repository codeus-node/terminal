package terminal_test

import (
	"testing"

	"github.com/codeus-node/terminal"
)

func TestConstantsColorsExist(t *testing.T) {
	testCases := []struct {
		name     string
		constant string
	}{
		{"EndFormat", terminal.EndFormat},
		{"Red", terminal.Red},
		{"Green", terminal.Green},
		{"Blue", terminal.Blue},
		{"Purple", terminal.Purple},
		{"Cyan", terminal.Cyan},
		{"White", terminal.White},
		{"Black", terminal.Black},
		{"Yellow", terminal.Yellow},
		{"LightYellow", terminal.LightYellow},

		{"BackBlack", terminal.BackBlack},
		{"BackWhite", terminal.BackWhite},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.constant == "" {
				t.Errorf("Constant %s should not be empty", tc.name)
			}
		})
	}
}
