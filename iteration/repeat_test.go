package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	excepted := "aaaaa"

	if repeated != excepted {
		t.Errorf("expected %q, bug got %q", excepted, repeated)
	}
}


func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}