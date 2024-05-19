package banking

import "testing"

func TestWallet(t *testing.T) {
	wallet := &Wallet{}
	wallet.Deposit(42)

	got := wallet.Balance()
	want := 42

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
