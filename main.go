package main

import (
	"fmt"

	animalhash "github.com/foxyblue/animal-hash/animal-hash"
)

func main() {
	// Output: 🐭🐱🐹
	fmt.Println(animalhash.Hash("dog", 23))
}
