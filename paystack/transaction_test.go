package paystack

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"github.com/google/go-cmp/cmp"
	"time"
)

func TestTransactionService_Initialize(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/transaction/initialize", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Authorization URL created",
		  "data": {
			"authorization_url": "https://standard.paystack.co/pay/0peioxfhpn",
			"access_code": "0peioxfhpn",
			"reference": "7PVGX8MEk85tgeEpVDtD"
		  }
		}`)
	})

	tranxRequest := TransactionRequest{Email: String("customer@email.com"), Reference:String("7PVGX8MEk85tgeEpVDtD")}

	tranx, _, err := client.Transaction.Initialize(context.Background(), &tranxRequest)
	if err != nil {
		t.Errorf("Transaction.Initialize returned error: %v", err)
	}

	want := &TransactionAuthorization{ AuthorizationUrl:String("https://standard.paystack.co/pay/0peioxfhpn"), AccessCode:String("0peioxfhpn"), Reference:String("7PVGX8MEk85tgeEpVDtD")}
	if !cmp.Equal(tranx, want) {
		t.Errorf("Transaction.Initialize returned %+v, want %+v", tranx, want)

	}
}

func TestTransactionService_Verify(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/transaction/verify/DG4uishudoq90LD", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
		   "status":true,
		   "message":"Verification successful",
		   "data":{
			  "amount":27000,
			  "currency":"NGN",
			  "transaction_date":"2016-10-01T11:03:09.000Z",
			  "status":"success",
			  "reference":"DG4uishudoq90LD",
			  "domain":"test",
			  "gateway_response":"Successful",
			  "message":null,
			  "channel":"card",
			  "ip_address":"41.1.25.1",
			  "log":{
				 "time_spent":9,
				 "attempts":1,
				 "authentication":null,
				 "errors":0,
				 "success":true,
				 "mobile":false,
				 "channel":null,
				 "history":[
					{
					   "type":"input",
					   "message":"Filled these fields: card number, card expiry, card cvv",
					   "time":7
					},
					{
					   "type":"action",
					   "message":"Attempted to pay",
					   "time":7
					}
				 ]
			  },
			  "fees":null,
			  "authorization":{
				 "authorization_code":"AUTH_8dfhjjdt",
				 "card_type":"visa",
				 "last4":"1381",
				 "exp_month":"08",
				 "exp_year":"2018",
				 "bin":"412345",
				 "bank":"TEST BANK",
				 "channel":"card",
				 "signature": "SIG_idyuhgd87dUYSHO92D",
				 "reusable":true,
				 "country_code":"NG"
			  },
			  "customer":{
				 "id":84312,
				 "customer_code":"CUS_hdhye17yj8qd2tx",
				 "first_name":"BoJack",
				 "last_name":"Horseman",
				 "email":"bojack@horseman.com"
			  },
			  "plan":"PLN_0as2m9n02cl0kp6"
		   }
		}`)
	})

	tranx, _, err := client.Transaction.Verify(context.Background(), "DG4uishudoq90LD")
	if err != nil {
		t.Errorf("Transaction.Verify returned error: %v", err)
	}
	tranxDate := time.Date(2016, 10, 01, 11, 3, 9, 0, time.UTC)

	//Metadata and Input Omitted in test
	want := &TransactionVerify{ Amount:Int(27000), Currency:String("NGN"), TransactionDate: &tranxDate, Status:String("success"), Reference:String("DG4uishudoq90LD"), Domain:String("test"), GatewayResponse:String("Successful"), Channel:String("card"), IpAddress:String("41.1.25.1"), Log:Log{TimeSpent:Int(9), Attempts: Int(1), Errors:Int(0), Success:Bool(true), Mobile:Bool(false), History:[]History{
		{Type:String("input"), Message:String("Filled these fields: card number, card expiry, card cvv"), Time:Int(7)},
		{Type:String("action"), Message:String("Attempted to pay"), Time:Int(7)}}},
		Authorization:Authorization{AuthorizationCode:String("AUTH_8dfhjjdt"), CardType:String("visa"), Last4:String("1381"), ExpMonth:String("08"), ExpYear:String("2018"), Bin:String("412345"), Bank:String("TEST BANK"), Channel:String("card"), Signature:String("SIG_idyuhgd87dUYSHO92D"), Reusable:Bool(true), CountryCode:String("NG")},
		Customer:Customer{Id:Int(84312), CustomerCode:String("CUS_hdhye17yj8qd2tx"), FirstName:String("BoJack"), LastName:String("Horseman"), Email:String("bojack@horseman.com")}, Plan:String("PLN_0as2m9n02cl0kp6")}
	if !cmp.Equal(tranx, want) {
		t.Errorf("Transaction.Verify returned %+v, want %+v", tranx, want)
	}
}
