package paystack

import (
	"context"
	"fmt"
	"net/http"
	//"reflect"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
	"time"
)

func TestCustomerService_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Customer created",
		  "data": {
			"email": "bojack@horsinaround.com",
			"integration": 100032,
			"domain": "test",
			"customer_code": "CUS_xnxdt6s1zg1f4nx",
			"id": 1173,
			"createdAt": "2016-03-29T20:03:09.584Z",
			"updatedAt": "2016-03-29T20:03:09.584Z"
		  }
		}`)
	})

	customerRequest := CustomerRequest{Email: String("horsinaround"), FirstName: String("horsinaround"), LastName: String("horsinaround"), Phone: String("horsinaround")}

	customer, _, err := client.Customer.Create(context.Background(), &customerRequest)
	if err != nil {
		t.Errorf("Customer.Create returned error: %v", err)
	}
	createdAt := time.Date(2016, 3, 29, 20, 3, 9, 584000000, time.UTC)
	updatedAt := time.Date(2016, 3, 29, 20, 3, 9, 584000000, time.UTC)

	want := &Customer{Email: String("bojack@horsinaround.com"), Integration: Int(100032), Domain: String("test"), CustomerCode: String("CUS_xnxdt6s1zg1f4nx"), Id: Int(1173), CreatedAt: &createdAt, UpdatedAt: &updatedAt}
	if !cmp.Equal(customer, want) {
		t.Errorf("Customer.Create returned %+v, want %+v", customer, want)

	}
}

func TestCustomerService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Customers retrieved",
		  "data": [
			{
			  "integration": 100032,
			  "first_name": "Diane",
			  "last_name": "Nguyen",
			  "email": "diane@writersclub.com",
			  "phone": "16504173147",
			  "metadata": null,
			  "domain": "test",
			  "customer_code": "CUS_1uld4hluw0g2gn0",
			  "id": 63,
			  "createdAt": "2016-03-29T20:03:09.0Z",
			  "updatedAt": "2016-03-29T20:03:09.0Z"
			},
			{
			  "integration": 100032,
			  "first_name": null,
			  "last_name": null,
			  "email": "todd@chavez.com",
			  "phone": null,
			  "metadata": null,
			  "domain": "test",
			  "customer_code": "CUS_soirsjdqkyjfwcr",
			  "id": 65,
			  "createdAt": "2016-03-29T20:03:09.0Z",
			  "updatedAt": "2016-03-29T20:03:09.0Z"
			}
		  ],
		  "meta": {
			"total": 2,
			"skipped": 0,
			"perPage": 50,
			"page": 1,
			"pageCount": 1
		  }
		}`)
	})

	customer, _, err := client.Customer.List(context.Background(), nil)
	if err != nil {
		t.Errorf("Customer.List returned error: %v", err)
	}

	createdAt1 := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)
	updatedAt1 := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)
	createdAt2 := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)
	updatedAt2 := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)

	want := []*Customer{
		{Email: String("diane@writersclub.com"), FirstName: String("Diane"), LastName: String("Nguyen"), Integration: Int(100032), Domain: String("test"), Phone: String("16504173147"), CustomerCode: String("CUS_1uld4hluw0g2gn0"), Id: Int(63), CreatedAt: &createdAt1, UpdatedAt: &updatedAt1},
		{Email: String("todd@chavez.com"), Integration: Int(100032), Domain: String("test"), CustomerCode: String("CUS_soirsjdqkyjfwcr"), Id: Int(65), CreatedAt: &createdAt2, UpdatedAt: &updatedAt2},
	}
	if !cmp.Equal(customer, want) {
		t.Errorf("Customer.List returned %+v \n want %+v", customer, want)
	}
}

