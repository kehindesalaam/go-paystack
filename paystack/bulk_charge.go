package paystack

import (
	"context"
	"fmt"
	"time"
)

type BulkChargeService service

type BulkBatchRequest struct {
	Authorization *string `json:"authorization, omitempty"`
	Amount        *int    `json:"amount, omitempty"`
}

type BulkBatch struct {
	Domain      *string    `json:"domain, omitempty"`
	BatchCode   *string    `json:"batch_code, omitempty"`
	Status      *string    `json:"status, omitempty"`
	Id          *int       `json:"id, omitempty"`
	Integration *int       `json:"integration, omitempty"`
	CreatedAt   *time.Time `json:"createdAt, omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt, omitempty"`
	TotalCharges *int      `json:"total_charges, omitempty"`
	PendingCharges *int    `json:"pending_charges, omitempty"`
}

type BulkCharge struct {
	Integration   *int          `json:"integration, omitempty"`
	Bulkcharge    *int          `json:"bulkcharge, omitempty"`
	Customer      Customer      `json:"customer, omitempty"`
	Authorization Authorization `json:"authorization, omitempty"`
	Transaction   Transaction   `json:"transaction, omitempty"`
	Domain        *string       `json:"domain, omitempty"`
	Amount        *int          `json:"amount, omitempty"`
	Currency      *string       `json:"currency, omitempty"`
	Status        *string       `json:"status, omitempty"`
	Id            *int          `json:"id, omitempty"`
	CreatedAt     *time.Time    `json:"created_at, omitempty"`
	UpdatedAt     *time.Time    `json:"updated_at, omitempty"`
}

//Initiate
//
// Paystack API reference:
// https://developers.paystack.co/reference#initiate-bulk-charge
func (s *BulkChargeService) Initiate(ctx context.Context, request []*BulkBatchRequest) (*BulkBatch, *Response, error) {
	u := fmt.Sprintf("bulkcharge")
	req, err := s.client.NewRequest("POST", u, request)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	bb := new(BulkBatch)
	MapDecoder(r.Data, bb)
	return bb, resp, nil
}

//ListBatches
//
// Paystack API reference:
// https://developers.paystack.co/reference#list-transactions
func (s *BulkChargeService) ListBatches(ctx context.Context, opt *ListOptions) ([]BulkBatch, *Response, error) {
	u := fmt.Sprintf("bulkcharge")
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
	var bba []BulkBatch
	bb := new(BulkBatch)
	for _, x := range lr.Data {
		MapDecoder(x, bb)
		bba = append(bba, *bb)
	}
	return bba, resp, nil

}

//FetchBatch
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-bulk-charge-batch
func (s *BulkChargeService) FetchBatch(ctx context.Context, id string) (*BulkBatch, *Response, error) {
	u := fmt.Sprintf("bulkcharge/" + id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	r := new(StandardResponse)
	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}
	m := new(BulkBatch)
	MapDecoder(r.Data, m)
	return m, resp, nil
}

//FetchBatchCharges
//
// Paystack API reference:
// https://developers.paystack.co/reference#fetch-transaction
func (s *BulkChargeService) FetchBatchCharges(ctx context.Context, id string, opt *BullkChargeOptions) ([]*BulkCharge, *Response, error) {
	u := fmt.Sprintf("bulkcharge/" + id + "charges")
	//Response is erroneous if opt.Page or opt.PerPage = 0

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	lr := new(StandardListResponse)
	resp, err := s.client.Do(ctx, req, lr)

	if err != nil {
		return nil, resp, err
	}
	var ta []*BulkCharge
	c := new(BulkCharge)
	for _, x := range lr.Data {
		MapDecoder(x, c)
		ta = append(ta, c)
	}
	return ta, resp, nil
}

//PauseBatch
//
// Paystack API reference:
// https://developers.paystack.co/reference#pause-bulk-charge-batch
func (s *BulkChargeService) PauseBatch(ctx context.Context, batch_code string) (*Message, *Response, error) {
	u := fmt.Sprintf("bulkcharge/pause" + batch_code)
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

//ResumeBatch
//
// Paystack API reference:
// https://developers.paystack.co/reference#resume-bulk-charge-batch
func (s *BulkChargeService) ResumeBatch(ctx context.Context, batch_code string) (*Message, *Response, error) {
	u := fmt.Sprintf("bulkcharge/resume" + batch_code)
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
