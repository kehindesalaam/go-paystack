# go-paystack #

go-paystack is a Go client library for accessing the Paystack.

**Build Status:** ![Build Status](https://travis-ci.com/kehindesalaam/go-paystack.svg?token=jEi76ESgT7V1Uzbsyqb8&branch=master)

## Usage ##
```go
import "github.com/kehindesalaam/go-paystack/paystack"
```

Construct a new Paystack client, then use the various services on the client to access different parts of the Paystack API. For example:

```go
client := paystack.NewClient(nil, paystack.SecretKey("sk_test_your_secret_key"))
newCustomer, _, err := client.Customer.Fetch(ctx, "12345");
```

The services of a client divide the API into logical chunks and correspond to the structure of the Paystack API documentation at https://developers.paystack.co/reference

### Creating and Updating Resources ###

All structs for Paystack resources use pointer values for all non-repeated fields.
This allows distinguishing between unset fields and those set to a zero-value.
All requests should be made with the Client request objects as in `paystack.CustomerRequest`
Helper functions have been provided to easily create these pointers for string,bool, and int values. For example:

```go
// create a new customer with email "foo@testing.com"
cust := &paystack.CustomerRequest{
    Email: paystack.String("foo@testing.com"),
    FirstName: paystack.String("Kehinde"),
    LastName: paystack.String("Salaam"),
    Phone: paystack.String("123456789")
 }
client.Customer.Create(ctx, "", cust)
```

Users who have worked with protocol buffers should find this pattern familiar.

### Pagination ###

All requests for resource collections (repos, pull requests, issues, etc.) support pagination. Pagination options are described in the
`paystack.ListOptions` struct and passed to the list methods directly or as an embedded type of a more specific list options struct (for example `paystack.TransactionOptions`). Pages information is available via the `paystack.Response` struct.

```go
client := paystack.NewClient(nil, paystack.SecretKey("sk_test_your_secret_key"))

opt := &paystack.TransactionOptions{
	ListOptions: paystack.ListOptions{PerPage: 10},
}
// get all pages of results
var allTxns []*paystack.Transaction
for {
	txns, resp, err := client.Transaction.List(ctx, opt)
	if err != nil {
		return err
	}
	allTxns = append(allTxns, txns...)
	if resp.NextPage == 0 {
		break
	}
	opt.Page = resp.NextPage
}
```

## LICENSE ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.