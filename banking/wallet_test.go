package banking

import "testing"

func TestWallet(t *testing.T) {
	wallet := &Wallet{}
	wallet.Deposit(Bitcoin(42))

	got := wallet.Balance()
	want := Bitcoin(42)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
