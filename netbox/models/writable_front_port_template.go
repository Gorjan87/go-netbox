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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableFrontPortTemplate writable front port template
// swagger:model WritableFrontPortTemplate
type WritableFrontPortTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Rear port
	// Required: true
	RearPort *int64 `json:"rear_port"`

	// Rear port position
	// Maximum: 64
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// Type
	// Required: true
	// Enum: [8p8c 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st]
	Type *string `json:"type"`
}

// Validate validates this writable front port template
func (m *WritableFrontPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableFrontPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateRearPortPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", int64(m.RearPortPosition), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", int64(m.RearPortPosition), 64, false); err != nil {
		return err
	}

	return nil
}

var writableFrontPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableFrontPortTemplateTypeTypePropEnum = append(writableFrontPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableFrontPortTemplateTypeNr8p8c captures enum value "8p8c"
	WritableFrontPortTemplateTypeNr8p8c string = "8p8c"

	// WritableFrontPortTemplateTypeNr110Punch captures enum value "110-punch"
	WritableFrontPortTemplateTypeNr110Punch string = "110-punch"

	// WritableFrontPortTemplateTypeBnc captures enum value "bnc"
	WritableFrontPortTemplateTypeBnc string = "bnc"

	// WritableFrontPortTemplateTypeMrj21 captures enum value "mrj21"
	WritableFrontPortTemplateTypeMrj21 string = "mrj21"

	// WritableFrontPortTemplateTypeFc captures enum value "fc"
	WritableFrontPortTemplateTypeFc string = "fc"

	// WritableFrontPortTemplateTypeLc captures enum value "lc"
	WritableFrontPortTemplateTypeLc string = "lc"

	// WritableFrontPortTemplateTypeLcApc captures enum value "lc-apc"
	WritableFrontPortTemplateTypeLcApc string = "lc-apc"

	// WritableFrontPortTemplateTypeLsh captures enum value "lsh"
	WritableFrontPortTemplateTypeLsh string = "lsh"

	// WritableFrontPortTemplateTypeLshApc captures enum value "lsh-apc"
	WritableFrontPortTemplateTypeLshApc string = "lsh-apc"

	// WritableFrontPortTemplateTypeMpo captures enum value "mpo"
	WritableFrontPortTemplateTypeMpo string = "mpo"

	// WritableFrontPortTemplateTypeMtrj captures enum value "mtrj"
	WritableFrontPortTemplateTypeMtrj string = "mtrj"

	// WritableFrontPortTemplateTypeSc captures enum value "sc"
	WritableFrontPortTemplateTypeSc string = "sc"

	// WritableFrontPortTemplateTypeScApc captures enum value "sc-apc"
	WritableFrontPortTemplateTypeScApc string = "sc-apc"

	// WritableFrontPortTemplateTypeSt captures enum value "st"
	WritableFrontPortTemplateTypeSt string = "st"
)

// prop value enum
func (m *WritableFrontPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableFrontPortTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableFrontPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableFrontPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableFrontPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableFrontPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
