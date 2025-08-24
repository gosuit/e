# E

This GoLang library allows you to handle errors more flexibly.

A custom error allows you to carry a large amount of additional error data and convert it to various formats.

## Installation

```zsh
go get github.com/gosuit/e
```

## Features

- **Additional information**: message, status, tags, underlying error and error source.
- **Codes managing**: getting HTTP and gRPC code based on status.
- **Сonversion**: converting errors to and from various formats.

## Usage

```golang
package main

import (
	"fmt"

	"github.com/gosuit/e"
)

func main() {
	err := e.New("error", e.Internal)

	fmt.Println(err.GetHttpCode())
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
