package main

import (
	"context"
	"go-paystack/paystack"
)

func main() {
	ctx := context.Background()

	client := paystack.NewClient(nil, paystack.SecretKey("sk_test_87e6f6be4e0b68e9a286d14d89b813ac6353efa5"))
	//c := paystack.Customer{Email: paystack.String("kenny@testing.com"), FirstName: paystack.String("Kehinde"), LastName: paystack.String("Sxcvvxcalaam"), Phone: paystack.String("08161759814"), CustomerCode: paystack.String("CUS_ya21eig6qit06v9")}
	//newCustomer, _, err := client.Customer.CreateCustomer(ctx, &c)
	lo := paystack.Options{Page: 0, PerPage: 2}
	clist, _, err := client.Customer.List(ctx, &lo) //, &lo)
	//newCustomer, _, err := client.Customer.Fetch(ctx, "687027");
	//newCustomer, _, err := client.Customer.Update(ctx, &c)
	//r := paystack.RiskActionPayload{RiskAction: paystack.Deny, CustomerCode: paystack.String("CUS_ya21eig6qit06v9")}
	//newCustomer, _, err := client.Customer.SetRiskAction(ctx, &r)

	if err != nil {
		println(err.Error())
		return
	}
	//println(newCustomer.GetEmail())
	//println(newCustomer.GetIntegration())
	//println(newCustomer.GetDomain())
	//println(newCustomer.GetUpdatedAt().String())
	//println(newCustomer.GetId())
	for _, c := range clist {
		println(c.GetEmail())
	}
}
