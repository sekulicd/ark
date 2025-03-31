// Code generated by go-swagger; DO NOT EDIT.

package explorer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExplorerServiceSubscribeForAddressParams creates a new ExplorerServiceSubscribeForAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExplorerServiceSubscribeForAddressParams() *ExplorerServiceSubscribeForAddressParams {
	return &ExplorerServiceSubscribeForAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExplorerServiceSubscribeForAddressParamsWithTimeout creates a new ExplorerServiceSubscribeForAddressParams object
// with the ability to set a timeout on a request.
func NewExplorerServiceSubscribeForAddressParamsWithTimeout(timeout time.Duration) *ExplorerServiceSubscribeForAddressParams {
	return &ExplorerServiceSubscribeForAddressParams{
		timeout: timeout,
	}
}

// NewExplorerServiceSubscribeForAddressParamsWithContext creates a new ExplorerServiceSubscribeForAddressParams object
// with the ability to set a context for a request.
func NewExplorerServiceSubscribeForAddressParamsWithContext(ctx context.Context) *ExplorerServiceSubscribeForAddressParams {
	return &ExplorerServiceSubscribeForAddressParams{
		Context: ctx,
	}
}

// NewExplorerServiceSubscribeForAddressParamsWithHTTPClient creates a new ExplorerServiceSubscribeForAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewExplorerServiceSubscribeForAddressParamsWithHTTPClient(client *http.Client) *ExplorerServiceSubscribeForAddressParams {
	return &ExplorerServiceSubscribeForAddressParams{
		HTTPClient: client,
	}
}

/*
ExplorerServiceSubscribeForAddressParams contains all the parameters to send to the API endpoint

	for the explorer service subscribe for address operation.

	Typically these are written to a http.Request.
*/
type ExplorerServiceSubscribeForAddressParams struct {

	// Address.
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the explorer service subscribe for address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExplorerServiceSubscribeForAddressParams) WithDefaults() *ExplorerServiceSubscribeForAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the explorer service subscribe for address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExplorerServiceSubscribeForAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) WithTimeout(timeout time.Duration) *ExplorerServiceSubscribeForAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) WithContext(ctx context.Context) *ExplorerServiceSubscribeForAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) WithHTTPClient(client *http.Client) *ExplorerServiceSubscribeForAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) WithAddress(address string) *ExplorerServiceSubscribeForAddressParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the explorer service subscribe for address params
func (o *ExplorerServiceSubscribeForAddressParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *ExplorerServiceSubscribeForAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
