// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasTagsListParams creates a new ExtrasTagsListParams object
// with the default values initialized.
func NewExtrasTagsListParams() *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsListParamsWithTimeout creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTagsListParamsWithTimeout(timeout time.Duration) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		timeout: timeout,
	}
}

// NewExtrasTagsListParamsWithContext creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTagsListParamsWithContext(ctx context.Context) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		Context: ctx,
	}
}

// NewExtrasTagsListParamsWithHTTPClient creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTagsListParamsWithHTTPClient(client *http.Client) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{
		HTTPClient: client,
	}
}

/*ExtrasTagsListParams contains all the parameters to send to the API endpoint
for the extras tags list operation typically these are written to a http.Request
*/
type ExtrasTagsListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string
	/*SlugIc*/
	SlugIc *string
	/*SlugIe*/
	SlugIe *string
	/*SlugIew*/
	SlugIew *string
	/*SlugIsw*/
	SlugIsw *string
	/*Slugn*/
	Slugn *string
	/*SlugNic*/
	SlugNic *string
	/*SlugNie*/
	SlugNie *string
	/*SlugNiew*/
	SlugNiew *string
	/*SlugNisw*/
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras tags list params
func (o *ExtrasTagsListParams) WithTimeout(timeout time.Duration) *ExtrasTagsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags list params
func (o *ExtrasTagsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags list params
func (o *ExtrasTagsListParams) WithContext(ctx context.Context) *ExtrasTagsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags list params
func (o *ExtrasTagsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags list params
func (o *ExtrasTagsListParams) WithHTTPClient(client *http.Client) *ExtrasTagsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags list params
func (o *ExtrasTagsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the extras tags list params
func (o *ExtrasTagsListParams) WithLimit(limit *int64) *ExtrasTagsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras tags list params
func (o *ExtrasTagsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras tags list params
func (o *ExtrasTagsListParams) WithName(name *string) *ExtrasTagsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras tags list params
func (o *ExtrasTagsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras tags list params
func (o *ExtrasTagsListParams) WithNameIc(nameIc *string) *ExtrasTagsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras tags list params
func (o *ExtrasTagsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras tags list params
func (o *ExtrasTagsListParams) WithNameIe(nameIe *string) *ExtrasTagsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras tags list params
func (o *ExtrasTagsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras tags list params
func (o *ExtrasTagsListParams) WithNameIew(nameIew *string) *ExtrasTagsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras tags list params
func (o *ExtrasTagsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras tags list params
func (o *ExtrasTagsListParams) WithNameIsw(nameIsw *string) *ExtrasTagsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras tags list params
func (o *ExtrasTagsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras tags list params
func (o *ExtrasTagsListParams) WithNamen(namen *string) *ExtrasTagsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras tags list params
func (o *ExtrasTagsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras tags list params
func (o *ExtrasTagsListParams) WithNameNic(nameNic *string) *ExtrasTagsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras tags list params
func (o *ExtrasTagsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras tags list params
func (o *ExtrasTagsListParams) WithNameNie(nameNie *string) *ExtrasTagsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras tags list params
func (o *ExtrasTagsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras tags list params
func (o *ExtrasTagsListParams) WithNameNiew(nameNiew *string) *ExtrasTagsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras tags list params
func (o *ExtrasTagsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras tags list params
func (o *ExtrasTagsListParams) WithNameNisw(nameNisw *string) *ExtrasTagsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras tags list params
func (o *ExtrasTagsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras tags list params
func (o *ExtrasTagsListParams) WithOffset(offset *int64) *ExtrasTagsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras tags list params
func (o *ExtrasTagsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras tags list params
func (o *ExtrasTagsListParams) WithQ(q *string) *ExtrasTagsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras tags list params
func (o *ExtrasTagsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the extras tags list params
func (o *ExtrasTagsListParams) WithSlug(slug *string) *ExtrasTagsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras tags list params
func (o *ExtrasTagsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugIc(slugIc *string) *ExtrasTagsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugIe(slugIe *string) *ExtrasTagsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugIew(slugIew *string) *ExtrasTagsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugIsw(slugIsw *string) *ExtrasTagsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugn(slugn *string) *ExtrasTagsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugNic(slugNic *string) *ExtrasTagsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugNie(slugNie *string) *ExtrasTagsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugNiew(slugNiew *string) *ExtrasTagsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the extras tags list params
func (o *ExtrasTagsListParams) WithSlugNisw(slugNisw *string) *ExtrasTagsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the extras tags list params
func (o *ExtrasTagsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string
		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {
			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}

	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string
		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {
			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}

	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string
		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {
			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}

	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string
		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {
			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}

	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string
		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {
			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}

	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string
		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {
			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}

	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string
		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {
			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}

	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string
		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {
			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}

	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string
		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {
			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
