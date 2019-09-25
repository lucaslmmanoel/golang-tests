package main

import "testing"

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Lucas", "")
        want := "Hello, Lucas"
        assertCorrectMessage(t, got, want)
    })


    t.Run("saying hello to french people", func(t *testing.T) {
        got := Hello("Lucas", "French")
        want := "Bonjour, Lucas"
        assertCorrectMessage(t, got, want)
    })


    t.Run("saying hello to spanish people", func(t *testing.T) {
        got := Hello("Lucas", "Spanish")
        want := "Hola, Lucas"
        assertCorrectMessage(t, got, want)
    })



}