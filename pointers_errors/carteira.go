package pointerserrors

import (
	"errors"
	"fmt"
)

// Isto pode ser muito útil quando queremos adicionar funcionalidades de domínios específicos a tipos já existentes.

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

//quando uma função ou um método é invocado, os argumentos são copiados.
/*
Quando func (c Carteira) Depositar(quantidade int) é chamado,
o c é uma cópia do valor de qualquer lugar que o método tenha sido chamado.
*/
func (c *Carteira) Depositar(quantidade Bitcoin) {
	//fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

// func (c *Carteira) Retirar(quantidade Bitcoin) error {
// 	c.saldo -= quantidade
// 	return nil
// }

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= quantidade
	return nil
}
