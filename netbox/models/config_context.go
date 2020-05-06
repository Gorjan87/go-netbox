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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigContext config context
// swagger:model ConfigContext
type ConfigContext struct {

	// cluster groups
	// Unique: true
	ClusterGroups []*NestedClusterGroup `json:"cluster_groups"`

	// clusters
	// Unique: true
	Clusters []*NestedCluster `json:"clusters"`

	// Data
	// Required: true
	Data *string `json:"data"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is active
	IsActive bool `json:"is_active,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// platforms
	// Unique: true
	Platforms []*NestedPlatform `json:"platforms"`

	// regions
	// Unique: true
	Regions []*NestedRegion `json:"regions"`

	// roles
	// Unique: true
	Roles []*NestedDeviceRole `json:"roles"`

	// sites
	// Unique: true
	Sites []*NestedSite `json:"sites"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// tenant groups
	// Unique: true
	TenantGroups []*NestedTenantGroup `json:"tenant_groups"`

	// tenants
	// Unique: true
	Tenants []*NestedTenant `json:"tenants"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this config context
func (m *ConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigContext) validateClusterGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("cluster_groups", "body", m.ClusterGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterGroups); i++ {
		if swag.IsZero(m.ClusterGroups[i]) { // not required
			continue
		}

		if m.ClusterGroups[i] != nil {
			if err := m.ClusterGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusters", "body", m.Clusters); err != nil {
		return err
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *ConfigContext) validatePlatforms(formats strfmt.Registry) error {

	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	if err := validate.UniqueItems("platforms", "body", m.Platforms); err != nil {
		return err
	}

	for i := 0; i < len(m.Platforms); i++ {
		if swag.IsZero(m.Platforms[i]) { // not required
			continue
		}

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateSites(formats strfmt.Registry) error {

	if swag.IsZero(m.Sites) { // not required
		return nil
	}

	if err := validate.UniqueItems("sites", "body", m.Sites); err != nil {
		return err
	}

	for i := 0; i < len(m.Sites); i++ {
		if swag.IsZero(m.Sites[i]) { // not required
			continue
		}

		if m.Sites[i] != nil {
			if err := m.Sites[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.Pattern("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), `^[-a-zA-Z0-9_]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConfigContext) validateTenantGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.TenantGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenant_groups", "body", m.TenantGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.TenantGroups); i++ {
		if swag.IsZero(m.TenantGroups[i]) { // not required
			continue
		}

		if m.TenantGroups[i] != nil {
			if err := m.TenantGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenant_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateTenants(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenants", "body", m.Tenants); err != nil {
		return err
	}

	for i := 0; i < len(m.Tenants); i++ {
		if swag.IsZero(m.Tenants[i]) { // not required
			continue
		}

		if m.Tenants[i] != nil {
			if err := m.Tenants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigContext) validateWeight(formats strfmt.Registry) error {

	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", int64(*m.Weight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", int64(*m.Weight), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigContext) UnmarshalBinary(b []byte) error {
	var res ConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
