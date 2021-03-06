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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerPort power port
// swagger:model PowerPort
type PowerPort struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// connection status
	ConnectionStatus *PowerPortConnectionStatus `json:"connection_status,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Maximum draw
	//
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// tags
	Tags []string `json:"tags"`

	// type
	Type *PowerPortType `json:"type,omitempty"`
}

// Validate validates this power port
func (m *PowerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *PowerPort) validateAllocatedDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *PowerPort) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *PowerPort) validateMaximumDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", int64(*m.MaximumDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", int64(*m.MaximumDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *PowerPort) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *PowerPort) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPort) UnmarshalBinary(b []byte) error {
	var res PowerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerPortConnectionStatus Connection status
// swagger:model PowerPortConnectionStatus
type PowerPortConnectionStatus struct {

	// label
	// Required: true
	// Enum: [Not Connected Connected]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [false true]
	Value *bool `json:"value"`
}

// Validate validates this power port connection status
func (m *PowerPortConnectionStatus) Validate(formats strfmt.Registry) error {
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

var powerPortConnectionStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Connected","Connected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortConnectionStatusTypeLabelPropEnum = append(powerPortConnectionStatusTypeLabelPropEnum, v)
	}
}

const (

	// PowerPortConnectionStatusLabelNotConnected captures enum value "Not Connected"
	PowerPortConnectionStatusLabelNotConnected string = "Not Connected"

	// PowerPortConnectionStatusLabelConnected captures enum value "Connected"
	PowerPortConnectionStatusLabelConnected string = "Connected"
)

// prop value enum
func (m *PowerPortConnectionStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerPortConnectionStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("connection_status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerPortConnectionStatusTypeValuePropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortConnectionStatusTypeValuePropEnum = append(powerPortConnectionStatusTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *PowerPortConnectionStatus) validateValueEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, powerPortConnectionStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("connection_status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortConnectionStatus) UnmarshalBinary(b []byte) error {
	var res PowerPortConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerPortType Type
// swagger:model PowerPortType
type PowerPortType struct {

	// label
	// Required: true
	// Enum: [C6 C8 C14 C16 C20 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 5-15P NEMA 5-20P NEMA 5-30P NEMA 5-50P NEMA 6-15P NEMA 6-20P NEMA 6-30P NEMA 6-50P NEMA L5-15P NEMA L5-20P NEMA L5-30P NEMA L6-15P NEMA L6-20P NEMA L6-30P NEMA L6-50P CS6361C CS6365C CS8165C CS8265C CS8365C CS8465C ITA Type E (CEE 7/5) ITA Type F (CEE 7/4) ITA Type E/F (CEE 7/7) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-20p nema-l6-30p nema-l6-50p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o]
	Value *string `json:"value"`
}

// Validate validates this power port type
func (m *PowerPortType) Validate(formats strfmt.Registry) error {
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

var powerPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C6","C8","C14","C16","C20","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 5-15P","NEMA 5-20P","NEMA 5-30P","NEMA 5-50P","NEMA 6-15P","NEMA 6-20P","NEMA 6-30P","NEMA 6-50P","NEMA L5-15P","NEMA L5-20P","NEMA L5-30P","NEMA L6-15P","NEMA L6-20P","NEMA L6-30P","NEMA L6-50P","CS6361C","CS6365C","CS8165C","CS8265C","CS8365C","CS8465C","ITA Type E (CEE 7/5)","ITA Type F (CEE 7/4)","ITA Type E/F (CEE 7/7)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTypeTypeLabelPropEnum = append(powerPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerPortTypeLabelC6 captures enum value "C6"
	PowerPortTypeLabelC6 string = "C6"

	// PowerPortTypeLabelC8 captures enum value "C8"
	PowerPortTypeLabelC8 string = "C8"

	// PowerPortTypeLabelC14 captures enum value "C14"
	PowerPortTypeLabelC14 string = "C14"

	// PowerPortTypeLabelC16 captures enum value "C16"
	PowerPortTypeLabelC16 string = "C16"

	// PowerPortTypeLabelC20 captures enum value "C20"
	PowerPortTypeLabelC20 string = "C20"

	// PowerPortTypeLabelPNE4H captures enum value "P+N+E 4H"
	PowerPortTypeLabelPNE4H string = "P+N+E 4H"

	// PowerPortTypeLabelPNE6H captures enum value "P+N+E 6H"
	PowerPortTypeLabelPNE6H string = "P+N+E 6H"

	// PowerPortTypeLabelPNE9H captures enum value "P+N+E 9H"
	PowerPortTypeLabelPNE9H string = "P+N+E 9H"

	// PowerPortTypeLabelNr2PE4H captures enum value "2P+E 4H"
	PowerPortTypeLabelNr2PE4H string = "2P+E 4H"

	// PowerPortTypeLabelNr2PE6H captures enum value "2P+E 6H"
	PowerPortTypeLabelNr2PE6H string = "2P+E 6H"

	// PowerPortTypeLabelNr2PE9H captures enum value "2P+E 9H"
	PowerPortTypeLabelNr2PE9H string = "2P+E 9H"

	// PowerPortTypeLabelNr3PE4H captures enum value "3P+E 4H"
	PowerPortTypeLabelNr3PE4H string = "3P+E 4H"

	// PowerPortTypeLabelNr3PE6H captures enum value "3P+E 6H"
	PowerPortTypeLabelNr3PE6H string = "3P+E 6H"

	// PowerPortTypeLabelNr3PE9H captures enum value "3P+E 9H"
	PowerPortTypeLabelNr3PE9H string = "3P+E 9H"

	// PowerPortTypeLabelNr3PNE4H captures enum value "3P+N+E 4H"
	PowerPortTypeLabelNr3PNE4H string = "3P+N+E 4H"

	// PowerPortTypeLabelNr3PNE6H captures enum value "3P+N+E 6H"
	PowerPortTypeLabelNr3PNE6H string = "3P+N+E 6H"

	// PowerPortTypeLabelNr3PNE9H captures enum value "3P+N+E 9H"
	PowerPortTypeLabelNr3PNE9H string = "3P+N+E 9H"

	// PowerPortTypeLabelNEMA515P captures enum value "NEMA 5-15P"
	PowerPortTypeLabelNEMA515P string = "NEMA 5-15P"

	// PowerPortTypeLabelNEMA520P captures enum value "NEMA 5-20P"
	PowerPortTypeLabelNEMA520P string = "NEMA 5-20P"

	// PowerPortTypeLabelNEMA530P captures enum value "NEMA 5-30P"
	PowerPortTypeLabelNEMA530P string = "NEMA 5-30P"

	// PowerPortTypeLabelNEMA550P captures enum value "NEMA 5-50P"
	PowerPortTypeLabelNEMA550P string = "NEMA 5-50P"

	// PowerPortTypeLabelNEMA615P captures enum value "NEMA 6-15P"
	PowerPortTypeLabelNEMA615P string = "NEMA 6-15P"

	// PowerPortTypeLabelNEMA620P captures enum value "NEMA 6-20P"
	PowerPortTypeLabelNEMA620P string = "NEMA 6-20P"

	// PowerPortTypeLabelNEMA630P captures enum value "NEMA 6-30P"
	PowerPortTypeLabelNEMA630P string = "NEMA 6-30P"

	// PowerPortTypeLabelNEMA650P captures enum value "NEMA 6-50P"
	PowerPortTypeLabelNEMA650P string = "NEMA 6-50P"

	// PowerPortTypeLabelNEMAL515P captures enum value "NEMA L5-15P"
	PowerPortTypeLabelNEMAL515P string = "NEMA L5-15P"

	// PowerPortTypeLabelNEMAL520P captures enum value "NEMA L5-20P"
	PowerPortTypeLabelNEMAL520P string = "NEMA L5-20P"

	// PowerPortTypeLabelNEMAL530P captures enum value "NEMA L5-30P"
	PowerPortTypeLabelNEMAL530P string = "NEMA L5-30P"

	// PowerPortTypeLabelNEMAL615P captures enum value "NEMA L6-15P"
	PowerPortTypeLabelNEMAL615P string = "NEMA L6-15P"

	// PowerPortTypeLabelNEMAL620P captures enum value "NEMA L6-20P"
	PowerPortTypeLabelNEMAL620P string = "NEMA L6-20P"

	// PowerPortTypeLabelNEMAL630P captures enum value "NEMA L6-30P"
	PowerPortTypeLabelNEMAL630P string = "NEMA L6-30P"

	// PowerPortTypeLabelNEMAL650P captures enum value "NEMA L6-50P"
	PowerPortTypeLabelNEMAL650P string = "NEMA L6-50P"

	// PowerPortTypeLabelCS6361C captures enum value "CS6361C"
	PowerPortTypeLabelCS6361C string = "CS6361C"

	// PowerPortTypeLabelCS6365C captures enum value "CS6365C"
	PowerPortTypeLabelCS6365C string = "CS6365C"

	// PowerPortTypeLabelCS8165C captures enum value "CS8165C"
	PowerPortTypeLabelCS8165C string = "CS8165C"

	// PowerPortTypeLabelCS8265C captures enum value "CS8265C"
	PowerPortTypeLabelCS8265C string = "CS8265C"

	// PowerPortTypeLabelCS8365C captures enum value "CS8365C"
	PowerPortTypeLabelCS8365C string = "CS8365C"

	// PowerPortTypeLabelCS8465C captures enum value "CS8465C"
	PowerPortTypeLabelCS8465C string = "CS8465C"

	// PowerPortTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE 7/5)"
	PowerPortTypeLabelITATypeECEE75 string = "ITA Type E (CEE 7/5)"

	// PowerPortTypeLabelITATypeFCEE74 captures enum value "ITA Type F (CEE 7/4)"
	PowerPortTypeLabelITATypeFCEE74 string = "ITA Type F (CEE 7/4)"

	// PowerPortTypeLabelITATypeEFCEE77 captures enum value "ITA Type E/F (CEE 7/7)"
	PowerPortTypeLabelITATypeEFCEE77 string = "ITA Type E/F (CEE 7/7)"

	// PowerPortTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerPortTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerPortTypeLabelITATypeH captures enum value "ITA Type H"
	PowerPortTypeLabelITATypeH string = "ITA Type H"

	// PowerPortTypeLabelITATypeI captures enum value "ITA Type I"
	PowerPortTypeLabelITATypeI string = "ITA Type I"

	// PowerPortTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerPortTypeLabelITATypeJ string = "ITA Type J"

	// PowerPortTypeLabelITATypeK captures enum value "ITA Type K"
	PowerPortTypeLabelITATypeK string = "ITA Type K"

	// PowerPortTypeLabelITATypeLCEI2350 captures enum value "ITA Type L (CEI 23-50)"
	PowerPortTypeLabelITATypeLCEI2350 string = "ITA Type L (CEI 23-50)"

	// PowerPortTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerPortTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerPortTypeLabelITATypeN captures enum value "ITA Type N"
	PowerPortTypeLabelITATypeN string = "ITA Type N"

	// PowerPortTypeLabelITATypeO captures enum value "ITA Type O"
	PowerPortTypeLabelITATypeO string = "ITA Type O"
)

// prop value enum
func (m *PowerPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerPortTypeTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-20p","nema-l6-30p","nema-l6-50p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTypeTypeValuePropEnum = append(powerPortTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerPortTypeValueIec60320C6 captures enum value "iec-60320-c6"
	PowerPortTypeValueIec60320C6 string = "iec-60320-c6"

	// PowerPortTypeValueIec60320C8 captures enum value "iec-60320-c8"
	PowerPortTypeValueIec60320C8 string = "iec-60320-c8"

	// PowerPortTypeValueIec60320C14 captures enum value "iec-60320-c14"
	PowerPortTypeValueIec60320C14 string = "iec-60320-c14"

	// PowerPortTypeValueIec60320C16 captures enum value "iec-60320-c16"
	PowerPortTypeValueIec60320C16 string = "iec-60320-c16"

	// PowerPortTypeValueIec60320C20 captures enum value "iec-60320-c20"
	PowerPortTypeValueIec60320C20 string = "iec-60320-c20"

	// PowerPortTypeValueIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	PowerPortTypeValueIec60309pne4h string = "iec-60309-p-n-e-4h"

	// PowerPortTypeValueIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	PowerPortTypeValueIec60309pne6h string = "iec-60309-p-n-e-6h"

	// PowerPortTypeValueIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	PowerPortTypeValueIec60309pne9h string = "iec-60309-p-n-e-9h"

	// PowerPortTypeValueIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	PowerPortTypeValueIec603092pe4h string = "iec-60309-2p-e-4h"

	// PowerPortTypeValueIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	PowerPortTypeValueIec603092pe6h string = "iec-60309-2p-e-6h"

	// PowerPortTypeValueIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	PowerPortTypeValueIec603092pe9h string = "iec-60309-2p-e-9h"

	// PowerPortTypeValueIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	PowerPortTypeValueIec603093pe4h string = "iec-60309-3p-e-4h"

	// PowerPortTypeValueIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	PowerPortTypeValueIec603093pe6h string = "iec-60309-3p-e-6h"

	// PowerPortTypeValueIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	PowerPortTypeValueIec603093pe9h string = "iec-60309-3p-e-9h"

	// PowerPortTypeValueIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	PowerPortTypeValueIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// PowerPortTypeValueIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	PowerPortTypeValueIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// PowerPortTypeValueIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	PowerPortTypeValueIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// PowerPortTypeValueNema515p captures enum value "nema-5-15p"
	PowerPortTypeValueNema515p string = "nema-5-15p"

	// PowerPortTypeValueNema520p captures enum value "nema-5-20p"
	PowerPortTypeValueNema520p string = "nema-5-20p"

	// PowerPortTypeValueNema530p captures enum value "nema-5-30p"
	PowerPortTypeValueNema530p string = "nema-5-30p"

	// PowerPortTypeValueNema550p captures enum value "nema-5-50p"
	PowerPortTypeValueNema550p string = "nema-5-50p"

	// PowerPortTypeValueNema615p captures enum value "nema-6-15p"
	PowerPortTypeValueNema615p string = "nema-6-15p"

	// PowerPortTypeValueNema620p captures enum value "nema-6-20p"
	PowerPortTypeValueNema620p string = "nema-6-20p"

	// PowerPortTypeValueNema630p captures enum value "nema-6-30p"
	PowerPortTypeValueNema630p string = "nema-6-30p"

	// PowerPortTypeValueNema650p captures enum value "nema-6-50p"
	PowerPortTypeValueNema650p string = "nema-6-50p"

	// PowerPortTypeValueNemaL515p captures enum value "nema-l5-15p"
	PowerPortTypeValueNemaL515p string = "nema-l5-15p"

	// PowerPortTypeValueNemaL520p captures enum value "nema-l5-20p"
	PowerPortTypeValueNemaL520p string = "nema-l5-20p"

	// PowerPortTypeValueNemaL530p captures enum value "nema-l5-30p"
	PowerPortTypeValueNemaL530p string = "nema-l5-30p"

	// PowerPortTypeValueNemaL550p captures enum value "nema-l5-50p"
	PowerPortTypeValueNemaL550p string = "nema-l5-50p"

	// PowerPortTypeValueNemaL620p captures enum value "nema-l6-20p"
	PowerPortTypeValueNemaL620p string = "nema-l6-20p"

	// PowerPortTypeValueNemaL630p captures enum value "nema-l6-30p"
	PowerPortTypeValueNemaL630p string = "nema-l6-30p"

	// PowerPortTypeValueNemaL650p captures enum value "nema-l6-50p"
	PowerPortTypeValueNemaL650p string = "nema-l6-50p"

	// PowerPortTypeValueCs6361c captures enum value "cs6361c"
	PowerPortTypeValueCs6361c string = "cs6361c"

	// PowerPortTypeValueCs6365c captures enum value "cs6365c"
	PowerPortTypeValueCs6365c string = "cs6365c"

	// PowerPortTypeValueCs8165c captures enum value "cs8165c"
	PowerPortTypeValueCs8165c string = "cs8165c"

	// PowerPortTypeValueCs8265c captures enum value "cs8265c"
	PowerPortTypeValueCs8265c string = "cs8265c"

	// PowerPortTypeValueCs8365c captures enum value "cs8365c"
	PowerPortTypeValueCs8365c string = "cs8365c"

	// PowerPortTypeValueCs8465c captures enum value "cs8465c"
	PowerPortTypeValueCs8465c string = "cs8465c"

	// PowerPortTypeValueItae captures enum value "ita-e"
	PowerPortTypeValueItae string = "ita-e"

	// PowerPortTypeValueItaf captures enum value "ita-f"
	PowerPortTypeValueItaf string = "ita-f"

	// PowerPortTypeValueItaEf captures enum value "ita-ef"
	PowerPortTypeValueItaEf string = "ita-ef"

	// PowerPortTypeValueItag captures enum value "ita-g"
	PowerPortTypeValueItag string = "ita-g"

	// PowerPortTypeValueItah captures enum value "ita-h"
	PowerPortTypeValueItah string = "ita-h"

	// PowerPortTypeValueItai captures enum value "ita-i"
	PowerPortTypeValueItai string = "ita-i"

	// PowerPortTypeValueItaj captures enum value "ita-j"
	PowerPortTypeValueItaj string = "ita-j"

	// PowerPortTypeValueItak captures enum value "ita-k"
	PowerPortTypeValueItak string = "ita-k"

	// PowerPortTypeValueItal captures enum value "ita-l"
	PowerPortTypeValueItal string = "ita-l"

	// PowerPortTypeValueItam captures enum value "ita-m"
	PowerPortTypeValueItam string = "ita-m"

	// PowerPortTypeValueItan captures enum value "ita-n"
	PowerPortTypeValueItan string = "ita-n"

	// PowerPortTypeValueItao captures enum value "ita-o"
	PowerPortTypeValueItao string = "ita-o"
)

// prop value enum
func (m *PowerPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerPortTypeTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortType) UnmarshalBinary(b []byte) error {
	var res PowerPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
