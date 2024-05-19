package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := &Wallet{}
	wallet.Deposit(Bitcoin(42))

	got := wallet.Balance()
	want := Bitcoin(42)

	if got != want {
		t.Errorf("got %d want %d", got, want)
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
