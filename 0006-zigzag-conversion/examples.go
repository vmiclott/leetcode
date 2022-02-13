package zigzagconversion

import (
	"fmt"
)

func example(s string, numRows int) {
	fmt.Printf("Input: s = \"%v\", numRows = %v\n", s, numRows)
	fmt.Printf("Output: %v\n", convert(s, numRows))
}

func Examples() {
	fmt.Println("6. Zigzag Conversion")
	example("PAYPALISHIRING", 3)
	example("PAYPALISHIRING", 4)
}
