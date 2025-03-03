package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %v, but want %v", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(8)}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(18))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(8))
		assertBalance(t, wallet, Bitcoin(12) )
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(12)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(18))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}