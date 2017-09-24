// Copyright 2017 The go-paystack AUTHORS. All rights reserved.

package paystack

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

//TransactionService handles the communication with the Transactions related parts of the Paystack API
type TransactionService service

type TransactionRequest struct {
	CallbackUrl       *string  `json:"callback_url, omitempty"`
	Reference         *string  `json:"reference, omitempty"`
	AuthorizationCode *string  `json:"authorization_code, omitempty"`
	Amount            *string  `json:"amount, omitempty"`
	Currency          *string  `json:"currency"`
	Email             *string  `json:"email, omitempty"`
	Plan              *string  `json:"plan, omitempty"`
	InvoiceLimit      *int32   `json:"invoice_limit, omitempty"`
	Metadata          Metadata `json:"metadata, omitempty"`
	Subaccount        *string  `json:"subaccount, omitempty"`
	TransactionCharge *int32   `json:"transaction_charge, omitempty"`
	Bearer            *string  `json:"bearer, omitempty"`
	Channels          []string `json:"channels, omitempty"`
}

type Transaction struct {
	Amount          *int          `json:"amount, omitempty"`
	Currency        *string       `json:"currency, omitempty"`
	TransactionDate *time.Time    `json:"transaction_date, omitempty"`
	Status          *string       `json:"status, omitempty"`
	Reference       *string       `json:"reference, omitempty"`
	Domain          *string       `json:"domain, omitempty"`
	Metadata        Metadata      `json:"metadata, omitempty"`
	GatewayResponse *string       `json:"gateway_response, omitempty"`
	Message         *string       `json:"message, omitempty"`
	Channel         *string       `json:"channel, omitempty"`
	IpAddress       *string       `json:"ip_address, omitempty"`
	Log             Log           `json:"log, omitempty"`
	Fees            *int          `json:"fees, omitempty"`
	Authorization   Authorization `json:"authorization, omitempty"`
	Customer        Customer      `json:"customer, omitempty"`
	Plan            Plan          `json:"plan, omitempty"`
	Id              *int          `json:"id, omitempty"`
	PaidAt          *time.Time    `json:"paid_at, omitempty"`
	CreatedAt       *time.Time    `json:"created_at, omitempty"`
	FeesSplit       *int          `json:"fees_split, omitempty"`
	Subaccount      Subaccount    `json:"subaccount, omitempty"`
}

type TransactionOptions struct {
	Options
	Customer int32     `json:"customer, omitempty"`
	Status   string    `json:"status, omitempty"`
	From     time.Time `json:"from, omitempty"`
	To       time.Time `json:"to, omitempty"`
	Amount   string    `json:"amount, omitempty"`
}

type TransactionTimeline struct {
	TimeSpent      *int          `json:"time_spent, omitempty"`
	Attempts       *int          `json:"attempts, omitempty"`
	Authentication *string       `json:"authentication, omitempty"`
	Errors         *int          `json:"errors, omitempty"`
	Success        *bool         `json:"success, omitempty"`
	Mobile         *bool         `json:"mobile, omitempty"`
	Input          []interface{} `json:"input, omitempty"`
	Channel        *string       `json:"channel, omitempty"`
	History        History       `json:"history, omitempty"`
}

type TransactionTotal struct {
	TotalTransactions     int               `json:"total_transactions, omitempty"`
	UniqueCustomers       int               `json:"unique_customers, omitempty"`
	TotalVolume           int64             `json:"total_volume, omitempty"`
	TotalVolumeByCurrency []FieldByCurrency `json:"total_volume_by_currency, omitempty"`
}

type ExportRequest struct {
	From        *time.Time `json:"from, omitempty"`
	To          *time.Time `json:"to, omitempty"`
	Settled     *bool      `json:"settled, omitempty"`
	PaymentPage *int32     `json:"payment_page, omitempty"`
	Customer    *int32     `json:"customer, omitempty"`
	Currency    *string    `json:"currency, omitempty"`
	Settlement  *string    `json:"settlement, omitempty"`
	Amount      *int32     `json:"amount, omitempty"`
	Status      *string    `json:"status, omitempty"`
}

type ExportPath struct {
	Path string `json:"path"`
}
type Reauthorization struct {
	ReauthorizationUrl *string `json:"reauthorization_url, omitempty"`
	Reference          *string `json:"reference, omitempty"`
}

