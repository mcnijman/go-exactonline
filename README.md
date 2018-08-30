# go-exactonline #

[![GoDoc](https://godoc.org/github.com/mcnijman/go-exactonline/github?status.svg)](https://godoc.org/github.com/mcnijman/go-exactonline/github) [![Build Status](https://travis-ci.org/mcnijman/go-exactonline.svg?branch=master)](https://travis-ci.org/mcnijman/go-exactonline) [![Test Coverage](https://coveralls.io/repos/github/mcnijman/go-exactonline/badge.svg?branch=master)](https://coveralls.io/github/mcnijman/go-exactonline?branch=master)

go-exactonline is a Go client library for accessing the Exact Online API. This is a limited implemenation, that doesn't yet support all endpoints and operations.

## Usage ##

```go
import "github.com/mcnijman/go-exactonline/exactonline"
```

## Authentication ##

This library doesn't directly handle authentication. You should provide a `http.Client` that handles the authentication for you.

## Divisions ##

The last used division by the user is available in the `Me.GetLastUsedDivision()` service. All user accessible divisions are accessible through the `Divisions` service.

## Implementation status ##


## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.