package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing Celsius = 0
	BoilingC Celsius = 100
)
