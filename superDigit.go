package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(superDigit("9875", 4))
}

func superDigit(s string, k int32) int32 {
	// Write your code here

	var suma int32
	const asciizero = 48

	if len(s) == 1 {
		return int32(s[0]) - asciizero
	} else {
		fmt.Println("s: ", s)

		for i := 0; i < len(s); i++ {
			suma += (int32(s[i]) - asciizero) * k
		}

		sdRecursivo := strconv.Itoa(suma)

		return superDigit(sdRecursivo, 1)
	}
}
