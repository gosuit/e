# E

This GoLang library provides a custom error interface that enhances the standard error handling capabilities in Go. It encapsulates error messages, status codes, and conversion methods for gRPC and HTTP responses, allowing developers to manage errors more effectively in their applications.

## Features

- Custom error type with additional context and functionality.
- Methods to retrieve error messages, status, underlying errors, and associated tags.
- Support for logging errors with metadata.
- Conversion methods for gRPC and HTTP status codes.
- Structured logging attributes for better log analysis.

## Installation

To install the library, use the following command:
```zsh
go get github.com/gosuit/e
```

## Usage

```golang
package main

import (
    "github.com/gosuit/e"
)

func main() {
    value, err := some()
    if err != nil {
        fmt.Println(err.ToHttpCode()) // Output: 500
    }

    // Do something
}

func some() (int, e.Error) {
    // Do something

    return 0, e.New("Someting going wrong...", e.Internal)
}
```

## Methods

The Error interface includes the following methods:

- GetMessage() string
- GetError() error
- GetTag(key string) any
- GetCode() Status
- GetSource() (string, int)
- Log(msg ...string)
- WithMessage(string) Error
- WithErr(error) Error
- WithTag(key string, value any) Error
- WithCtx(c lec.Context) Error
- WithCode(Status) Error
- ToJson() jsonError
- ToGRPCCode() codes.Code
- ToHttpCode() int
- Error() string
- SlErr() slog.Attr
- ToGRPCErr() error

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
