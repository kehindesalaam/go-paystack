package paystack

import (
	"context"
	"fmt"
)

type ChargeService service

type ChargeRequest struct {
	Email             *string  `json:"email, omitempty"`
	Card              Card     `json:"card, omitempty"`
	Bank              Bank     `json:"bank, omitempty"`
	AuthorizationCode *string  `json:"authorization_code, omitempty"`
	Pin               *string  `json:"pin, omitempty"`
	Metadata          Metadata `json:"metadata, omitempty"`
}

type Card struct {
	Number      *string `json:"number, omitempty"`
	CVV         *string `json:"cvv, omitempty"`
	ExpiryMonth *string `json:"expiry_month, omitempty"`
	ExpiryYear  *string `json:"expiry_year, omitempty"`
}

type PinRequest struct {
	Pin       *string `json:"pin, omitempty"`
	Reference *string `json:"reference, omitempty"`
}

type OTPRequest struct {
	OTP       *string `json:"otp, omitempty"`
	Reference *string `json:"reference, omitempty"`
}

type PhoneRequest struct {
	OTP       *string `json:"otp, omitempty"`
	Reference *string `json:"reference, omitempty"`
}

type BirthdayRequest struct {
	Birthday  *string `json:"birthday, omitempty"`
	Reference *string `json:"reference, omitempty"`
}

// Tokenize
//
// Paystack API reference:
// https://developers.paystack.co/reference#charge-tokenize
func (s *ChargeService) Tokenize(ctx context.Context, request *ChargeRequest) (*Authorization, *Response, error) {
	u := fmt.Sprintf("charge/tokenize")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Authorization)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// Charge
//
// Paystack API reference:
// https://developers.paystack.co/reference#charge
func (s *ChargeService) Charge(ctx context.Context, request *ChargeRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// SubmitPIN
//
// Paystack API reference:
// https://developers.paystack.co/reference#submit-pin
func (s *ChargeService) SubmitPIN(ctx context.Context, request *PinRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge/submit_pin")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// SubmitOTP
//
// Paystack API reference:
// https://developers.paystack.co/reference#submit-otp
func (s *ChargeService) SubmitOTP(ctx context.Context, request *OTPRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge/submit_pin")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// SubmitPhone
//
// Paystack API reference:
// https://developers.paystack.co/reference#submit-phone
func (s *ChargeService) SubmitPhone(ctx context.Context, request *PhoneRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge/submit_pin")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// SubmitBirthday
//
// Paystack API reference:
// https://developers.paystack.co/reference#submit-birthday
func (s *ChargeService) SubmitBirthday(ctx context.Context, request *BirthdayRequest) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge/submit_pin")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}

// CheckPending
//
// Paystack API reference:
// https://developers.paystack.co/reference#check-pending-charge
func (s *ChargeService) CheckPending(ctx context.Context, reference *string) (*Transaction, *Response, error) {
	u := fmt.Sprintf("charge/submit_pin")
	req, err := s.client.NewRequest("GET", u, reference)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Transaction)
	mapDecoder(r.Data, m)
	return m, resp, nil
}
