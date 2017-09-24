package paystack

import (
	"context"
	"fmt"
	"time"
)

//SettlementService handles the communication with the Settlements related parts of the Paystack API
type SettlementService service

type Settlement struct {
	Integration *int        `json:"integration, omitempty"`
	Subaccount  Subaccount  `json:"subaccount, omitempty"`
	SettledBy   interface{} `json:"settled_by, omitempty"`
	SettledDate *time.Time  `json:"settled_date, omitempty"`
	Domain      *string     `json:"domain, omitempty"`
	TotalAmount *int        `json:"total_amount, omitempty"`
	Status      *string     `json:"status, omitempty"`
	Id          *int        `json:"id, omitempty"`
	CreatedAt   *time.Time  `json:"created_at, omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at, omitempty"`
}

// Fetch gets all settlements made to your bank accounts
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-settlements
func (s *SettlementService) Fetch(ctx context.Context, opt *Options) ([]Settlement, *Response, error) {
	u := fmt.Sprintf("settlement")
	//Response is erroneous if opt.Page or opt.PerPage = 0
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	lr := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, lr)

	if err != nil {
		return nil, resp, err
	}
	var sa []Settlement
	c := new(Settlement)
	for _, x := range lr.Data {
		MapDecoder(x, c)
		sa = append(sa, *c)
	}
	return sa, resp, nil
}
