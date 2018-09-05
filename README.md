# go-exactonline #

[![GoDoc](https://godoc.org/github.com/mcnijman/go-exactonline?status.svg)](https://godoc.org/github.com/mcnijman/go-exactonline) [![Build Status](https://travis-ci.org/mcnijman/go-exactonline.svg?branch=master)](https://travis-ci.org/mcnijman/go-exactonline) [![Test Coverage](https://coveralls.io/repos/github/mcnijman/go-exactonline/badge.svg?branch=master)](https://coveralls.io/github/mcnijman/go-exactonline?branch=master) [![Maintainability](https://api.codeclimate.com/v1/badges/a2ca34f94cb3bc58e6a1/maintainability)](https://codeclimate.com/github/mcnijman/go-exactonline/maintainability) [![go report](https://goreportcard.com/badge/github.com/mcnijman/go-exactonline)](https://goreportcard.com/report/github.com/mcnijman/go-exactonline)

go-exactonline is a Go client library for accessing the Exact Online API. This library is tested for Go 1.10 and above.

> This is library is incomplete and under development. API's are subject to change, but they will be stable in the near future.

## Usage ##

```go
import "github.com/mcnijman/go-exactonline"
```

Note that this library doens't directly handle authentication, see [Authentication](#authentication).
We first contstruct a client and then access the various API endpoints.

```go
client := exactonline.NewClient(nil)
ctx := context.Background()

// Get the last used division
divisionID, err := client.GetCurrentDivisionID(ctx)

// Fetch all transactions in the division
transactions, err := client.FinancialTransaction.Transactions.List(ctx, divisionID, false, nil)
```

## Authentication ##

This library doesn't directly handle authentication. You should provide a `http.Client` that handles the authentication for you.
There are multiple ways to do this, however these are the recommended ways:

```go
tokenSource := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
)
client := exactonline.NewClientFromTokenSource(context.Background(), tokenSource)
```

Or use your oauth2 configuration and the `oauth2` package will automatically refresh the token for you.

```go
token := &oauth2.Token{} // Your previously stored or fetched token

ctx := context.Background()
config := &oauth2.Config{
    RedirectURL:  "the registered redirect URL",
    ClientID:     "the registered client ID",
    ClientSecret: "the registered client secret",
    Endpoint: oauth2.Endpoint{
        AuthURL:  "https://start.exactonline.nl/api/oauth2/auth",
        TokenURL: "https://start.exactonline.nl/api/oauth2/token",
    },
}

tokenSource := config.TokenSource(ctx, token) // this will refresh your access token if a valid refresh token is available
httpClient := oauth2.NewClient(ctx, tokenSource) // Create a http.Client that you want to tweak or use exactonline.NewClientFromTokenSource
client := exactonline.NewClient(nil)
```

## Divisions ##

The current division can be fetched using:

```go
divisionID, err := client.GetCurrentDivisionID(context.Background())
```

Other available divisions can be fecthed through te following enpoints:

```go
// To get all divisions which are accessible for the user that granted the app permission, use:
divisions, err := client.System.Divisions.List(context.Background(), true, nil)
// or if you need to retrieve the divisions for the current license, of the user that granted
// the app permission, use:
divisions, err := client.HRM.Divisions.List(context.Background(), true, nil)
```

## Pagination ##

TODO

## Issues ##

Issues and/or pull requests are welcome. Note that the services are generated using `gen-services.go`. If there are issues with the services, take a look at this file or the templates. The service files shouldn't be edites.

## Versioning ##

This library uses symantic versions using git tags. However since this library is still in development, API methods are subject to change.

## TODO ##

- Implement all CRUD operations (currently only `List` is supported)
- Add support for non standard endpoints
- Integration tests
- Error handling
- Web hooks support
- Documentation and examples

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.