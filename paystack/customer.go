// Copyright 2017 The go-paystack AUTHORS. All rights reserved.

package paystack

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//enum that allows selection of the risk action type for customers
type RiskAction int

const (
	Allow RiskAction = 1 + iota
	Deny
)

var riskActions = [...]string{
	"allow",
	"deny",
}

//returns string equivalent of a RiskAction
func (r RiskAction) String() string { return riskActions[r-1] }

//RiskActionPayload is sent when setting the risk action for a seller
type RiskActionPayload struct {
	CustomerCode *string    `json:"customer"`
	RiskAction   RiskAction `json:"risk_action"`
}

//CustomerService handles the communication with the Customer related
type CustomerService service

type Customer struct {
	Email          *string         `json:"email, omitempty"`
	FirstName      *string         `json:"first_name,omitempty"`
	LastName       *string         `json:"last_name,omitempty"`
	Phone          *string         `json:"phone,omitempty"`
	Integration    *int            `json:"integration,omitempty"`
	Domain         *string         `json:"domain,omitempty"`
	CustomerCode   *string         `json:"customer_code,omitempty"`
	Id             *int            `json:"id,omitempty"`
	CreatedAt      *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time      `json:"updatedAt,omitempty"`
	Metadata       Metadata        `json:"metadata,omitempty"`
	RiskAction     *string         `json:"risk_action,omitempty"`
	Transactions   []Transaction   `json:"transactions,omitempty"`
	Subscriptions  []Subscription  `json:"subscriptions,omitempty"`
	Authorizations []Authorization `json:"authorizations,omitempty"`
}

type CustomerRequest struct {
	Email     *string  `json:"email, omitempty"`
	FirstName *string  `json:"first_name,omitempty"`
	LastName  *string  `json:"last_name,omitempty"`
	Phone     *string  `json:"phone,omitempty"`
	Metadata  Metadata `json:"metadata,omitempty"`
}

// Create returns a new customer
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-customer
func (s *CustomerService) Create(ctx context.Context, cr *CustomerRequest) (*Customer, *Response, error) {
	u := fmt.Sprintf("customer")
	req, err := s.client.NewRequest("POST", u, cr)
	if err != nil {
		//return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	c := new(Customer)
	MapDecoder(r.Data, c)
	return c, resp, nil
}

// List returns an array all customers
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-customers
func (s *CustomerService) List(ctx context.Context, opt *ListOptions) ([]Customer, *Response, error) {
	u := fmt.Sprintf("customer")
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
	var ca []Customer
	c := new(Customer)
	for _, x := range lr.Data {
		MapDecoder(x, c)
		ca = append(ca, *c)
	}
	return ca, resp, nil
}

// Fetch returns a new customer with the id
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-customer
func (s *CustomerService) Fetch(ctx context.Context, id string) (*Customer, *Response, error) {
	u := fmt.Sprintf("customer/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var c Customer
	err = MapDecoder(r.Data, &c)
	if err != nil {
		return nil, resp, err
	}
	return &c, resp, nil
}

// Update updates a customer model
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-customer
func (s *CustomerService) Update(ctx context.Context, cr *CustomerRequest, customerCode string) (*Customer, *Response, error) {
	u := fmt.Sprintf("customer/" + customerCode)
	req, err := s.client.NewRequest("PUT", u, cr)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var c Customer
	MapDecoder(&r.Data, c)
	return &c, resp, nil
}

//SetRiskAction takes a risk action on a customer
//A RiskActionPayload has to be supplied
//
// Paystack API reference:
// https://developers.paystack.co/reference#whiteblacklist-customer
func (s *CustomerService) SetRiskAction(ctx context.Context, rap *RiskActionPayload) (*Customer, *Response, error) {
	if rap.GetCustomerCode() == "" {
		return nil, nil, errors.New("RiskActionPayload must have both a riskAction and customer code")
	}
	u := fmt.Sprintf("customer/set_risk_action")
	req, err := s.client.NewRequest("POST", u, rap)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var c Customer
	MapDecoder(r.Data, c)
	return &c, resp, nil
}

//DeactivateAuthorization forgets a customer's card
//the card authentication code is supplied
// Paystack API reference:
// https://developers.paystack.co/reference#deactivate-authorization
func (s *CustomerService) DeactivateAuthorization(ctx context.Context, a *Authorization) (*Customer, *Response, error) {

	u := fmt.Sprintf("customer/deactivate_authorization")
	req, err := s.client.NewRequest("POST", u, a)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	var c Customer
	MapDecoder(r.Data, c)
	return &c, resp, nil
}
