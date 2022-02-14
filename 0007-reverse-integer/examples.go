package reverseinteger

import (
	"fmt"
)

func example(x int) {
	fmt.Printf("Input: x = %v\n", x)
	fmt.Printf("Output: %v\n", reverse(x))
}

func Examples() {
	fmt.Println("7. Reverse Integer")
	example(123)
	example(-123)
	example(120)
}
