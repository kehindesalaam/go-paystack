package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kehindesalaam/mapstructure"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	libraryVersion = "1.0"
	defaultBaseURL = "https://api.paystack.co/"
	userAgent      = "go-paystack/" + libraryVersion
)

//A Client manages communication with the Paystack API
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.
	// Base URL for API requests. Defaults to the  Paystack API url.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	//User agent for communicating with the Paystack API
	UserAgent string

	Secret string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Paystack API.
	Balance           *BalanceService
	BulkCharge        *BulkChargeService
	Charge            *ChargeService
	Customer          *CustomerService
	Integration       *IntegrationService
	Miscellaneous     *MiscellaneousService
	Page              *PageService
	Plan              *PlanService
	Settlement        *SettlementService
	Subaccount        *SubaccountService
	Subscription      *SubscriptionService
	Transaction       *TransactionService
	Transfer          *TransferService
	TransferRecipient *TransferRecipientService
}

type service struct {
	client *Client
}

//The standard response type when the Paystack API returns a single data object
type StandardResponse struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	Meta    Meta                   `json:"meta"`
}

//Object to model the meta json object returned by the Paystack API
type MetaResponse struct {
	Meta Meta `json:"meta"`
}

//The standard response type when the Paystack API returns a list of data object
type StandardListResponse struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
	Meta    Meta                     `json:"meta"`
}

type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"` //awkward
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type Message struct {
	Status  *bool   `json:"status, omitempty"`
	Message *string `json:"message, omitempty"`
}

