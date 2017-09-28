package paystack

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBalanceService_Check(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
				  "status": true,
				  "message": "Balances retrieved",
				  "data": [
					{
					  "currency": "NGN",
					  "balance": 1700000
					}
				  ]
				}`)
	})

	balance, _, err := client.Balance.Check(context.Background())
	if err != nil {
		t.Errorf("Balance.Check returned error: %v", err)
	}
	want := []*Balance{{Currency: String("NGN"), Balance: Int(1700000)}}
	if !reflect.DeepEqual(balance, want) {
		t.Errorf("Balance.Check returned %+v, want %+v", balance, want)
	}
}
