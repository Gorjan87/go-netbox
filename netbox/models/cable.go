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

// Cable cable
// swagger:model Cable
type Cable struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	// Max Length: 100
	Label string `json:"label,omitempty"`

	// Length
	// Maximum: 32767
	// Minimum: 0
	Length *int64 `json:"length,omitempty"`

	// length unit
	LengthUnit *CableLengthUnit `json:"length_unit,omitempty"`

	// status
	Status *CableStatus `json:"status,omitempty"`

	// Termination a
	// Read Only: true
	Terminationa map[string]string `json:"termination_a,omitempty"`

	// Termination a id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationaID *int64 `json:"termination_a_id"`

	// Termination a type
	// Required: true
	TerminationaType *string `json:"termination_a_type"`

	// Termination b
	// Read Only: true
	Terminationb map[string]string `json:"termination_b,omitempty"`

	// Termination b id
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	TerminationbID *int64 `json:"termination_b_id"`

	// Termination b type
	// Required: true
	TerminationbType *string `json:"termination_b_type"`

	// Type
	// Enum: [cat3 cat5 cat5e cat6 cat6a cat7 dac-active dac-passive mrj21-trunk coaxial mmf mmf-om1 mmf-om2 mmf-om3 mmf-om4 smf smf-os1 smf-os2 aoc power]
	Type string `json:"type,omitempty"`
}

