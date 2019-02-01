package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertEquals(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(0)
		assertEquals(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw inssuficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		w := Wallet{startBalance}
		err := w.Withdraw(100)
		assertEquals(t, w, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)

	})
}

func assertEquals(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("Got %s; Want: %s", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Got error but didn't expect one")
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Error expected, but didn't ge one.")
	}

	if err != want {
		t.Errorf("got: '%s'; want '%s'", err, want)
	}
}
