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

// APIKey A PEM-format RSA credential for securing requests to the Oracle Bare Metal Cloud Services REST API. Also known
// as an *API signing key*. Specifically, this is the public key from the key pair. The private key remains with
// the user calling the API. For information about generating a key pair
// in the required PEM format, see [Required Keys and OCIDs](/Content/API/Concepts/apisigningkey.htm).
//
// **Important:** This is **not** the SSH key for accessing compute instances.
//
// Each user can have a maximum of three API signing keys.
//
// For more information about user credentials, see [User Credentials](/Content/Identity/Concepts/usercredentials.htm).
//
// swagger:model ApiKey
type APIKey struct {

	// The key's fingerprint (e.g., 12:34:56:78:90:ab:cd:ef:12:34:56:78:90:ab:cd:ef).
	Fingerprint string `json:"fingerprint,omitempty"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`

	// An Oracle-assigned identifier for the key, in this format:
	// TENANCY_OCID/USER_OCID/KEY_FINGERPRINT.
	//
	KeyID string `json:"keyId,omitempty"`

	// The key's value.
	KeyValue string `json:"keyValue,omitempty"`

	// The API key's current state. After creating an `ApiKey` object, make sure its `lifecycleState` changes from
	// CREATING to ACTIVE before using it.
	//
	// Max Length: 64
	// Min Length: 1
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Date and time the `ApiKey` object was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The OCID of the user the key belongs to.
	UserID string `json:"userId,omitempty"`
}

// Validate validates this Api key
func (m *APIKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apiKeyTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","INACTIVE","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiKeyTypeLifecycleStatePropEnum = append(apiKeyTypeLifecycleStatePropEnum, v)
	}
}

const (
	// APIKeyLifecycleStateCREATING captures enum value "CREATING"
	APIKeyLifecycleStateCREATING string = "CREATING"
	// APIKeyLifecycleStateACTIVE captures enum value "ACTIVE"
	APIKeyLifecycleStateACTIVE string = "ACTIVE"
	// APIKeyLifecycleStateINACTIVE captures enum value "INACTIVE"
	APIKeyLifecycleStateINACTIVE string = "INACTIVE"
	// APIKeyLifecycleStateDELETING captures enum value "DELETING"
	APIKeyLifecycleStateDELETING string = "DELETING"
	// APIKeyLifecycleStateDELETED captures enum value "DELETED"
	APIKeyLifecycleStateDELETED string = "DELETED"
)

// prop value enum
func (m *APIKey) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, apiKeyTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *APIKey) validateLifecycleState(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleState) { // not required
		return nil
	}

	if err := validate.MinLength("lifecycleState", "body", string(m.LifecycleState), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("lifecycleState", "body", string(m.LifecycleState), 64); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	return nil
}