func TestCustomerService_Fetch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer/1173", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Customer retrieved",
		  "data": {
			  "integration": 100032,
			  "first_name": "Bojack",
			  "last_name": "Horseman",
			  "email": "bojack@horsinaround.com",
			  "phone": null,
			  "metadata": {
				"photos": [
				  {
					"type": "twitter",
					"typeId": "twitter",
					"typeName": "Twitter",
					"url": "https://d2ojpxxtu63wzl.cloudfront.net/static/61b1a0a1d4dda2c9fe9e165fed07f812_a722ae7148870cc2e33465d1807dfdc6efca33ad2c4e1f8943a79eead3c21311",
					"isPrimary": true
				  }
				]
			  },
			  "domain": "test",
			  "customer_code": "CUS_xnxdt6s1zg1f4nx",
			  "id": 1173,
			  "transactions": [],
			  "subscriptions": [],
			  "authorizations": [],
			  "createdAt": "2016-03-29T20:03:09.000Z",
			  "updatedAt": "2016-03-29T20:03:10.000Z"
			}
		}`)
	})

	customer, _, err := client.Customer.Fetch(context.Background(), "1173", nil)
	if err != nil {
		t.Errorf("Customer.Fetch returned error: %v", err)
	}
	createdAt := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)
	updatedAt := time.Date(2016, 3, 29, 20, 3, 10, 0, time.UTC)

	want := &Customer{FirstName: String("Bojack"), LastName: String("Horseman"), Email: String("bojack@horsinaround.com"), Integration: Int(100032), Metadata: Metadata{Photos: []Photo{{Type: String("twitter"), TypeId: String("twitter"), TypeName: String("Twitter"), URL: String("https://d2ojpxxtu63wzl.cloudfront.net/static/61b1a0a1d4dda2c9fe9e165fed07f812_a722ae7148870cc2e33465d1807dfdc6efca33ad2c4e1f8943a79eead3c21311"), IsPrimary: Bool(true)}}}, Domain: String("test"), CustomerCode: String("CUS_xnxdt6s1zg1f4nx"), Id: Int(1173), Transactions: []Transaction{}, Subscriptions: []Subscription{}, Authorizations: []Authorization{}, CreatedAt: &createdAt, UpdatedAt: &updatedAt}
	if !reflect.DeepEqual(customer, want) {
		t.Errorf("Customer.Fetch returned %+v, want %+v", customer, want)
	}
}

func TestCustomerService_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer/CUS_xnxdt6s1zg1f4nx", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Customer updated",
		  "data": {
			  "integration": 100032,
			  "first_name": "Bojack",
			  "last_name": "Horseman",
			  "email": "bojack@horsinaround.com",
			  "phone": null,
			  "metadata": {
				"photos": [
				  {
					"type": "twitter",
					"typeId": "twitter",
					"typeName": "Twitter",
					"url": "https://d2ojpxxtu63wzl.cloudfront.net/static/61b1a0a1d4dda2c9fe9e165fed07f812_a722ae7148870cc2e33465d1807dfdc6efca33ad2c4e1f8943a79eead3c21311",
					"isPrimary": true
				  }
				]
			  },
			  "domain": "test",
			  "customer_code": "CUS_xnxdt6s1zg1f4nx",
			  "id": 1173,
			  "transactions": [],
			  "subscriptions": [],
			  "authorizations": [],
			  "createdAt": "2016-03-29T20:03:09.000Z",
			  "updatedAt": "2016-03-29T20:03:10.000Z"
			}
		}`)
	})
	customerRequest := &CustomerRequest{FirstName: String("BoJack")}

	customer, _, err := client.Customer.Update(context.Background(), customerRequest, "CUS_xnxdt6s1zg1f4nx")
	if err != nil {
		t.Errorf("Customer.Update returned error: %v", err)
	}
	createdAt := time.Date(2016, 3, 29, 20, 3, 9, 0, time.UTC)
	updatedAt := time.Date(2016, 3, 29, 20, 3, 10, 0, time.UTC)

	want := &Customer{FirstName: String("Bojack"), LastName: String("Horseman"), Email: String("bojack@horsinaround.com"), Integration: Int(100032), Metadata: Metadata{Photos: []Photo{{Type: String("twitter"), TypeId: String("twitter"), TypeName: String("Twitter"), URL: String("https://d2ojpxxtu63wzl.cloudfront.net/static/61b1a0a1d4dda2c9fe9e165fed07f812_a722ae7148870cc2e33465d1807dfdc6efca33ad2c4e1f8943a79eead3c21311"), IsPrimary: Bool(true)}}}, Domain: String("test"), CustomerCode: String("CUS_xnxdt6s1zg1f4nx"), Id: Int(1173), Transactions: []Transaction{}, Subscriptions: []Subscription{}, Authorizations: []Authorization{}, CreatedAt: &createdAt, UpdatedAt: &updatedAt}
	if !reflect.DeepEqual(customer, want) {
		t.Errorf("Customer.Update returned %+v, want %+v", customer, want)
	}
}

func TestCustomerService_SetRiskAction(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer/set_risk_action", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Customer updated",
		  "data": {
			"first_name": "Peter",
			"last_name": "Griffin",
			"email": "peter@familyguy.com",
			"phone": null,
			"metadata": {},
			"domain": "test",
			"customer_code": "CUS_xr58yrr2ujlft9k",
			"risk_action": "allow",
			"id": 2109,
			"integration": 100032,
			"createdAt": "2016-01-26T13:43:38.000Z",
			"updatedAt": "2016-08-23T03:56:43.000Z"
		  }
		}`)
	})
	rap := &RiskActionPayload{CustomerCode: String("CUS_xr58yrr2ujlft9k"), RiskAction: Allow}

	customer, _, err := client.Customer.SetRiskAction(context.Background(), rap)
	if err != nil {
		t.Errorf("Customer.SetRiskAction returned error: %v", err)
	}
	createdAt := time.Date(2016, 1, 26, 13, 43, 38, 0, time.UTC)
	updatedAt := time.Date(2016, 8, 23, 3, 56, 43, 0, time.UTC)

	want := &Customer{FirstName: String("Peter"), LastName: String("Griffin"), Email: String("peter@familyguy.com"), Integration: Int(100032), Metadata: Metadata{}, Domain: String("test"), CustomerCode: String("CUS_xr58yrr2ujlft9k"), RiskAction: String("allow"), Id: Int(2109), CreatedAt: &createdAt, UpdatedAt: &updatedAt}
	if !reflect.DeepEqual(customer, want) {
		t.Errorf("Customer.SetRiskAction returned %+v, want %+v", customer, want)
	}
}

func TestCustomerService_DeactivateAuthorization(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/customer/deactivate_authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Authorization has been deactivated"
		}`)
	})
	auth := &Authorization{AuthorizationCode: String("AUTH_au6hc0de")}
	message, _, err := client.Customer.DeactivateAuthorization(context.Background(), auth)
	if err != nil {
		t.Errorf("Customer.DeactivateAuthorization returned error: %v", err)
	}
	want := &Message{Status: Bool(true), Message: String("Authorization has been deactivated")}
	if !reflect.DeepEqual(message, want) {
		t.Errorf("Customer.DeactivateAuthorization returned %+v, want %+v", message, want)
	}
}
