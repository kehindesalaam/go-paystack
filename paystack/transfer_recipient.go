package paystack

import (
	"context"
	"fmt"
	"time"
)

//TransferService handles the communication with the Transfers related parts of the Paystack API
type TransferRecipientService service

type TransferRecipientRequest struct {
	Type          *string  `json:"type, omitempty"`
	Currency      *string  `json:"currency, omitempty"`
	Name          *string  `json:"name, omitempty"`
	Description   *string  `json:"description"`
	Metadata      Metadata `json:"metadata, omitempty"`
	AccountNumber *string  `json:"account_number, omitempty"`
	BankCode      *string  `json:"bank_code, omitempty"`
}

type TransferRecipient struct {
	Domain        *string                  `json:"domain, omitempty"`
	Type          *string                  `json:"type, omitempty"`
	Currency      *string                  `json:"currency, omitempty"`
	Name          *string                  `json:"name, omitempty"`
	Details       TransferRecipientDetails `json:"details, omitempty"`
	Description   *string                  `json:"description"`
	Metadata      Metadata                 `json:"metadata, omitempty"`
	RecipientCode *string                  `json:"recipient_code, omitempty"`
	Active        *bool                    `json:"active, omitempty"`
	Id            *int                     `json:"id, omitempty"`
	Integration   *int                     `json:"integration, omitempty, omitempty"`
	CreatedAt     *time.Time               `json:"created_at, omitempty"`
	UpdatedAt     *time.Time               `json:"updated_at, omitempty"`
}

type TransferRecipientDetails struct {
	AccountNumber *string `json:"account_number, omitempty"`
	AccountName   *string `json:"account_name, omitempty"`
	BankCode      *string `json:"bank_code, omitempty"`
	BankName      *string `json:"bank_name, omitempty"`
}

//Create creates a new transfer recipient
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-transfer-recipient
func (s *TransferRecipientService) Create(ctx context.Context, t *TransferRecipientRequest) (*TransferRecipient, *Response, error) {
	u := fmt.Sprintf("transferrecipient")
	req, err := s.client.NewRequest("POST", u, t)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	tr := new(TransferRecipient)
	MapDecoder(r.Data, tr)
	return tr, resp, nil
}

// List returns all created transfer recipients
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-transfers
func (s *TransferRecipientService) List(ctx context.Context, opt *Options) ([]TransferRecipient, *Response, error) {
	u := fmt.Sprintf("transfer")
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
	var ca []TransferRecipient
	c := new(TransferRecipient)
	for _, x := range r.Data {
		MapDecoder(x, c)
		ca = append(ca, *c)
	}
	return ca, resp, nil
}
