// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewGetSystemScanAllScheduleParams creates a new GetSystemScanAllScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemScanAllScheduleParams() *GetSystemScanAllScheduleParams {
	return &GetSystemScanAllScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemScanAllScheduleParamsWithTimeout creates a new GetSystemScanAllScheduleParams object
// with the ability to set a timeout on a request.
func NewGetSystemScanAllScheduleParamsWithTimeout(timeout time.Duration) *GetSystemScanAllScheduleParams {
	return &GetSystemScanAllScheduleParams{
		timeout: timeout,
	}
}

// NewGetSystemScanAllScheduleParamsWithContext creates a new GetSystemScanAllScheduleParams object
// with the ability to set a context for a request.
func NewGetSystemScanAllScheduleParamsWithContext(ctx context.Context) *GetSystemScanAllScheduleParams {
	return &GetSystemScanAllScheduleParams{
		Context: ctx,
	}
}

// NewGetSystemScanAllScheduleParamsWithHTTPClient creates a new GetSystemScanAllScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemScanAllScheduleParamsWithHTTPClient(client *http.Client) *GetSystemScanAllScheduleParams {
	return &GetSystemScanAllScheduleParams{
		HTTPClient: client,
	}
}

/* GetSystemScanAllScheduleParams contains all the parameters to send to the API endpoint
   for the get system scan all schedule operation.

   Typically these are written to a http.Request.
*/
type GetSystemScanAllScheduleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system scan all schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemScanAllScheduleParams) WithDefaults() *GetSystemScanAllScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system scan all schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemScanAllScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) WithTimeout(timeout time.Duration) *GetSystemScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) WithContext(ctx context.Context) *GetSystemScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) WithHTTPClient(client *http.Client) *GetSystemScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system scan all schedule params
func (o *GetSystemScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