// InitializeTransaction
//
// Paystack API reference:
// https://developers.paystack.co/reference#initialize-a-transaction
func (s *TransactionService) InitializeTransaction(ctx context.Context, tr *TransactionRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("transaction/initialize")
	req, err := s.client.NewRequest("POST", u, tr)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	var t Transaction
	mapstructure.Decode(r.Data, t)
	return &t, resp, nil
}

//VerifyTransaction creates a new customer
//
// Paystack API reference:
// https://developers.paystack.co/reference#initialize-a-transaction
func (s *TransactionService) VerifyTransaction(ctx context.Context, reference string) (*Transaction, *Response, error) {
	u := fmt.Sprintf("transaction/verify" + reference)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	var t Transaction
	mapstructure.Decode(r.Data, t)
	return &t, resp, nil
}

// ListTransactions lists all transactions
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-transactions
func (s *TransactionService) ListTransactions(ctx context.Context, opt *TransactionOptions) ([]Transaction, *Response, error) {
	u := fmt.Sprintf("transaction")
	//Response is erroneous if opt.Page or opt.PerPage = 0
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, r)

	if err != nil {
		return nil, resp, err
	}

	var ta []Transaction
	t := new(Transaction)
	for _, x := range r.Data {
		MapDecoder(x, t)
		ta = append(ta, *t)
	}
	return ta, resp, nil
}

// FetchTransaction fetches a transaction
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-transaction
func (s *TransactionService) FetchTransaction(ctx context.Context, id string) (*Transaction, *Response, error) {
	u := fmt.Sprintf("transaction/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var t Transaction
	mapstructure.Decode(r.Data, t)
	return &t, resp, nil
}

//ChargeAuthorization
//
// Paystack API reference:
// https://developers.paystack.co/reference#charge-authorization
func (s *TransactionService) ChargeAuthorization(ctx context.Context, tr *TransactionRequest) (*Transaction, *Response, error) {

	u := fmt.Sprintf("transaction/charge_authorization")
	req, err := s.client.NewRequest("POST", u, tr)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var t Transaction
	mapstructure.Decode(r.Data, t)
	return &t, resp, nil
}

// Timeline fetches a transaction timeline
//
// Paystack API reference:
// https://developers.paystack.co/reference#view-transaction-timeline
func (s *TransactionService) Timeline(ctx context.Context, id string) (*TransactionTimeline, *Response, error) {
	u := fmt.Sprintf("transaction/timeline" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var tt TransactionTimeline
	mapstructure.Decode(r.Data, tt)
	return &tt, resp, nil
}

// Totals fetches a transaction timeline
//
// Paystack API reference:
// https://developers.paystack.co/reference#transaction-totals
func (s *TransactionService) Totals(ctx context.Context, opt *TransactionOptions) (*TransactionTotal, *Response, error) {
	u := fmt.Sprintf("transaction/totals")

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
	var t TransactionTotal
	mapstructure.Decode(r.Data, t)
	return &t, resp, nil
}

// Export a transaction
//
// Paystack API reference:
// https://developers.paystack.co/reference#export-transactions
func (s *TransactionService) ExportTransactions(ctx context.Context) (*ExportPath, *Response, error) {
	u := fmt.Sprintf("transaction/export")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var ep ExportPath
	mapstructure.Decode(r.Data, ep)
	return &ep, resp, nil
}

// RequestReauthorization
//
// Paystack API reference:
// https://developers.paystack.co/reference#request-reauthorization
func (s *TransactionService) RequestReauthorization(ctx context.Context, opt *TransactionRequest) (*Reauthorization, *Response, error) {
	u := fmt.Sprintf("transaction/request_reauthorization")

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var ep Reauthorization
	mapstructure.Decode(r.Data, ep)
	return &ep, resp, nil
}

// CheckAuthorization
//
// Paystack API reference:
// https://developers.paystack.co/reference#check-authorization
func (s *TransactionService) CheckAuthorization(ctx context.Context, opt *TransactionRequest) (*FieldByCurrency, *Response, error) {
	u := fmt.Sprintf("transaction/check_authorization ")

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	var fbc FieldByCurrency
	mapstructure.Decode(r.Data, fbc)
	return &fbc, resp, nil
}
