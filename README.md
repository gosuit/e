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

### Error usage

```golang
package main

import (
	"errors"

	"github.com/gosuit/e"
)

func main() {
	err := e.New("error message", e.Internal).
		WithTag("key", "value").
		WithErr(errors.New("underlying error"))

	err.GetMessage()           // error message
	err.GetStatus().ToString() // Internal
	err.GetError()             // underlying error
	err.GetTag("key")          // value
	err.GetSource()            // <your path>/main.go 12
	err.GetHttpCode()          // 500
	err.GetGrpcCode()          // Internal
	err.ToJson()               // {error message}
	err.Error()                // error message: underlying error
	err.ToGRPC()               // rpc error: code = Internal desc = error message
	err.SlErr()                // error=error message: underlying error
}
```

### Error converting

```golang
package main

import (
	"errors"

	"github.com/gosuit/e"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	err := e.E(errors.New("error"))

	// Use err....

	err = e.FromGRPC(status.Error(codes.Internal, "msg"))

	// Use err...
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
