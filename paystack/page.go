package paystack

import (
	"context"
	"fmt"
	"time"
)

//PageService handles the communication with the Pages related parts of the Paystack API
type PageService service

type Page struct {
	Integration  *int         `json:"integration, omitempty, omitempty"`
	Plan         *int         `json:"plan, omitempty"`
	Domain       *string      `json:"domain, omitempty"`
	Name         *string      `json:"name, omitempty"`
	Description  *string      `json:"description, omitempty"`
	Amount       *int         `json:"amount, omitempty"`
	Currency     *string      `json:"currency, omitempty"`
	Slug         *string      `json:"slug, omitempty"`
	CustomFields CustomFields `json:"custom_fields, omitempty"`
	RedirectUrl  *string      `json:"redirect_url, omitempty"`
	Active       *bool        `json:"active, omitempty"`
	Migrate      interface{}  `json:"migrate, omitempty"`
	Id           *int         `json:"id, omitempty"`
	CreatedAt    *time.Time   `json:"created_at, omitempty"`
	UpdatedAt    *time.Time   `json:"updated_at, omitempty"`
}

type PageRequest struct {
	Name         *string      `json:"name, omitempty"`
	Description  *string      `json:"description, omitempty"`
	Amount       *int         `json:"amount, omitempty"`
	Currency     *string      `json:"currency, omitempty"`
	Slug         *string      `json:"slug, omitempty"`
	CustomFields CustomFields `json:"custom_fields, omitempty"`
	RedirectUrl  *string      `json:"redirect_url, omitempty"`
	Active       *bool        `json:"active, omitempty"`
	Id           *int         `json:"id, omitempty"`
}

type CustomFields struct {
	DisplayName  *string `json:"display_name, omitempty"`
	VariableName *string `json:"variable_name, omitempty"`
	Value        *string `json:"value, omitempty"`
}

// Create returns a new page
//
// Paystack API reference:
// https://developers.paystack.co/reference#create-page
func (s *PageService) Create(ctx context.Context, sa *PageRequest) (*Page, *Response, error) {
	u := fmt.Sprintf("page")
	req, err := s.client.NewRequest("POST", u, sa)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	p := new(Page)
	MapDecoder(r.Data, p)
	return p, resp, nil
}

// List returns all created pages
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-pages
func (s *PageService) List(ctx context.Context, opt *ListOptions) ([]Page, *Response, error) {
	u := fmt.Sprintf("page")
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
	var pa []Page
	p := new(Page)
	for _, x := range lr.Data {
		MapDecoder(x, p)
		pa = append(pa, *p)
	}
	return pa, resp, nil
}

// Fetch fetches a page
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-page
func (s *PageService) Fetch(ctx context.Context, id string) (*Page, *Response, error) {
	u := fmt.Sprintf("page/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	c := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}
	p := new(Page)
	MapDecoder(c.Data, p)
	return p, resp, nil
}

// Update updates a page model with the page
//
// Paystack API reference:
// https://developers.paystack.co/reference#update-page
func (s *PageService) Update(ctx context.Context, sa *PageRequest, id string) (*Page, *Response, error) {
	u := fmt.Sprintf("page/" + id)
	req, err := s.client.NewRequest("PUT", u, sa)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	p := new(Page)
	MapDecoder(r.Data, p)
	return p, resp, nil
}

// CheckSlugAvailability checks if a slug is available
//
// Paystack API reference:
// https://developers.paystack.co/reference#check-slug-availability
func (s *PageService) CheckSlugAvailability(ctx context.Context, id string) (*Message, *Response, error) {
	u := fmt.Sprintf("page/check_slug_availability/fol-offerings")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(Message)
	MapDecoder(r.Data, m)
	return m, resp, nil
}
