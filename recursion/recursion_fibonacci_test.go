package recursion

import (
	"fmt"
	"testing"
)

func TestRecursionFibonacci(t *testing.T) {

	// TODO how to gc the wrapFibonacci closure as times go
	fibonacci := wrapFibonacci()
	result := fibonacci(3)

	fibonacci(15)
	fmt.Printf("%d", result)
	if result != 1 {
		t.Errorf("fibonacci on 3; want '%v', got '%v';", 1, result)
	}
	result = fibonacci(30)

	if result != 514229 {
		t.Errorf("fibonacci on 30; want '%v', got '%v';", 30, result)
	}

}
