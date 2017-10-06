package paystack

import (
	"context"
	"fmt"

	"time"
)

//SubaccountService handles the communication with the Subaccounts related parts of the Paystack API
type SubaccountService service

type Subaccount struct {
	Integration         *int       `json:"integration, omitempty"`
	Domain              *string    `json:"domain, omitempty"`
	SubaccountCode      *string    `json:"subaccount_code, omitempty"`
	BusinessName        *string    `json:"business_name, omitempty"`
	Description         *string    `json:"description, omitempty"`
	PrimaryContactName  *string    `json:"primary_contact_name, omitempty"`
	PrimaryContactEmail *string    `json:"primary_contact_email, omitempty"`
	PrimaryContactPhone *string    `json:"primary_contact_phone, omitempty"`
	Metadata            Metadata   `json:"metadata, omitempty"`
	PercentageCharge    *float32   `json:"percentage_charge, omitempty"`
	IsVerified          *bool      `json:"is_verified, omitempty"`
	SettlementBank      *string    `json:"settlement_bank, omitempty"`
	AccountNumber       *string    `json:"account_number, omitempty"`
	SettlementSchedule  *string    `json:"settlement_schedule, omitempty"`
	Active              *bool      `json:"active, omitempty"`
	Migrate             *bool      `json:"migrate, omitempty"`
	Id                  *int       `json:"id, omitempty"`
	CreatedAt           *time.Time `json:"created_at, omitempty"`
	UpdatedAt           *time.Time `json:"updated_at, omitempty"`
}

type SubaccountRequest struct {
	BusinessName        *string  `json:"business_name, omitempty"`
	PrimaryContactName  *string  `json:"primary_contact_name, omitempty"`
	PrimaryContactEmail *string  `json:"primary_contact_email, omitempty"`
	PrimaryContactPhone *string  `json:"primary_contact_phone, omitempty"`
	Metadata            Metadata `json:"metadata, omitempty"`
	PercentageCharge    *float32 `json:"percentage_charge, omitempty"`
	SettlementBank      *string  `json:"settlement_bank, omitempty"`
	AccountNumber       *string  `json:"account_number, omitempty"`
	SettlementSchedule  *string  `json:"settlement_schedule, omitempty"`
}

// Create returns a new subaccount
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-subaccount
func (s *SubaccountService) Create(ctx context.Context, sr *SubaccountRequest) (*Subaccount, *Response, error) {
	u := fmt.Sprintf("subaccount")
	req, err := s.client.NewRequest("POST", u, sr)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	var sa Subaccount
	mapDecoder(r.Data, sa)
	return &sa, resp, nil
}

// List returns an array of all created subaccounts
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-subaccounts
func (s *SubaccountService) List(ctx context.Context, opt *ListOptions) ([]Subaccount, *Response, error) {
	u := fmt.Sprintf("subaccount")
	//Response is erroneous if opt.Page or opt.PerPage = 0
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, r)

	if err != nil {
		return nil, resp, err
	}

	var saa []Subaccount
	sa := new(Subaccount)
	for _, x := range r.Data {
		mapDecoder(x, sa)
		saa = append(saa, *sa)
	}
	return saa, resp, nil
}

// Fetch returns a subaccount with the passed id
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-subaccount
func (s *SubaccountService) Fetch(ctx context.Context, id string) (*Subaccount, *Response, error) {
	u := fmt.Sprintf("subaccount/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var sa Subaccount
	mapDecoder(r.Data, sa)
	return &sa, resp, nil
}

// Update updates a subaccount model with the subaccount
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-subaccount
func (s *SubaccountService) Update(ctx context.Context, sa *SubaccountRequest, id string) (*Subaccount, *Response, error) {
	u := fmt.Sprintf("subaccount/" + id)
	req, err := s.client.NewRequest("PUT", u, sa)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var sar Subaccount
	mapDecoder(r.Data, sar)
	return &sar, resp, nil
}
