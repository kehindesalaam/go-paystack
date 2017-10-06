package paystack

import (
	"context"
	"fmt"
	"time"
)

//TransferService handles the communication with the Transfers related parts of the Paystack API
type TransferService service

type Transfer struct {
	Integration   *int              `json:"integration, omitempty"`
	Recipient     TransferRecipient `json:"recipient, omitempty"`
	Domain        *string           `json:"domain, omitempty"`
	Amount        *int              `json:"amount, omitempty"`
	Currency      *string           `json:"currency, omitempty"`
	Source        *string           `json:"source, omitempty"`
	SourceDetails *string           `json:"source_details, omitempty"`
	Reason        *string           `json:"reason, omitempty"`
	Status        *string           `json:"status, omitempty"`
	Failures      interface{}       `json:"failures, omitempty"`
	TransferCode  *string           `json:"transfer_code, omitempty"`
	Id            *int              `json:"id, omitempty"`
	CreatedAt     *time.Time        `json:"created_at, omitempty"`
	UpdatedAt     *time.Time        `json:"updated_at, omitempty"`
}

type TransferRequest struct {
	Recipient    *string `json:"recipient, omitempty"`
	Amount       *int    `json:"amount, omitempty"`
	Currency     *string `json:"currency, omitempty"`
	Source       *string `json:"source, omitempty"`
	Reason       *string `json:"reason, omitempty"`
	TransferCode *string `json:"transfer_code, omitempty"`
}

type FinalizeTransferRequest struct {
	TransferCode *string `json:"transfer_code, omitempty"`
	OTP          *string `json:"otp, omitempty"`
}

type BulkTransferRequest struct {
	Currency  *string           `json:"currency, omitempty"`
	Source    *string           `json:"source, omitempty"`
	Transfers []TransferRequest `json:"transfers, omitempty"`
}

//Initiate creates a new transfer
//
// Paystack API reference:
// https://developers.paystack.co/reference#initiate-transfer
func (s *TransferService) Initiate(ctx context.Context, t *TransferRequest) (*Transfer, *Response, error) {
	u := fmt.Sprintf("transfer")
	req, err := s.client.NewRequest("POST", u, t)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	c := new(Transfer)
	mapDecoder(r.Data, c)
	return c, resp, nil
}

// List returns all created transfers
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-transfers
func (s *TransferService) List(ctx context.Context, opt *ListOptions) ([]Transfer, *Response, error) {
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

	lr := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, lr)

	if err != nil {
		return nil, resp, err
	}
	var ta []Transfer
	c := new(Transfer)
	for _, x := range lr.Data {
		mapDecoder(x, c)
		ta = append(ta, *c)
	}
	return ta, resp, nil
}

// Fetch fetches a transfer
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-transfer
func (s *TransferService) Fetch(ctx context.Context, id string) (*Transfer, *Response, error) {
	u := fmt.Sprintf("transfer/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	c := new(Transfer)
	mapDecoder(r.Data, c)
	return c, resp, nil
}

// Finalize
//
// Paystack API reference:
// https://developers.paystack.co/reference#finalize-transfer
func (s *TransferService) Finalize(ctx context.Context, sa *FinalizeTransferRequest) (*Response, error) {
	u := fmt.Sprintf("transfer/finalize_transfer")
	u, err := addOptions(u, sa)
	req, err := s.client.NewRequest("POST", u, sa)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

//InitiateBulkTransfer creates a new transfer
//
// Paystack API reference:
// https://developers.paystack.co/reference#initiate-bulk-transfer
func (s *TransferService) InitiateBulkTransfer(ctx context.Context, t *BulkTransferRequest) (*Message, *Response, error) {
	u := fmt.Sprintf("transfer")
	u, err := addOptions(u, t)
	req, err := s.client.NewRequest("POST", u, t)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// ResendOTP
//
// Paystack API reference:
// https://developers.paystack.co/reference#resend-otp-for-transfer
func (s *TransferService) ResendOTP(ctx context.Context, sa *TransferRequest) (*Message, *Response, error) {
	u := fmt.Sprintf("transfer/resend_otp")
	req, err := s.client.NewRequest("POST", u, sa)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// DisableOTP
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-transfer
func (s *TransferService) DisableOTP(ctx context.Context) (*Message, *Response, error) {
	u := fmt.Sprintf("transfer/disable_otp")
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// DisableOTPFinalize
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-transfer
func (s *TransferService) DisableOTPFinalize(ctx context.Context) (*Message, *Response, error) {
	u := fmt.Sprintf("transfer/disable_otp_finalize")
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// EnableOTP
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-transfer
func (s *TransferService) EnableOTP(ctx context.Context) (*Message, *Response, error) {
	u := fmt.Sprintf("transfer/enable_otp")
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	mapDecoder(r.Data, m)
	return m, resp, nil
}
