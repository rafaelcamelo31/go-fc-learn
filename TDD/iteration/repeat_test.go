package iteration

import (
	"fmt"
	"testing"
)

func TestRepe(t *testing.T) {
	character := "a"
	count := 5

	repeated := Repeat(character, count)
	var expected string

	for i := 0; i < count; i++ {
		expected += character
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
