# A quine, in Go

Well, here it is:

```go
package main

import "fmt"

const p = `package main

import "fmt"

const p = %c%s%c

func main() {
	fmt.Printf(p, 96, p, 96)
}
`

func main() {
	fmt.Printf(p, 96, p, 96)
}

```
