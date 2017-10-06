package paystack

import (
	"context"
	"fmt"
	"time"
)

//PlanService handles the communication with the Plans related parts of the Paystack API
type PlanService service

type Plan struct {
	Name              *string            `json:"name, omitempty"`
	Description       *string            `json:"description, omitempty"`
	Amount            *int               `json:"amount, omitempty"`
	Interval          *string            `json:"interval, omitempty"`
	Domain            *string            `json:"domain, omitempty"`
	PlanCode          *string            `json:"plan_code, omitempty"`
	SendInvoices      *bool              `json:"send_invoices, omitempty"`
	SendSms           *bool              `json:"send_sms, omitempty"`
	HostedPage        *bool              `json:"hosted_page, omitempty"`
	Currency          *string            `json:"currency, omitempty"`
	InvoiceLimit      *string            `json:"invoice_limit, omitempty"`
	Id                *int               `json:"id, omitempty"`
	CreatedAt         *time.Time         `json:"created_at, omitempty"`
	UpdatedAt         *time.Time         `json:"updated_at, omitempty"`
	Subscriptions     []PlanSubscription `json:"subscriptions, omitempty"`
	Integration       *int               `json:"integration, omitempty, omitempty"`
	HostedPageURL     *string            `json:"hosted_page_url, omitempty"`
	HostedPageSummary *string            `json:"hosted_page_summary, omitempty"`
}

type PlanSubscription struct {
	Customer         *int        `json:"customer"`
	Plan             *int        `json:"plan, omitempty"`
	Integration      *int        `json:"integration, omitempty"`
	Domain           *string     `json:"domain, omitempty"`
	Start            *int64      `json:"start, omitempty"`
	Status           *string     `json:"status, omitempty"`
	Quantity         *int        `json:"quantity, omitempty"`
	Amount           *int        `json:"amount, omitempty"`
	SubscriptionCode *string     `json:"subscription_code, omitempty"`
	EmailToken       *string     `json:"email_token, omitempty"`
	Authorization    *int        `json:"authorization, omitempty"`
	EasyCronId       *int        `json:"easy_cron_id, omitempty"`
	CronExpression   *string     `json:"cron_expression, omitempty"`
	NextPaymentDate  *time.Time  `json:"next_payment_date, omitempty"`
	OpenInvoice      interface{} `json:"open_invoice, omitempty"`
	Id               *int        `json:"id, omitempty"`
	CreatedAt        *time.Time  `json:"created_at, omitempty"`
	UpdatedAt        *time.Time  `json:"updated_at, omitempty"`
}

type PlanRequest struct {
	Name         *string `json:"name, omitempty"`
	Description  *string `json:"description, omitempty"`
	Amount       *int    `json:"amount, omitempty"`
	Interval     *string `json:"interval, omitempty"`
	SendInvoices *bool   `json:"send_invoices, omitempty"`
	SendSms      *bool   `json:"send_sms, omitempty"`
	Currency     *string `json:"currency, omitempty"`
	InvoiceLimit *string `json:"invoice_limit, omitempty"`
}

// Create returns a new plan
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-plan
func (s *PlanService) Create(ctx context.Context, p *PlanRequest) (*Plan, *Response, error) {
	u := fmt.Sprintf("plan")
	req, err := s.client.NewRequest("POST", u, p)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var pr Plan
	mapDecoder(r.Data, pr)
	return &pr, resp, nil
}

// List returns an array all created plans
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-plans
func (s *PlanService) List(ctx context.Context, opt *PlanOptions) ([]Plan, *Response, error) {
	u := fmt.Sprintf("plan")
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

	var pa []Plan
	p := new(Plan)
	for _, x := range lr.Data {
		mapDecoder(x, p)
		pa = append(pa, *p)
	}
	return pa, resp, nil
}

// Fetch returns a plan with the passed id parameter
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-plan
func (s *PlanService) Fetch(ctx context.Context, id string) (*Plan, *Response, error) {
	u := fmt.Sprintf("plan/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var pr Plan
	mapDecoder(r.Data, pr)
	return &pr, resp, nil
}

// Update updates a plan model with the supplied id
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-plan
func (s *PlanService) Update(ctx context.Context, sa *Plan, id string) (*PlanSubscription, *Response, error) {
	u := fmt.Sprintf("plan/" + id)
	req, err := s.client.NewRequest("PUT", u, sa)
	if err != nil {
		return nil, nil, err
	}
	pr := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, pr)
	if err != nil {
		return nil, resp, err
	}
	var p PlanSubscription
	mapDecoder(pr.Data, p)
	return &p, resp, nil
}
