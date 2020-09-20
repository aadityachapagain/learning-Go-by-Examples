// package temp conv do the celcius and farhienheit computations

package tempconv

import (
   "fmt"
)
type Celsius float64
type fahrenheit float64

const (
   AbsoluteZeroC Celsius = -273.15
   FreezingC Celsius = 0
   BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(F Fahrenheit) Celsius { return Celsius( (f - 32) * 5 / 9 ) }
