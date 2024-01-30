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

	confirmaError := func(t *testing.T, resultado error, esperado error) {
		t.Helper()
		if resultado == nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu.")
		}

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

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(10)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaError(t, erro, ErroSaldoInsuficiente)
	})
}
