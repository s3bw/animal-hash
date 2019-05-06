<h1 align="center">
    Animal Hash
</h1>

<h4 align="center">
    A hash that returns a unique string of animal えもじ
</h4>

## Usage

```golang
package main

import (
	"fmt"

	animalhash "github.com/foxyblue/animal-hash/animal-hash"
)

func main() {
    // Output: 🐭🐱🐹
	fmt.Println(animalhash.Hash("dog", 23))
}
```
