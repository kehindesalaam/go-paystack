package paystack

import (
	"context"
	"fmt"
)

type BalanceService service

type Balance struct {
	Currency *string `json:"currency, omitempty"`
	Balance  *int    `json:"balance, omitempty"`
}

//Check returns an array of balances
//
// Paystack API reference:
// https://developers.paystack.co/reference#check-balance
func (s *BalanceService) Check(ctx context.Context) ([]Balance, *Response, error) {
	u := fmt.Sprintf("balance")
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	lr := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, lr)
	if err != nil {
		return nil, resp, err
	}
	var ba []Balance
	b := new(Balance)
	for _, x := range lr.Data {
		MapDecoder(x, b)
		ba = append(ba, *b)
	}
	return ba, resp, nil
}
