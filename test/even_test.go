package even

import (
	"testing"
)

func Test_Loop(t *testing.T) {
	res := Loop(uint64(3))
	if res != 6 {
		t.Error("Loop:", Loop(uint64(3)))
	}
}

func TestFactorial(t *testing.T) {
	res := Factorial(uint64(3))
	if res != 6 {
		t.Error("Factorial:", Factorial(uint64(3)))
	}
}

func BenchmarkLoop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Factorial(uint64(40))
	}
}
