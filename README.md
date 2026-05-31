# gen2

A lightweight collection of reusable Go utility packages designed to simplify common programming tasks.

## Features

### `solve`

Arithmetic utility functions for mathematical operations.

### `stringutil`

String manipulation helpers, including:

* Uppercase conversion
* Lowercase conversion
* Word capitalization
* Additional string utilities (planned)

## Installation

```bash
go get github.com/cyber-gentle/gen2
```

## Usage

### String Utilities

```go
package main

import (
	"fmt"

	"github.com/cyber-gentle/gen2/stringutil"
)

func main() {
	fmt.Println(stringutil.UpperCase("hello"))
	fmt.Println(stringutil.LowerCase("HELLO"))
	fmt.Println(stringutil.Capitalize("hello world"))
}
```

### Arithmetic Utilities

```go
package main

import (
	"fmt"

	"github.com/cyber-gentle/gen2/solve"
)

func main() {
	fmt.Println(solve.Sum(2, 3))
}
```

## Project Structure

```text
gen2/
├── solve/
│   └── arithmetics.go
├── stringutil/
│   └── strings.go
└── README.md
```

## Goals

* Simple and reusable APIs
* Idiomatic Go implementation
* Lightweight dependencies
* Educational and practical utility functions

## License

MIT License
