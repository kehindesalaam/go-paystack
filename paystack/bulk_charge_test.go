package paystack

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
	"time"
)

func TestBulkChargeService_Initiate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bulkcharge", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Charges have been queued",
		  "data": {
			"domain": "test",
			"status": "active",
			"id": 17,
			"integration": 100073,
			"batch_code": "BCH_180tl7oq7cayggh",
			"createdAt": "2017-02-04T05:44:19.000Z",
			"updatedAt": "2017-02-04T05:44:19.000Z"
		  }
		}`)
	})
	bbr := []*BulkBatchRequest{}
	bbr = append(bbr, &BulkBatchRequest{Authorization: String("AUTH_n95vpedf"), Amount: Int(2500)})
	bbr = append(bbr, &BulkBatchRequest{Authorization: String("AUTH_ljdt4e4j"), Amount: Int(1500)})

	bc, _, err := client.BulkCharge.Initiate(context.Background(), bbr)
	if err != nil {
		t.Errorf("BulkCharge.Initiate returned error: %v", err)
	}
	createdAt := time.Date(2017, 02, 04, 05, 44, 19, 0, time.UTC)
	updatedAt := time.Date(2017, 02, 04, 05, 44, 19, 0, time.UTC)

	want := &BulkBatch{
		Domain:      String("test"),
		BatchCode:   String("BCH_180tl7oq7cayggh"),
		Status:      String("active"),
		Integration: Int(100073),
		Id:          Int(17),
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
	}

	if !cmp.Equal(bc, want) {
		t.Errorf("BulkCharge.Initiate returned %+v, want %+v", bc, want)
	}

}

func TestBulkChargeService_ListBatches(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bulkcharge", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Bulk charges retrieved",
		  "data": [
			{
			  "domain": "test",
			  "batch_code": "BCH_1nV4L1D7cayggh",
			  "status": "complete",
			  "id": 1733,
			  "createdAt": "2017-02-04T05:44:19.000Z",
			  "updatedAt": "2017-02-04T05:45:02.000Z"
			}
		  ],
		  "meta": {
			"total": 1,
			"skipped": 0,
			"perPage": 50,
			"page": 1,
			"pageCount": 1
		  }
		}`)
	})

	bc, _, err := client.BulkCharge.ListBatches(context.Background(), nil)
	if err != nil {
		t.Errorf("BulkCharge.ListBatches returned error: %v", err)
	}

	createdAt := time.Date(2017, 02, 04, 05, 44, 19, 0, time.UTC)
	updatedAt := time.Date(2017, 02, 04, 05, 45, 02, 0, time.UTC)

	var want []BulkBatch
	want = append(want, BulkBatch{
		Domain:    String("test"),
		BatchCode: String("BCH_1nV4L1D7cayggh"),
		Status:    String("complete"),
		Id:        Int(1733),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	})

	if !cmp.Equal(bc, want) {
		t.Errorf("BulkCharge.ListBatches returned %+v, want %+v", bc, want)
	}
}

func TestBulkChargeService_FetchBatch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bulkcharge/BCH_180tl7oq7cayggh", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Bulk charge retrieved",
		  "data": {
			"domain": "test",
			"batch_code": "BCH_180tl7oq7cayggh",
			"status": "complete",
			"id": 17,
			"total_charges": 0,
			"pending_charges": 0,
			"createdAt": "2017-02-04T05:44:19.000Z",
			"updatedAt": "2017-02-04T05:45:02.000Z"
		  }
		}`)
	})

	bc, _, err := client.BulkCharge.FetchBatch(context.Background(), "BCH_180tl7oq7cayggh")
	if err != nil {
		t.Errorf("BulkCharge.FetchBatch returned error: %v", err)
	}

	createdAt := time.Date(2017, 02, 04, 05, 44, 19, 0, time.UTC)
	updatedAt := time.Date(2017, 02, 04, 05, 45, 02, 0, time.UTC)

	want := &BulkBatch{
		Domain:         String("test"),
		BatchCode:      String("BCH_180tl7oq7cayggh"),
		Status:         String("complete"),
		Id:             Int(17),
		TotalCharges:   Int(0),
		PendingCharges: Int(0),
		CreatedAt:      &createdAt,
		UpdatedAt:      &updatedAt,
	}

	if !cmp.Equal(bc, want) {
		t.Errorf("BulkCharge.FetchBatch returned %+v, want %+v", bc, want)
	}
}

func TestBulkChargeService_FetchBatchCharges(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/bulkcharge/BCH_180tl7oq7cayggh", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
		  "status": true,
		  "message": "Bulk charge retrieved",
		  "data": {
			"domain": "test",
			"batch_code": "BCH_180tl7oq7cayggh",
			"status": "complete",
			"id": 17,
			"total_charges": 0,
			"pending_charges": 0,
			"createdAt": "2017-02-04T05:44:19.000Z",
			"updatedAt": "2017-02-04T05:45:02.000Z"
		  }
		}`)
	})

	bc, _, err := client.BulkCharge.FetchBatch(context.Background(), "BCH_180tl7oq7cayggh")
	if err != nil {
		t.Errorf("BulkCharge.FetchBatch returned error: %v", err)
	}

	createdAt := time.Date(2017, 02, 04, 05, 44, 19, 0, time.UTC)
	updatedAt := time.Date(2017, 02, 04, 05, 45, 02, 0, time.UTC)

	want := &BulkBatch{
		Domain:         String("test"),
		BatchCode:      String("BCH_180tl7oq7cayggh"),
		Status:         String("complete"),
		Id:             Int(17),
		TotalCharges:   Int(0),
		PendingCharges: Int(0),
		CreatedAt:      &createdAt,
		UpdatedAt:      &updatedAt,
	}

	if !cmp.Equal(bc, want) {
		t.Errorf("BulkCharge.FetchBatch returned %+v, want %+v", bc, want)
	}
}
