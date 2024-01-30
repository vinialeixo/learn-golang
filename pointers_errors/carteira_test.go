package pointerserrors

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	// t.Run("Depositar", func(t *testing.T) {
	// 	carteira := Carteira{}

	// 	carteira.Depositar(Bitcoin(10))

	// 	resultado := carteira.Saldo()

	// 	//fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

	// 	esperado := Bitcoin(10)
	// 	if resultado != esperado {
	// 		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	// 	}

	// })

	// t.Run("Retirar", func(t *testing.T) {
	// 	carteira := Carteira{saldo: Bitcoin(20)}

	// 	carteira.Retirar(Bitcoin(10))

	// 	resultado := carteira.Saldo()

	// 	//fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

	// 	esperado := Bitcoin(10)
	// 	if resultado != esperado {
	// 		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	// 	}
	// })

	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		cartiera := Carteira{saldo: Bitcoin(20)}
		cartiera.Retirar(10)
		confirmaSaldo(t, cartiera, Bitcoin(10))
	})
}
