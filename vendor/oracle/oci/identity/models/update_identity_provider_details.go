package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// UpdateIdentityProviderDetails update identity provider details
// swagger:discriminator UpdateIdentityProviderDetails protocol
type UpdateIdentityProviderDetails interface {
	runtime.Validatable

	// The description you assign to the `IdentityProvider`. Does not have to
	// be unique, and it's changeable.
	//
	// Max Length: 400
	// Min Length: 1
	Description() string
	SetDescription(string)

	// The protocol used for federation.
	//
	// Example: `SAML2`
	//
	// Required: true
	Protocol() string
	SetProtocol(string)
}

// UnmarshalUpdateIdentityProviderDetailsSlice unmarshals polymorphic slices of UpdateIdentityProviderDetails
func UnmarshalUpdateIdentityProviderDetailsSlice(reader io.Reader, consumer runtime.Consumer) ([]UpdateIdentityProviderDetails, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []UpdateIdentityProviderDetails
	for _, element := range elements {
		obj, err := unmarshalUpdateIdentityProviderDetails(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalUpdateIdentityProviderDetails unmarshals polymorphic UpdateIdentityProviderDetails
func UnmarshalUpdateIdentityProviderDetails(reader io.Reader, consumer runtime.Consumer) (UpdateIdentityProviderDetails, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalUpdateIdentityProviderDetails(data, consumer)
}

func unmarshalUpdateIdentityProviderDetails(data []byte, consumer runtime.Consumer) (UpdateIdentityProviderDetails, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the protocol property.
	var getType struct {
		Protocol string `json:"protocol"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("protocol", "body", getType.Protocol); err != nil {
		return nil, err
	}

	// The value of protocol is used to determine which type to create and unmarshal the data into
	switch getType.Protocol {
	case "UpdateSaml2IdentityProviderDetails":
		var result UpdateSaml2IdentityProviderDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid protocol value: %q", getType.Protocol)

}
