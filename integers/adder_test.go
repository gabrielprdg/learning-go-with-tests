package integers

import (
	"fmt"
	"testing"
)

// %q -> string
// %d -> integer

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4
}

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
