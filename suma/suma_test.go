package suma

import "testing"

func TestSumar(t *testing.T) {
	// ejecuta en parelelo
	t.Parallel()

	t.Run("Suma Digitos ok", func(t *testing.T) {
		resultado := Sumar(1, 2)
		if resultado != 3 {
			t.Errorf("Sumar(1,2), expected: 3, got: %v", resultado)
		}
	})

	t.Run("Suma digitos negativos ok", func(t *testing.T) {
		resultado := Sumar(10, -5)
		if resultado != 5 {
			t.Errorf("Sumar(10,-5), expected: 5, got: %v", resultado)
		}
	})

}

func TestDividir(t *testing.T) {
	t.Parallel()

	t.Run("Division por 0 error", func(t *testing.T) {
		resultado, err := Dividir(5, 0)
		if err == nil && resultado != 0 {
			t.Errorf("Dividir(5,0), expected: error, got %v. Result: %v", err, resultado)
		}
	})

	t.Run("Division numeros ok", func(t *testing.T) {
		resultado, err := Dividir(10, 2)
		if resultado != 5 && err != nil {
			t.Errorf("Dividir(10,2), expected: 5, got %v, Error %v", resultado, err)
		}
	})

}
