package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't Sqrt negative number: %g\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z, d := x/2, 0.
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-d) < 1e-9 {
			break
		}
		d = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
