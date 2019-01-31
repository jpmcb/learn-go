package tempconv

import "testing"

func TestCToF(t *testing.T) {
	cases := []struct {
		in   Celsius
		want Fahrenheit
	}{
		{0, 32},
		{100, 212},
	}
	for _, c := range cases {
		got := CToF(c.in)
		if got != c.want {
			t.Errorf("CToF(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFToC(t *testing.T) {
	cases := []struct {
		in   Fahrenheit
		want Celsius
	}{
		{32, 0},
		{212, 100},
	}
	for _, c := range cases {
		got := FToC(c.in)
		if got != c.want {
			t.Errorf("FToC(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
