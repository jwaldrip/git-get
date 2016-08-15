# Tint [![Coverage Status](https://img.shields.io/coveralls/jwaldrip/tint.svg)](https://coveralls.io/r/jwaldrip/tint?branch=master) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/jwaldrip/tint)

## Usage

```go
package main

import (
	"fmt"

	"github.com/jwaldrip/tint"
)

func main() {
	msg := tint.Colorize("color me blue", tint.Blue)
	fmt.Println(msg)
}
```




