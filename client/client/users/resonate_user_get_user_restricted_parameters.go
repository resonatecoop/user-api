// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewResonateUserGetUserRestrictedParams creates a new ResonateUserGetUserRestrictedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResonateUserGetUserRestrictedParams() *ResonateUserGetUserRestrictedParams {
	return &ResonateUserGetUserRestrictedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResonateUserGetUserRestrictedParamsWithTimeout creates a new ResonateUserGetUserRestrictedParams object
// with the ability to set a timeout on a request.
func NewResonateUserGetUserRestrictedParamsWithTimeout(timeout time.Duration) *ResonateUserGetUserRestrictedParams {
	return &ResonateUserGetUserRestrictedParams{
		timeout: timeout,
	}
}

// NewResonateUserGetUserRestrictedParamsWithContext creates a new ResonateUserGetUserRestrictedParams object
// with the ability to set a context for a request.
func NewResonateUserGetUserRestrictedParamsWithContext(ctx context.Context) *ResonateUserGetUserRestrictedParams {
	return &ResonateUserGetUserRestrictedParams{
		Context: ctx,
	}
}

// NewResonateUserGetUserRestrictedParamsWithHTTPClient creates a new ResonateUserGetUserRestrictedParams object
// with the ability to set a custom HTTPClient for a request.
func NewResonateUserGetUserRestrictedParamsWithHTTPClient(client *http.Client) *ResonateUserGetUserRestrictedParams {
	return &ResonateUserGetUserRestrictedParams{
		HTTPClient: client,
	}
}

/* ResonateUserGetUserRestrictedParams contains all the parameters to send to the API endpoint
   for the resonate user get user restricted operation.

   Typically these are written to a http.Request.
*/
type ResonateUserGetUserRestrictedParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resonate user get user restricted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserGetUserRestrictedParams) WithDefaults() *ResonateUserGetUserRestrictedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resonate user get user restricted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserGetUserRestrictedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) WithTimeout(timeout time.Duration) *ResonateUserGetUserRestrictedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) WithContext(ctx context.Context) *ResonateUserGetUserRestrictedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) WithHTTPClient(client *http.Client) *ResonateUserGetUserRestrictedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) WithID(id string) *ResonateUserGetUserRestrictedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resonate user get user restricted params
func (o *ResonateUserGetUserRestrictedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResonateUserGetUserRestrictedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}