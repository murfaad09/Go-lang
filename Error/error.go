package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for i := 1; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
		if (z*z-x)/(2*z) < 0.00000001 {
			return z, nil
		}
	}
	return z, nil
}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
