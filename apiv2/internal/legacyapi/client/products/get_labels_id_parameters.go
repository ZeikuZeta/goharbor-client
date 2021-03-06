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
	"github.com/go-openapi/swag"
)

// NewGetLabelsIDParams creates a new GetLabelsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelsIDParams() *GetLabelsIDParams {
	return &GetLabelsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsIDParamsWithTimeout creates a new GetLabelsIDParams object
// with the ability to set a timeout on a request.
func NewGetLabelsIDParamsWithTimeout(timeout time.Duration) *GetLabelsIDParams {
	return &GetLabelsIDParams{
		timeout: timeout,
	}
}

// NewGetLabelsIDParamsWithContext creates a new GetLabelsIDParams object
// with the ability to set a context for a request.
func NewGetLabelsIDParamsWithContext(ctx context.Context) *GetLabelsIDParams {
	return &GetLabelsIDParams{
		Context: ctx,
	}
}

// NewGetLabelsIDParamsWithHTTPClient creates a new GetLabelsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelsIDParamsWithHTTPClient(client *http.Client) *GetLabelsIDParams {
	return &GetLabelsIDParams{
		HTTPClient: client,
	}
}

/* GetLabelsIDParams contains all the parameters to send to the API endpoint
   for the get labels ID operation.

   Typically these are written to a http.Request.
*/
type GetLabelsIDParams struct {

	/* ID.

	   Label ID

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get labels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsIDParams) WithDefaults() *GetLabelsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get labels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get labels ID params
func (o *GetLabelsIDParams) WithTimeout(timeout time.Duration) *GetLabelsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels ID params
func (o *GetLabelsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels ID params
func (o *GetLabelsIDParams) WithContext(ctx context.Context) *GetLabelsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels ID params
func (o *GetLabelsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels ID params
func (o *GetLabelsIDParams) WithHTTPClient(client *http.Client) *GetLabelsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels ID params
func (o *GetLabelsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get labels ID params
func (o *GetLabelsIDParams) WithID(id int64) *GetLabelsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get labels ID params
func (o *GetLabelsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
