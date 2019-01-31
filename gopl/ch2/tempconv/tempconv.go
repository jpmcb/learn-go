package tempconv

import (
	"fmt"
)

// Celsius Type for C temp value
type Celsius float64

// Fahrenheit Type for F temp value
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // Absolute zero point
	Freezing      Celsius = 0       // Freezing point
	BoilingC      Celsius = 100     // Boiling point
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºF", f)
}
