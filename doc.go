// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package go-exactonline provides a client for using the Exact Online API.

Usage:

	import "github.com/mcnijman/go-exactonline"

Note that this library doens't directly handle authentication, see [Authentication](#authentication).
We first contstruct a client and then access the various API endpoints.

	client := exactonline.NewClient(nil)

	// Get the last used division
	divisionID, err := client.GetCurrentDivisionID()

	// Fetch all transactions in the division
	transactions, err := client.FinancialTransaction.Transactions.List(context.Background(), divisionID, false)
*/
package exactonline
