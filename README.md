# plrl
Simple Go functions for handling language plurals.

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/plrl.svg)](https://pkg.go.dev/github.com/speedyhoon/plrl)
[![go report card](https://goreportcard.com/badge/github.com/speedyhoon/plrl)](https://goreportcard.com/report/github.com/speedyhoon/plrl)

## Example
```go
package main

import (
	"fmt"
	"github.com/speedyhoon/plrl"
)

func main() {
	want := 4
	fmt.Printf("At least %d item%s required.\n", want, plrl.Any(want, " is", "s are"))
	want = 1
	fmt.Printf("At least %d item%s required.\n", want, plrl.Any(want, " is", "s are"))
	days := 5
	fmt.Printf("%d day%s\n", days, plrl.S(days))
	days = -1
	fmt.Printf("%d day%s\n", days, plrl.S(days))
}
```
Output
```
At least 4 items are required.
At least 1 item is required.
5 days
-1 day
```
