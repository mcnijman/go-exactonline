# go-exactonline #

[![GoDoc](https://godoc.org/github.com/mcnijman/go-exactonline?status.svg)](https://godoc.org/github.com/mcnijman/go-exactonline) [![Build Status](https://travis-ci.org/mcnijman/go-exactonline.svg?branch=master)](https://travis-ci.org/mcnijman/go-exactonline) [![Test Coverage](https://coveralls.io/repos/github/mcnijman/go-exactonline/badge.svg?branch=master)](https://coveralls.io/github/mcnijman/go-exactonline?branch=master)

go-exactonline is a Go client library for accessing the Exact Online API.

This is library is incomplete and under development.

## Usage ##

```go
import "github.com/mcnijman/go-exactonline"
```

## Authentication ##

This library doesn't directly handle authentication. You should provide a `http.Client` that handles the authentication for you.

## Divisions ##

## Versioning ##

This library uses symantic versions using git tags. However since this library is still in development, API methods are subject to change.

## TODO ##

- Implement all CRUD operations
- Helpers for odata types creation
- Integration tests
- Error handling
- Web hooks support
- Documentation and examples

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.