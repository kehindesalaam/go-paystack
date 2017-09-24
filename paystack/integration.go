package paystack

import (
	"context"
	"fmt"
)

type IntegrationService service

type PaymentSession struct {
	PaymentSessionTimeout *int `json:"payment_session_timeout, omitempty"`
}

//FetchPaymentSessionTimeout
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-payment-session-timeout
func (s *IntegrationService) FetchPaymentSessionTimeout(ctx context.Context) (*PaymentSession, *Response, error) {
	u := fmt.Sprintf("integration/payment_session_timeout")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(PaymentSession)
	MapDecoder(r.Data, m)
	return m, resp, nil
}

//UpdatePaymentSessionTimeout
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-payment-session-timeout
func (s *IntegrationService) UpdatePaymentSessionTimeout(ctx context.Context, options Options) (*PaymentSession, *Response, error) {
	u := fmt.Sprintf("integration/payment_session_timeout")

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(PaymentSession)
	MapDecoder(r.Data, m)
	return m, resp, nil
}
