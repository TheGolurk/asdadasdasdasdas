package str

import (
	"errors"
	"fmt"
	"strconv"
)

// Karina
// -> Hello Karina
func Saludo(name string) (string, error) {
	if name == "" {
		return "", errors.New("cannot send empty name")
	}

	return fmt.Sprintf("Hello %v", name), nil
}

// " 10 "
// -> 10, nil
// ------------------
// "#"
// 0, error
func AtoiT(num string) (int, error) {
	return strconv.Atoi(num)
}
