package paystack

import (
	"context"
	"fmt"

	"time"
)

//SubscriptionService handles the communication with the Subscriptions related parts of the Paystack API
type SubscriptionService service

type Subscription struct {
	Customer         Customer      `json:"customer"`
	Plan             Plan          `json:"plan"`
	Integration      *int          `json:"integration, omitempty"`
	Authorization    Authorization `json:"authorization"`
	Domain           *string       `json:"domain, omitempty"`
	Start            *int64        `json:"start, omitempty"`
	Status           *string       `json:"status, omitempty"`
	Quantity         *int          `json:"quantity, omitempty"`
	Amount           *int          `json:"amount, omitempty"`
	SubscriptionCode *string       `json:"subscription_code, omitempty"`
	EmailToken       *string       `json:"email_token, omitempty"`
	EasyCronId       *int          `json:"easy_cron_id, omitempty"`
	CronExpression   *string       `json:"cron_expression, omitempty"`
	NextPaymentDate  *time.Time    `json:"next_payment_date, omitempty"`
	OpenInvoice      *string       `json:"open_invoice, omitempty"`
	Id               *int          `json:"id, omitempty"`
	CreatedAt        *time.Time    `json:"created_at, omitempty"`
	UpdatedAt        *time.Time    `json:"updated_at, omitempty"`
}

type SubscriptionResponse struct {
	Customer         *int       `json:"customer, omitempty"`
	Plan             *int       `json:"plan, omitempty"`
	Integration      *int       `json:"integration, omitempty"`
	Domain           *string    `json:"domain, omitempty"`
	Start            *int64     `json:"start, omitempty"`
	Status           *string    `json:"status, omitempty"`
	Quantity         *int       `json:"quantity, omitempty"`
	Amount           *int       `json:"amount, omitempty"`
	Authorization    *int       `json:"authorization, omitempty"`
	SubscriptionCode *string    `json:"subscription_code, omitempty"`
	EmailToken       *string    `json:"email_token, omitempty"`
	Id               *int       `json:"id, omitempty"`
	CreatedAt        *time.Time `json:"created_at, omitempty"`
	UpdatedAt        *time.Time `json:"updated_at, omitempty"`
}

type SubscriptionRequest struct {
	Customer      *string    `json:"customer, omitempty"`
	Plan          *string    `json:"plan, omitempty"`
	Authorization *string    `json:"authorization, omitempty"`
	StartDate     *time.Time `json:"start_date, omitempty"`
	Code          *string    `json:"code, omitempty"`
	Token         *string    `json:"token, omitempty"`
}

// Create creates a new subscription
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-subscription
func (s *SubscriptionService) Create(ctx context.Context, sa *Subscription) (*SubscriptionResponse, *Response, error) {
	u := fmt.Sprintf("subscription")
	req, err := s.client.NewRequest("POST", u, sa)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	var c SubscriptionResponse
	mapDecoder(r.Data, c)
	return &c, resp, nil
}

// List lists all created subscriptions
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-subscriptions
func (s *SubscriptionService) List(ctx context.Context, opt *SubscriptionOptions) ([]Subscription, *Response, error) {
	u := fmt.Sprintf("subscription")
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

	var sa []Subscription
	st := new(Subscription)
	for _, x := range lr.Data {
		mapDecoder(x, st)
		sa = append(sa, *st)
	}
	return sa, resp, nil
}

// Fetch fetches a subscription
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-subscription
func (s *SubscriptionService) Fetch(ctx context.Context, id string) (*Subscription, *Response, error) {
	u := fmt.Sprintf("subscription/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var st Subscription
	err = mapDecoder(r.Data, &st)
	if err != nil {
		return nil, resp, err
	}
	return &st, resp, nil
}

// Disable disables a subscription model with the supplied parameters
//
// Paystack API reference:
// https://developers.paystack.co/reference#disable-subscription
func (s *SubscriptionService) Disable(ctx context.Context, sa *SubscriptionRequest) (*Message, *Response, error) {
	u := fmt.Sprintf("subscription/disable")

	u, err := addOptions(u, sa)
	if err != nil {
		return nil, nil, err
	}
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

// Enable enables a subscription model with the supplied parameters
//
// Paystack API reference:
// https://developers.paystack.co/reference#enable-subscription
func (s *SubscriptionService) Enable(ctx context.Context, sa *SubscriptionRequest) (*Message, *Response, error) {
	u := fmt.Sprintf("subscription/enable")

	u, err := addOptions(u, sa)
	if err != nil {
		return nil, nil, err
	}
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
