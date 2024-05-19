package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(42))
		assertBalance(t, wallet, Bitcoin(42))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(33)}
		err := wallet.Withdraw(Bitcoin(2))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(31))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(9000)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(9999))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("expected no error but got error")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error but no error")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestBitcoin(t *testing.T) {
	b := Bitcoin(64)
	got := b.String()
	want := "64 BTC"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