// Validate validates this cable
func (m *Cable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLengthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminationbType(formats); err != nil {
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

func (m *Cable) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", string(m.Color), 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", string(m.Color), `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 100); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLength(formats strfmt.Registry) error {

	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if err := validate.MinimumInt("length", "body", int64(*m.Length), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("length", "body", int64(*m.Length), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateLengthUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.LengthUnit) { // not required
		return nil
	}

	if m.LengthUnit != nil {
		if err := m.LengthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("length_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Cable) validateTerminationaID(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_id", "body", m.TerminationaID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_a_id", "body", int64(*m.TerminationaID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_a_id", "body", int64(*m.TerminationaID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationaType(formats strfmt.Registry) error {

	if err := validate.Required("termination_a_type", "body", m.TerminationaType); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationbID(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_id", "body", m.TerminationbID); err != nil {
		return err
	}

	if err := validate.MinimumInt("termination_b_id", "body", int64(*m.TerminationbID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("termination_b_id", "body", int64(*m.TerminationbID), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *Cable) validateTerminationbType(formats strfmt.Registry) error {

	if err := validate.Required("termination_b_type", "body", m.TerminationbType); err != nil {
		return err
	}

	return nil
}

var cableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cat3","cat5","cat5e","cat6","cat6a","cat7","dac-active","dac-passive","mrj21-trunk","coaxial","mmf","mmf-om1","mmf-om2","mmf-om3","mmf-om4","smf","smf-os1","smf-os2","aoc","power"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableTypeTypePropEnum = append(cableTypeTypePropEnum, v)
	}
}

const (

	// CableTypeCat3 captures enum value "cat3"
	CableTypeCat3 string = "cat3"

	// CableTypeCat5 captures enum value "cat5"
	CableTypeCat5 string = "cat5"

	// CableTypeCat5e captures enum value "cat5e"
	CableTypeCat5e string = "cat5e"

	// CableTypeCat6 captures enum value "cat6"
	CableTypeCat6 string = "cat6"

	// CableTypeCat6a captures enum value "cat6a"
	CableTypeCat6a string = "cat6a"

	// CableTypeCat7 captures enum value "cat7"
	CableTypeCat7 string = "cat7"

	// CableTypeDacActive captures enum value "dac-active"
	CableTypeDacActive string = "dac-active"

	// CableTypeDacPassive captures enum value "dac-passive"
	CableTypeDacPassive string = "dac-passive"

	// CableTypeMrj21Trunk captures enum value "mrj21-trunk"
	CableTypeMrj21Trunk string = "mrj21-trunk"

	// CableTypeCoaxial captures enum value "coaxial"
	CableTypeCoaxial string = "coaxial"

	// CableTypeMmf captures enum value "mmf"
	CableTypeMmf string = "mmf"

	// CableTypeMmfOm1 captures enum value "mmf-om1"
	CableTypeMmfOm1 string = "mmf-om1"

	// CableTypeMmfOm2 captures enum value "mmf-om2"
	CableTypeMmfOm2 string = "mmf-om2"

	// CableTypeMmfOm3 captures enum value "mmf-om3"
	CableTypeMmfOm3 string = "mmf-om3"

	// CableTypeMmfOm4 captures enum value "mmf-om4"
	CableTypeMmfOm4 string = "mmf-om4"

	// CableTypeSmf captures enum value "smf"
	CableTypeSmf string = "smf"

	// CableTypeSmfOs1 captures enum value "smf-os1"
	CableTypeSmfOs1 string = "smf-os1"

	// CableTypeSmfOs2 captures enum value "smf-os2"
	CableTypeSmfOs2 string = "smf-os2"

	// CableTypeAoc captures enum value "aoc"
	CableTypeAoc string = "aoc"

	// CableTypePower captures enum value "power"
	CableTypePower string = "power"
)

// prop value enum
func (m *Cable) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Cable) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cable) UnmarshalBinary(b []byte) error {
	var res Cable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableLengthUnit Length unit
// swagger:model CableLengthUnit
type CableLengthUnit struct {

	// label
	// Required: true
	// Enum: [Meters Centimeters Feet Inches]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [m cm ft in]
	Value *string `json:"value"`
}

// Validate validates this cable length unit
func (m *CableLengthUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cableLengthUnitTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Meters","Centimeters","Feet","Inches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeLabelPropEnum = append(cableLengthUnitTypeLabelPropEnum, v)
	}
}

const (

	// CableLengthUnitLabelMeters captures enum value "Meters"
	CableLengthUnitLabelMeters string = "Meters"

	// CableLengthUnitLabelCentimeters captures enum value "Centimeters"
	CableLengthUnitLabelCentimeters string = "Centimeters"

	// CableLengthUnitLabelFeet captures enum value "Feet"
	CableLengthUnitLabelFeet string = "Feet"

	// CableLengthUnitLabelInches captures enum value "Inches"
	CableLengthUnitLabelInches string = "Inches"
)

// prop value enum
func (m *CableLengthUnit) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableLengthUnitTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("length_unit"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableLengthUnitTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["m","cm","ft","in"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableLengthUnitTypeValuePropEnum = append(cableLengthUnitTypeValuePropEnum, v)
	}
}

const (

	// CableLengthUnitValueM captures enum value "m"
	CableLengthUnitValueM string = "m"

	// CableLengthUnitValueCm captures enum value "cm"
	CableLengthUnitValueCm string = "cm"

	// CableLengthUnitValueFt captures enum value "ft"
	CableLengthUnitValueFt string = "ft"

	// CableLengthUnitValueIn captures enum value "in"
	CableLengthUnitValueIn string = "in"
)

// prop value enum
func (m *CableLengthUnit) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableLengthUnitTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableLengthUnit) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("length_unit"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("length_unit"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableLengthUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableLengthUnit) UnmarshalBinary(b []byte) error {
	var res CableLengthUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CableStatus Status
// swagger:model CableStatus
type CableStatus struct {

	// label
	// Required: true
	// Enum: [Connected Planned Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [connected planned decommissioning]
	Value *string `json:"value"`
}

// Validate validates this cable status
func (m *CableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cableStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Connected","Planned","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeLabelPropEnum = append(cableStatusTypeLabelPropEnum, v)
	}
}

const (

	// CableStatusLabelConnected captures enum value "Connected"
	CableStatusLabelConnected string = "Connected"

	// CableStatusLabelPlanned captures enum value "Planned"
	CableStatusLabelPlanned string = "Planned"

	// CableStatusLabelDecommissioning captures enum value "Decommissioning"
	CableStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *CableStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeValuePropEnum = append(cableStatusTypeValuePropEnum, v)
	}
}

const (

	// CableStatusValueConnected captures enum value "connected"
	CableStatusValueConnected string = "connected"

	// CableStatusValuePlanned captures enum value "planned"
	CableStatusValuePlanned string = "planned"

	// CableStatusValueDecommissioning captures enum value "decommissioning"
	CableStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *CableStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableStatus) UnmarshalBinary(b []byte) error {
	var res CableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
