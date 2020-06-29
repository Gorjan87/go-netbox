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

// PowerOutletTemplate power outlet template
// swagger:model PowerOutletTemplate
type PowerOutletTemplate struct {

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// feed leg
	FeedLeg *PowerOutletTemplateFeedLeg `json:"feed_leg,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// power port
	PowerPort *NestedPowerPortTemplate `json:"power_port,omitempty"`

	// type
	Type *PowerOutletTemplateType `json:"type,omitempty"`
}

// Validate validates this power outlet template
func (m *PowerOutletTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPort(formats); err != nil {
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

func (m *PowerOutletTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) validateFeedLeg(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	if m.FeedLeg != nil {
		if err := m.FeedLeg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feed_leg")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) validateName(formats strfmt.Registry) error {

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

func (m *PowerOutletTemplate) validatePowerPort(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if m.PowerPort != nil {
		if err := m.PowerPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutletTemplate) validateType(formats strfmt.Registry) error {

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
func (m *PowerOutletTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplate) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletTemplateFeedLeg Feed leg
// swagger:model PowerOutletTemplateFeedLeg
type PowerOutletTemplateFeedLeg struct {

	// label
	// Required: true
	// Enum: [A B C]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [A B C]
	Value *string `json:"value"`
}

// Validate validates this power outlet template feed leg
func (m *PowerOutletTemplateFeedLeg) Validate(formats strfmt.Registry) error {
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

var powerOutletTemplateFeedLegTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateFeedLegTypeLabelPropEnum = append(powerOutletTemplateFeedLegTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTemplateFeedLegLabelA captures enum value "A"
	PowerOutletTemplateFeedLegLabelA string = "A"

	// PowerOutletTemplateFeedLegLabelB captures enum value "B"
	PowerOutletTemplateFeedLegLabelB string = "B"

	// PowerOutletTemplateFeedLegLabelC captures enum value "C"
	PowerOutletTemplateFeedLegLabelC string = "C"
)

// prop value enum
func (m *PowerOutletTemplateFeedLeg) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerOutletTemplateFeedLegTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateFeedLeg) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("feed_leg"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTemplateFeedLegTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateFeedLegTypeValuePropEnum = append(powerOutletTemplateFeedLegTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTemplateFeedLegValueA captures enum value "A"
	PowerOutletTemplateFeedLegValueA string = "A"

	// PowerOutletTemplateFeedLegValueB captures enum value "B"
	PowerOutletTemplateFeedLegValueB string = "B"

	// PowerOutletTemplateFeedLegValueC captures enum value "C"
	PowerOutletTemplateFeedLegValueC string = "C"
)

// prop value enum
func (m *PowerOutletTemplateFeedLeg) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerOutletTemplateFeedLegTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateFeedLeg) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("feed_leg"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletTemplateFeedLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplateFeedLeg) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplateFeedLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletTemplateType Type
// swagger:model PowerOutletTemplateType
type PowerOutletTemplateType struct {

	// label
	// Required: true
	// Enum: [C5 C7 C13 C15 C19 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 5-15R NEMA 5-20R NEMA 5-30R NEMA 5-50R NEMA 6-15R NEMA 6-20R NEMA 6-30R NEMA 6-50R NEMA L5-15R NEMA L5-20R NEMA L5-30R NEMA L6-15R NEMA L6-20R NEMA L6-30R NEMA L6-50R CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ITA Type E (CEE7/5) ITA Type F (CEE7/3) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O HDOT Cx]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-20r nema-l6-30r nema-l6-50r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o hdot-cx]
	Value *string `json:"value"`
}

// Validate validates this power outlet template type
func (m *PowerOutletTemplateType) Validate(formats strfmt.Registry) error {
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

var powerOutletTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C5","C7","C13","C15","C19","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 5-15R","NEMA 5-20R","NEMA 5-30R","NEMA 5-50R","NEMA 6-15R","NEMA 6-20R","NEMA 6-30R","NEMA 6-50R","NEMA L5-15R","NEMA L5-20R","NEMA L5-30R","NEMA L6-15R","NEMA L6-20R","NEMA L6-30R","NEMA L6-50R","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ITA Type E (CEE7/5)","ITA Type F (CEE7/3)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","HDOT Cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateTypeTypeLabelPropEnum = append(powerOutletTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTemplateTypeLabelC5 captures enum value "C5"
	PowerOutletTemplateTypeLabelC5 string = "C5"

	// PowerOutletTemplateTypeLabelC7 captures enum value "C7"
	PowerOutletTemplateTypeLabelC7 string = "C7"

	// PowerOutletTemplateTypeLabelC13 captures enum value "C13"
	PowerOutletTemplateTypeLabelC13 string = "C13"

	// PowerOutletTemplateTypeLabelC15 captures enum value "C15"
	PowerOutletTemplateTypeLabelC15 string = "C15"

	// PowerOutletTemplateTypeLabelC19 captures enum value "C19"
	PowerOutletTemplateTypeLabelC19 string = "C19"

	// PowerOutletTemplateTypeLabelPNE4H captures enum value "P+N+E 4H"
	PowerOutletTemplateTypeLabelPNE4H string = "P+N+E 4H"

	// PowerOutletTemplateTypeLabelPNE6H captures enum value "P+N+E 6H"
	PowerOutletTemplateTypeLabelPNE6H string = "P+N+E 6H"

	// PowerOutletTemplateTypeLabelPNE9H captures enum value "P+N+E 9H"
	PowerOutletTemplateTypeLabelPNE9H string = "P+N+E 9H"

	// PowerOutletTemplateTypeLabelNr2PE4H captures enum value "2P+E 4H"
	PowerOutletTemplateTypeLabelNr2PE4H string = "2P+E 4H"

	// PowerOutletTemplateTypeLabelNr2PE6H captures enum value "2P+E 6H"
	PowerOutletTemplateTypeLabelNr2PE6H string = "2P+E 6H"

	// PowerOutletTemplateTypeLabelNr2PE9H captures enum value "2P+E 9H"
	PowerOutletTemplateTypeLabelNr2PE9H string = "2P+E 9H"

	// PowerOutletTemplateTypeLabelNr3PE4H captures enum value "3P+E 4H"
	PowerOutletTemplateTypeLabelNr3PE4H string = "3P+E 4H"

	// PowerOutletTemplateTypeLabelNr3PE6H captures enum value "3P+E 6H"
	PowerOutletTemplateTypeLabelNr3PE6H string = "3P+E 6H"

	// PowerOutletTemplateTypeLabelNr3PE9H captures enum value "3P+E 9H"
	PowerOutletTemplateTypeLabelNr3PE9H string = "3P+E 9H"

	// PowerOutletTemplateTypeLabelNr3PNE4H captures enum value "3P+N+E 4H"
	PowerOutletTemplateTypeLabelNr3PNE4H string = "3P+N+E 4H"

	// PowerOutletTemplateTypeLabelNr3PNE6H captures enum value "3P+N+E 6H"
	PowerOutletTemplateTypeLabelNr3PNE6H string = "3P+N+E 6H"

	// PowerOutletTemplateTypeLabelNr3PNE9H captures enum value "3P+N+E 9H"
	PowerOutletTemplateTypeLabelNr3PNE9H string = "3P+N+E 9H"

	// PowerOutletTemplateTypeLabelNEMA515R captures enum value "NEMA 5-15R"
	PowerOutletTemplateTypeLabelNEMA515R string = "NEMA 5-15R"

	// PowerOutletTemplateTypeLabelNEMA520R captures enum value "NEMA 5-20R"
	PowerOutletTemplateTypeLabelNEMA520R string = "NEMA 5-20R"

	// PowerOutletTemplateTypeLabelNEMA530R captures enum value "NEMA 5-30R"
	PowerOutletTemplateTypeLabelNEMA530R string = "NEMA 5-30R"

	// PowerOutletTemplateTypeLabelNEMA550R captures enum value "NEMA 5-50R"
	PowerOutletTemplateTypeLabelNEMA550R string = "NEMA 5-50R"

	// PowerOutletTemplateTypeLabelNEMA615R captures enum value "NEMA 6-15R"
	PowerOutletTemplateTypeLabelNEMA615R string = "NEMA 6-15R"

	// PowerOutletTemplateTypeLabelNEMA620R captures enum value "NEMA 6-20R"
	PowerOutletTemplateTypeLabelNEMA620R string = "NEMA 6-20R"

	// PowerOutletTemplateTypeLabelNEMA630R captures enum value "NEMA 6-30R"
	PowerOutletTemplateTypeLabelNEMA630R string = "NEMA 6-30R"

	// PowerOutletTemplateTypeLabelNEMA650R captures enum value "NEMA 6-50R"
	PowerOutletTemplateTypeLabelNEMA650R string = "NEMA 6-50R"

	// PowerOutletTemplateTypeLabelNEMAL515R captures enum value "NEMA L5-15R"
	PowerOutletTemplateTypeLabelNEMAL515R string = "NEMA L5-15R"

	// PowerOutletTemplateTypeLabelNEMAL520R captures enum value "NEMA L5-20R"
	PowerOutletTemplateTypeLabelNEMAL520R string = "NEMA L5-20R"

	// PowerOutletTemplateTypeLabelNEMAL530R captures enum value "NEMA L5-30R"
	PowerOutletTemplateTypeLabelNEMAL530R string = "NEMA L5-30R"

	// PowerOutletTemplateTypeLabelNEMAL615R captures enum value "NEMA L6-15R"
	PowerOutletTemplateTypeLabelNEMAL615R string = "NEMA L6-15R"

	// PowerOutletTemplateTypeLabelNEMAL620R captures enum value "NEMA L6-20R"
	PowerOutletTemplateTypeLabelNEMAL620R string = "NEMA L6-20R"

	// PowerOutletTemplateTypeLabelNEMAL630R captures enum value "NEMA L6-30R"
	PowerOutletTemplateTypeLabelNEMAL630R string = "NEMA L6-30R"

	// PowerOutletTemplateTypeLabelNEMAL650R captures enum value "NEMA L6-50R"
	PowerOutletTemplateTypeLabelNEMAL650R string = "NEMA L6-50R"

	// PowerOutletTemplateTypeLabelCS6360C captures enum value "CS6360C"
	PowerOutletTemplateTypeLabelCS6360C string = "CS6360C"

	// PowerOutletTemplateTypeLabelCS6364C captures enum value "CS6364C"
	PowerOutletTemplateTypeLabelCS6364C string = "CS6364C"

	// PowerOutletTemplateTypeLabelCS8164C captures enum value "CS8164C"
	PowerOutletTemplateTypeLabelCS8164C string = "CS8164C"

	// PowerOutletTemplateTypeLabelCS8264C captures enum value "CS8264C"
	PowerOutletTemplateTypeLabelCS8264C string = "CS8264C"

	// PowerOutletTemplateTypeLabelCS8364C captures enum value "CS8364C"
	PowerOutletTemplateTypeLabelCS8364C string = "CS8364C"

	// PowerOutletTemplateTypeLabelCS8464C captures enum value "CS8464C"
	PowerOutletTemplateTypeLabelCS8464C string = "CS8464C"

	// PowerOutletTemplateTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE7/5)"
	PowerOutletTemplateTypeLabelITATypeECEE75 string = "ITA Type E (CEE7/5)"

	// PowerOutletTemplateTypeLabelITATypeFCEE73 captures enum value "ITA Type F (CEE7/3)"
	PowerOutletTemplateTypeLabelITATypeFCEE73 string = "ITA Type F (CEE7/3)"

	// PowerOutletTemplateTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerOutletTemplateTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerOutletTemplateTypeLabelITATypeH captures enum value "ITA Type H"
	PowerOutletTemplateTypeLabelITATypeH string = "ITA Type H"

	// PowerOutletTemplateTypeLabelITATypeI captures enum value "ITA Type I"
	PowerOutletTemplateTypeLabelITATypeI string = "ITA Type I"

	// PowerOutletTemplateTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerOutletTemplateTypeLabelITATypeJ string = "ITA Type J"

	// PowerOutletTemplateTypeLabelITATypeK captures enum value "ITA Type K"
	PowerOutletTemplateTypeLabelITATypeK string = "ITA Type K"

	// PowerOutletTemplateTypeLabelITATypeLCEI2350 captures enum value "ITA Type L (CEI 23-50)"
	PowerOutletTemplateTypeLabelITATypeLCEI2350 string = "ITA Type L (CEI 23-50)"

	// PowerOutletTemplateTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerOutletTemplateTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerOutletTemplateTypeLabelITATypeN captures enum value "ITA Type N"
	PowerOutletTemplateTypeLabelITATypeN string = "ITA Type N"

	// PowerOutletTemplateTypeLabelITATypeO captures enum value "ITA Type O"
	PowerOutletTemplateTypeLabelITATypeO string = "ITA Type O"

	// PowerOutletTemplateTypeLabelHDOTCx captures enum value "HDOT Cx"
	PowerOutletTemplateTypeLabelHDOTCx string = "HDOT Cx"
)

// prop value enum
func (m *PowerOutletTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerOutletTemplateTypeTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-20r","nema-l6-30r","nema-l6-50r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","hdot-cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTemplateTypeTypeValuePropEnum = append(powerOutletTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTemplateTypeValueIec60320C5 captures enum value "iec-60320-c5"
	PowerOutletTemplateTypeValueIec60320C5 string = "iec-60320-c5"

	// PowerOutletTemplateTypeValueIec60320C7 captures enum value "iec-60320-c7"
	PowerOutletTemplateTypeValueIec60320C7 string = "iec-60320-c7"

	// PowerOutletTemplateTypeValueIec60320C13 captures enum value "iec-60320-c13"
	PowerOutletTemplateTypeValueIec60320C13 string = "iec-60320-c13"

	// PowerOutletTemplateTypeValueIec60320C15 captures enum value "iec-60320-c15"
	PowerOutletTemplateTypeValueIec60320C15 string = "iec-60320-c15"

	// PowerOutletTemplateTypeValueIec60320C19 captures enum value "iec-60320-c19"
	PowerOutletTemplateTypeValueIec60320C19 string = "iec-60320-c19"

	// PowerOutletTemplateTypeValueIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	PowerOutletTemplateTypeValueIec60309pne4h string = "iec-60309-p-n-e-4h"

	// PowerOutletTemplateTypeValueIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	PowerOutletTemplateTypeValueIec60309pne6h string = "iec-60309-p-n-e-6h"

	// PowerOutletTemplateTypeValueIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	PowerOutletTemplateTypeValueIec60309pne9h string = "iec-60309-p-n-e-9h"

	// PowerOutletTemplateTypeValueIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	PowerOutletTemplateTypeValueIec603092pe4h string = "iec-60309-2p-e-4h"

	// PowerOutletTemplateTypeValueIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	PowerOutletTemplateTypeValueIec603092pe6h string = "iec-60309-2p-e-6h"

	// PowerOutletTemplateTypeValueIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	PowerOutletTemplateTypeValueIec603092pe9h string = "iec-60309-2p-e-9h"

	// PowerOutletTemplateTypeValueIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	PowerOutletTemplateTypeValueIec603093pe4h string = "iec-60309-3p-e-4h"

	// PowerOutletTemplateTypeValueIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	PowerOutletTemplateTypeValueIec603093pe6h string = "iec-60309-3p-e-6h"

	// PowerOutletTemplateTypeValueIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	PowerOutletTemplateTypeValueIec603093pe9h string = "iec-60309-3p-e-9h"

	// PowerOutletTemplateTypeValueIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	PowerOutletTemplateTypeValueIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// PowerOutletTemplateTypeValueIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	PowerOutletTemplateTypeValueIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// PowerOutletTemplateTypeValueIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	PowerOutletTemplateTypeValueIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// PowerOutletTemplateTypeValueNema515r captures enum value "nema-5-15r"
	PowerOutletTemplateTypeValueNema515r string = "nema-5-15r"

	// PowerOutletTemplateTypeValueNema520r captures enum value "nema-5-20r"
	PowerOutletTemplateTypeValueNema520r string = "nema-5-20r"

	// PowerOutletTemplateTypeValueNema530r captures enum value "nema-5-30r"
	PowerOutletTemplateTypeValueNema530r string = "nema-5-30r"

	// PowerOutletTemplateTypeValueNema550r captures enum value "nema-5-50r"
	PowerOutletTemplateTypeValueNema550r string = "nema-5-50r"

	// PowerOutletTemplateTypeValueNema615r captures enum value "nema-6-15r"
	PowerOutletTemplateTypeValueNema615r string = "nema-6-15r"

	// PowerOutletTemplateTypeValueNema620r captures enum value "nema-6-20r"
	PowerOutletTemplateTypeValueNema620r string = "nema-6-20r"

	// PowerOutletTemplateTypeValueNema630r captures enum value "nema-6-30r"
	PowerOutletTemplateTypeValueNema630r string = "nema-6-30r"

	// PowerOutletTemplateTypeValueNema650r captures enum value "nema-6-50r"
	PowerOutletTemplateTypeValueNema650r string = "nema-6-50r"

	// PowerOutletTemplateTypeValueNemaL515r captures enum value "nema-l5-15r"
	PowerOutletTemplateTypeValueNemaL515r string = "nema-l5-15r"

	// PowerOutletTemplateTypeValueNemaL520r captures enum value "nema-l5-20r"
	PowerOutletTemplateTypeValueNemaL520r string = "nema-l5-20r"

	// PowerOutletTemplateTypeValueNemaL530r captures enum value "nema-l5-30r"
	PowerOutletTemplateTypeValueNemaL530r string = "nema-l5-30r"

	// PowerOutletTemplateTypeValueNemaL550r captures enum value "nema-l5-50r"
	PowerOutletTemplateTypeValueNemaL550r string = "nema-l5-50r"

	// PowerOutletTemplateTypeValueNemaL620r captures enum value "nema-l6-20r"
	PowerOutletTemplateTypeValueNemaL620r string = "nema-l6-20r"

	// PowerOutletTemplateTypeValueNemaL630r captures enum value "nema-l6-30r"
	PowerOutletTemplateTypeValueNemaL630r string = "nema-l6-30r"

	// PowerOutletTemplateTypeValueNemaL650r captures enum value "nema-l6-50r"
	PowerOutletTemplateTypeValueNemaL650r string = "nema-l6-50r"

	// PowerOutletTemplateTypeValueCS6360C captures enum value "CS6360C"
	PowerOutletTemplateTypeValueCS6360C string = "CS6360C"

	// PowerOutletTemplateTypeValueCS6364C captures enum value "CS6364C"
	PowerOutletTemplateTypeValueCS6364C string = "CS6364C"

	// PowerOutletTemplateTypeValueCS8164C captures enum value "CS8164C"
	PowerOutletTemplateTypeValueCS8164C string = "CS8164C"

	// PowerOutletTemplateTypeValueCS8264C captures enum value "CS8264C"
	PowerOutletTemplateTypeValueCS8264C string = "CS8264C"

	// PowerOutletTemplateTypeValueCS8364C captures enum value "CS8364C"
	PowerOutletTemplateTypeValueCS8364C string = "CS8364C"

	// PowerOutletTemplateTypeValueCS8464C captures enum value "CS8464C"
	PowerOutletTemplateTypeValueCS8464C string = "CS8464C"

	// PowerOutletTemplateTypeValueItae captures enum value "ita-e"
	PowerOutletTemplateTypeValueItae string = "ita-e"

	// PowerOutletTemplateTypeValueItaf captures enum value "ita-f"
	PowerOutletTemplateTypeValueItaf string = "ita-f"

	// PowerOutletTemplateTypeValueItag captures enum value "ita-g"
	PowerOutletTemplateTypeValueItag string = "ita-g"

	// PowerOutletTemplateTypeValueItah captures enum value "ita-h"
	PowerOutletTemplateTypeValueItah string = "ita-h"

	// PowerOutletTemplateTypeValueItai captures enum value "ita-i"
	PowerOutletTemplateTypeValueItai string = "ita-i"

	// PowerOutletTemplateTypeValueItaj captures enum value "ita-j"
	PowerOutletTemplateTypeValueItaj string = "ita-j"

	// PowerOutletTemplateTypeValueItak captures enum value "ita-k"
	PowerOutletTemplateTypeValueItak string = "ita-k"

	// PowerOutletTemplateTypeValueItal captures enum value "ita-l"
	PowerOutletTemplateTypeValueItal string = "ita-l"

	// PowerOutletTemplateTypeValueItam captures enum value "ita-m"
	PowerOutletTemplateTypeValueItam string = "ita-m"

	// PowerOutletTemplateTypeValueItan captures enum value "ita-n"
	PowerOutletTemplateTypeValueItan string = "ita-n"

	// PowerOutletTemplateTypeValueItao captures enum value "ita-o"
	PowerOutletTemplateTypeValueItao string = "ita-o"

	// PowerOutletTemplateTypeValueHdotCx captures enum value "hdot-cx"
	PowerOutletTemplateTypeValueHdotCx string = "hdot-cx"
)

// prop value enum
func (m *PowerOutletTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, powerOutletTemplateTypeTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletTemplateType) validateValue(formats strfmt.Registry) error {

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
func (m *PowerOutletTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletTemplateType) UnmarshalBinary(b []byte) error {
	var res PowerOutletTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
