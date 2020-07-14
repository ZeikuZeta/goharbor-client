// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessLog access log
//
// swagger:model AccessLog
type AccessLog struct {

	// The ID of the log entry.
	LogID int64 `json:"log_id,omitempty"`

	// The time when this operation is triggered.
	OpTime string `json:"op_time,omitempty"`

	// The operation against the repository in this log entry.
	Operation string `json:"operation,omitempty"`

	// Name of the repository in this log entry.
	RepoName string `json:"repo_name,omitempty"`

	// Tag of the repository in this log entry.
	RepoTag string `json:"repo_tag,omitempty"`

	// Username of the user in this log entry.
	Username string `json:"username,omitempty"`
}

// Validate validates this access log
func (m *AccessLog) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessLog) UnmarshalBinary(b []byte) error {
	var res AccessLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}