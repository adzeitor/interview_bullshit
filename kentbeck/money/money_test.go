package money

import (
	"testing"
)

func TestDollar_Times(t *testing.T) {
	five := Dollar(5)

	assertMoney(t, Dollar(10), five.Times(2))
	assertMoney(t, Dollar(15), five.Times(3))
}

func TestDollar_Equal(t *testing.T) {
	assertTrue(t, Dollar(5).Equal(Dollar(5)))
	assertFalse(t, Dollar(5).Equal(Dollar(6)))
	assertTrue(t, Franc(5).Equal(Franc(5)))
	assertFalse(t, Franc(5).Equal(Franc(6)))

	dollarFive := Dollar(5)
	francFive := Franc(5)
	assertFalse(t, francFive.Equal(dollarFive))

	moneyFranc := Money{
		amount:   5,
		currency: "CHF",
	}
	assertTrue(t, moneyFranc.Equal(Franc(5)))
}

func TestMoney_Currency(t *testing.T) {
	assert(t, "USD", Dollar(1).currency)
	assert(t, "CHF", Franc(1).currency)
}
