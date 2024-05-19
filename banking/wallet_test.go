package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(42))
		assertBalance(t, wallet, Bitcoin(42))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(33)}
		wallet.Withdraw(Bitcoin(2))
		assertBalance(t, wallet, Bitcoin(31))
	})

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("expected error but no error")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(9000)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(9999))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, "not enough money bud")
	})
}

func TestBitcoin(t *testing.T) {
	b := Bitcoin(64)
	got := b.String()
	want := "64 BTC"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
