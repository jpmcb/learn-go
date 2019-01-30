package tempconv

import "testing"

func TestCToF(t *testing.T) {
	cases := []struct {
		in Celsius
        want Fahrenheit
	}{
		{0, 32},
	}
	for _, c := range cases {
		got := CToF(c.in)
		if got != c.want {
			t.Errorf("CToF(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
