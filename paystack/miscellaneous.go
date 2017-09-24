package paystack

import (
	"context"
	"fmt"
	"time"
)

//MiscellaneousService handles the communication with the Miscellaneous related parts of the Paystack API
type MiscellaneousService service

type Bank struct {
	Name      *string     `json:"name, omitempty"`
	Slug      *string     `json:"slug, omitempty"`
	Code      *string     `json:"code, omitempty"`
	Longcode  *string     `json:"longcode, omitempty"`
	Gateway   *string     `json:"gateway, omitempty"`
	Active    *bool       `json:"active, omitempty"`
	IsDeleted interface{} `json:"is_deleted, omitempty"`
	Id        *int        `json:"id, omitempty"`
	CreatedAt *time.Time  `json:"created_at, omitempty"`
	UpdatedAt *time.Time  `json:"updated_at, omitempty"`
}

type Bin struct {
	Bin          *string `json:"bin, omitempty"`
	Brand        *string `json:"brand, omitempty"`
	SubBrand     *string `json:"sub_brand, omitempty"`
	CountryCode  *string `json:"country_code, omitempty"`
	CountryName  *string `json:"country_name, omitempty"`
	CardType     *string `json:"card_type, omitempty"`
	Bank         *string `json:"bank, omitempty"`
	LinkedBankId *int    `json:"linked_bank_id, omitempty"`
}

type BvnData struct {
	FirstName *string `json:"first_name, omitempty"`
	LastName  *string `json:"last_name, omitempty"`
	Dob       *string `json:"dob, omitempty"`
	Mobile    *string `json:"mobile, omitempty"`
	Bvn       *string `json:"bvn, omitempty"`
}

type AccountData struct {
	AccountNumber *string `json:"account_number, omitempty"`
	AccountName   *string `json:"account_name, omitempty"`
}

// ListBanks lists all banks
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-banks
func (s *MiscellaneousService) ListBanks(ctx context.Context, opt *Options) ([]Bank, *Response, error) {
	u := fmt.Sprintf("bank")
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
	var ba []Bank
	b := new(Bank)
	for _, x := range lr.Data {
		MapDecoder(x, b)
		ba = append(ba, *b)
	}
	return ba, resp, nil
}

// ResolveCardBin
//
// Paystack API reference:
// https://developers.paystack.co/reference#resolve-card-bin
func (s *MiscellaneousService) ResolveCardBin(ctx context.Context, id string) (*Bin, *Response, error) {
	u := fmt.Sprintf("decision/bin/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	b := new(Bin)
	MapDecoder(r.Data, b)
	return b, resp, nil
}

// ResolveBvn
//
// Paystack API reference:
// https://developers.paystack.co/reference#resolve-bvn
func (s *MiscellaneousService) ResolveBvn(ctx context.Context, id string) (*BvnData, *Response, error) {
	u := fmt.Sprintf("bank/resolve_bvn/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	b := new(BvnData)
	MapDecoder(r.Data, b)
	return b, resp, nil
}

// ResolveAccountNumber
//
// Paystack API reference:
// https://developers.paystack.co/reference#resolve-account-number
func (s *MiscellaneousService) ResolveAccountNumber(ctx context.Context, opt *Options) (*AccountData, *Response, error) {
	u := fmt.Sprintf("bank/resolve")
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	c := new(AccountData)
	MapDecoder(r.Data, c)
	return c, resp, nil
}
