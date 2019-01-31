package tempconv

// CToF Converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC Converts Fahrenheit
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