// ListOptions specifies the optional parameters to various List methods
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"perPage,omitempty"`
}

type BullkChargeOptions struct {
	ListOptions
	Status string `json:"status, omitempty"`
}

type IntegrationOptions struct {
	Timeout *int `json:"timeout, omitempty"`
}

type TransactionOptions struct {
	ListOptions
	Customer    int32     `json:"customer, omitempty"`
	Status      string    `json:"status, omitempty"`
	From        time.Time `json:"from, omitempty"`
	To          time.Time `json:"to, omitempty"`
	Amount      string    `json:"amount, omitempty"`
	Settled     *bool     `json:"settled, omitempty"`
	PaymentPage *int      `json:"payment_page, omitempty"`
	Currency    *string   `json:"currency, omitempty"`
	Settlement  *int      `json:"settlement, omitempty"`
}

type SettlementOptions struct {
	From       time.Time `json:"from, omitempty"`
	To         time.Time `json:"to, omitempty"`
	Subaccount *string   `json:"subaccount, omitempty"`
}

type PlanOptions struct {
	ListOptions
	Interval *string `json:"interval, omitempty"`
	Amount   string  `json:"amount, omitempty"`
}

type SubscriptionOptions struct {
	ListOptions
	Customer int32  `json:"customer, omitempty"`
	Plan     *int32 `json:"plan, omitempty"`
}

type Log struct {
	TimeSpent      *int32        `json:"time_spent, omitempty"`
	Attempts       *int32        `json:"attempts, omitempty"`
	Authentication interface{}   `json:"authentication, omitempty"`
	Errors         *int32        `json:"errors, omitempty"`
	Success        *bool         `json:"success, omitempty"`
	Mobile         *bool         `json:"mobile, omitempty"`
	Input          []interface{} `json:"input, omitempty"`
	Channel        *string       `json:"channel, omitempty"`
	History        []History     `json:"history, omitempty"`
}

type History struct {
	Type    *string `json:"type, omitempty"`
	Message *string `json:"message, omitempty"`
	Time    *string `json:"time, omitempty"`
}

type Authorization struct {
	AuthorizationCode *string `json:"authorization_code, omitempty"`
	CardType          *string `json:"card_type, omitempty"`
	Last4             *string `json:"last4, omitempty"`
	ExpMonth          *string `json:"exp_month, omitempty"`
	ExpYear           *string `json:"exp_year, omitempty"`
	Bin               *string `json:"bin, omitempty"`
	Bank              *string `json:"bank, omitempty"`
	Channel           *string `json:"channel, omitempty"`
	Signature         *string `json:"signature, omitempty"`
	Reusable          *bool   `json:"reusable, omitempty"`
	CountryCode       *string `json:"country_code, omitempty"`
	Customer          *string `json:"customer, omitempty"`
}

type Photo struct {
	Type      string `json:"type, omitempty"`
	TypeId    string `json:"typeId, omitempty"`
	TypeName  string `json:"typeName, omitempty"`
	URL       string `json:"url, omitempty"`
	IsPrimary bool   `json:"isPrimary, omitempty"`
}

type FieldByCurrency struct {
	Currency *string `json:"currency, omitempty"`
	Amount   *string `json:"amount, omitempty"`
}

type Metadata struct {
	CustomFields []map[string]interface{} `json:"custom_fields, omitempty"`
	Photos       []Photo                  `json:"photos,omitempty"`
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// SecretKey is a referential function that sets the users secret
// key when initializing the client
func SecretKey(key string) func(*Client) {
	return func(c *Client) {
		c.setSecret(key)
	}
}

// Sets the client's authorization secret
func (c *Client) setSecret(secret string) {
	c.Secret = secret
}

// NewClient returns a new Paystack API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, options ...func(*Client)) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.common.client = c
	c.Balance = (*BalanceService)(&c.common)
	c.BulkCharge = (*BulkChargeService)(&c.common)
	c.Charge = (*ChargeService)(&c.common)
	c.Customer = (*CustomerService)(&c.common)
	c.Integration = (*IntegrationService)(&c.common)
	c.Miscellaneous = (*MiscellaneousService)(&c.common)
	c.Page = (*PageService)(&c.common)
	c.Plan = (*PlanService)(&c.common)
	c.Settlement = (*SettlementService)(&c.common)
	c.Subaccount = (*SubaccountService)(&c.common)
	c.Subscription = (*SubscriptionService)(&c.common)
	c.Transaction = (*TransactionService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)
	c.TransferRecipient = (*TransferRecipientService)(&c.common)

	for _, option := range options {
		option(c)
	}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.Secret != "" {
		//user's application secret
		req.Header.Set("Authorization", "Bearer "+c.Secret)
	}
	return req, nil
}

// Response is a Paystack API response. This wraps the standard http.Response
// returned from Paystack and provides convenient access to things like
// pagination .
type Response struct {
	*http.Response

	// These fields provide the page values for paginating through a set of
	// results. Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.

	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	err := response.populatePageValues()
	if err != nil {
		println(err.Error())
	}
	return response
}

// populatePageValues parses the JSON meta response from Paystack and populates the
// various pagination values in the Response.
func (r *Response) populatePageValues() error {
	//drain bytes from body into t
	t, _ := ioutil.ReadAll(r.Body)
	//buffer to decode the response
	var b MetaResponse
	err := json.NewDecoder(bytes.NewBuffer(t)).Decode(&b)
	if err != nil {
		return err
	}
	//restore r.Body to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(t))

	meta := b.Meta

	if meta.PageCount > 0 {
		r.FirstPage = 1
		r.LastPage = meta.PageCount
		if meta.Page != meta.PageCount {
			r.NextPage = meta.Page + 1
		} else {
			r.NextPage = 0
		}
		if meta.Page > 1 {
			r.PrevPage = meta.Page - 1
		} else {
			r.PrevPage = 0
		}
	}
	return nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

/*
An ErrorResponse reports one or more errors caused by an API request.
Paystack docs: https://developers.paystack.co/docs/errors
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// AuthError occurs when the request was not authorized.
// This can be triggered by passing an invalid secret key
// in the authorization header or the lack of one.
type AuthError ErrorResponse

func (r *AuthError) Error() string { return (*ErrorResponse)(r).Error() }

//BadRequestError occurs when a validation or client side error occurred and
// the request was not fulfilled.
type BadRequestError ErrorResponse

func (r *BadRequestError) Error() string { return (*ErrorResponse)(r).Error() }

// NotFoundError occurs when the request could not be fulfilled as the
// request resource does not exist.
type NotFoundError ErrorResponse

func (r *NotFoundError) Error() string { return (*ErrorResponse)(r).Error() }

// ServerError Occurs when the request could not be fulfilled due to an error on
// Paystack's end. This shouldn't happen so please report to paystack as soon as you
// encounter any instance of this.
type ServerError ErrorResponse

func (r *ServerError) Error() string { return (*ErrorResponse)(r).Error() }

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
//
// The error types are listed as follows,
// *ServerError for status code between 500 and 504 inclusive,
// *BadRequestError for 400 status codes
// *NotFoundError for 404 status codes
// and *AuthError for authentication errors.
func CheckResponse(r *http.Response) error {

	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	if c := r.StatusCode; 500 <= c && c <= 504 {
		return (*ServerError)(errorResponse)
	}

	switch {
	case r.StatusCode == http.StatusBadRequest:
		return (*BadRequestError)(errorResponse)

	case r.StatusCode == http.StatusNotFound:
		return (*NotFoundError)(errorResponse)

	case r.StatusCode == http.StatusUnauthorized:
		return (*AuthError)(errorResponse)

	default:
		return errorResponse
	}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// MapDecoder is a helper function that decodes map[string]interface{}
// objects into another interface
func MapDecoder(source interface{}, target interface{}) error {
	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(time.RFC3339, data.(string))
		}

		return data, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Result:     target,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}

	err = decoder.Decode(source)
	if err != nil {
		return err
	}
	return nil
}
