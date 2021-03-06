// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationTask The replication task
//
// swagger:model ReplicationTask
type ReplicationTask struct {

	// The destination resource
	DstResource string `json:"dst_resource,omitempty"`

	// The end time
	EndTime string `json:"end_time,omitempty"`

	// The execution ID
	ExecutionID int64 `json:"execution_id,omitempty"`

	// The ID
	ID int64 `json:"id,omitempty"`

	// The job ID
	JobID string `json:"job_id,omitempty"`

	// The resource type
	ResourceType string `json:"resource_type,omitempty"`

	// The source resource
	SrcResource string `json:"src_resource,omitempty"`

	// The start time
	StartTime string `json:"start_time,omitempty"`

	// The status
	Status string `json:"status,omitempty"`
}

// Validate validates this replication task
func (m *ReplicationTask) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replication task based on context it is used
func (m *ReplicationTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTask) UnmarshalBinary(b []byte) error {
	var res ReplicationTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
