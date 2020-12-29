package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 声明是辅助函数, 报告行数将在调用方
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("zf")
		want := "hello, zf"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world, when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})

}
