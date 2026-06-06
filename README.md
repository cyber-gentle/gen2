# gen2

A lightweight collection of reusable Go utility packages designed to simplify common programming tasks. Built with simplicity, readability, and reusability in mind.

## Features

### `solve`

Arithmetic utility functions, including:

* Sum
* Difference
* Additional mathematical helpers (planned)

### `stringutil`

String manipulation helpers, including:

* Uppercase conversion
* Lowercase conversion
* Word capitalization
* Additional string utilities (planned)

### `toolkit`

General-purpose utility functions, including:

* Slice helpers
* Collection utilities
* Common reusable helpers
* Additional utilities (planned)

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
	fmt.Println(solve.Difference(10, 2, 3))
}
```

### Toolkit Utilities

```go
package main

import (
	"fmt"

	"github.com/cyber-gentle/gen2/toolkit"
)

func main() {
	fmt.Println(toolkit.MergeSlices(
		[]int{1, 2},
		[]int{3, 4},
	))
}
```

## Project Structure

```text
gen2/
├── solve/
│   └── arithmetics.go
├── stringutil/
│   └── strings.go
├── toolkit/
│   └── toolkit.go
├── go.mod
└── README.md
```

## Design Goals

* Simple and intuitive APIs
* Idiomatic Go implementation
* Lightweight and dependency-free
* Reusable across projects
* Educational and beginner-friendly
* Continuously expanding utility collection

## Versioning

This project follows Semantic Versioning (SemVer).

* Patch releases (`v1.2.x`) contain fixes and improvements.
* Minor releases (`v1.x.0`) introduce new functionality without breaking compatibility.
* Major releases (`v2.0.0`) may contain breaking changes.

## Contributing

Contributions, suggestions, and bug reports are welcome.

## License

MIT License
