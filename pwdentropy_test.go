package pwdentropy

import "testing"

func TestCalculate(t *testing.T) {
	tests := map[string]struct {
		input string
		want  float64
	}{
		"4 characters: letters of the same case": {
			input: "abcd",
			want:  18.80,
		},
		"8 characters: letters of the same case": {
			input: "abcdefgh",
			want:  37.60,
		},
		"8 characters: letters of upper and lower case": {
			input: "abcdEFGH",
			want:  45.60,
		},
		"8 characters: letters of upper and lower case, number and symbol": {
			input: "Abcd3FG!",
			want:  52.56,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Calculate(tc.input)
			if got != tc.want {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}
