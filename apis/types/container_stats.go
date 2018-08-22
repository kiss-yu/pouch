// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerStats container stats almost from cgroup resource usage.
// swagger:model ContainerStats

type ContainerStats struct {

	// blkio stats
	BlkioStats *BlkioStats `json:"blkio_stats,omitempty"`

	// cpu stats
	CPUStats *CPUStats `json:"cpu_stats,omitempty"`

	// container id
	ID string `json:"id,omitempty"`

	// memory stats
	MemoryStats *MemoryStats `json:"memory_stats,omitempty"`

	// container name
	Name string `json:"name,omitempty"`

	// networks
	Networks map[string]NetworkStats `json:"networks,omitempty"`

	// pids stats
	PidsStats *PidsStats `json:"pids_stats,omitempty"`

	// precpu stats
	PrecpuStats *CPUStats `json:"precpu_stats,omitempty"`

	// read time of container stats.
	Read strfmt.DateTime `json:"read,omitempty"`
}

/* polymorph ContainerStats blkio_stats false */

/* polymorph ContainerStats cpu_stats false */

/* polymorph ContainerStats id false */

/* polymorph ContainerStats memory_stats false */

/* polymorph ContainerStats name false */

/* polymorph ContainerStats networks false */

/* polymorph ContainerStats pids_stats false */

/* polymorph ContainerStats precpu_stats false */

/* polymorph ContainerStats read false */

// Validate validates this container stats
func (m *ContainerStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlkioStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCPUStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemoryStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePidsStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrecpuStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerStats) validateBlkioStats(formats strfmt.Registry) error {

	if swag.IsZero(m.BlkioStats) { // not required
		return nil
	}

	if m.BlkioStats != nil {

		if err := m.BlkioStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blkio_stats")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerStats) validateCPUStats(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUStats) { // not required
		return nil
	}

	if m.CPUStats != nil {

		if err := m.CPUStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu_stats")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerStats) validateMemoryStats(formats strfmt.Registry) error {

	if swag.IsZero(m.MemoryStats) { // not required
		return nil
	}

	if m.MemoryStats != nil {

		if err := m.MemoryStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_stats")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerStats) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	return nil
}

func (m *ContainerStats) validatePidsStats(formats strfmt.Registry) error {

	if swag.IsZero(m.PidsStats) { // not required
		return nil
	}

	if m.PidsStats != nil {

		if err := m.PidsStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pids_stats")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerStats) validatePrecpuStats(formats strfmt.Registry) error {

	if swag.IsZero(m.PrecpuStats) { // not required
		return nil
	}

	if m.PrecpuStats != nil {

		if err := m.PrecpuStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("precpu_stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerStats) UnmarshalBinary(b []byte) error {
	var res ContainerStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
