package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("check output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})

	t.Run("check sleep write order", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("incorrect order of calls to sleeper, want %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
