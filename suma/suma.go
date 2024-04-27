package suma

import "errors"

func Sumar(a, b int) int {
	return a + b
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}
