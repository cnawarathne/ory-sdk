/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.16.10
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UpdateRegistrationFlowBody - Update Registration Request Body
type UpdateRegistrationFlowBody struct {
	UpdateRegistrationFlowWithCodeMethod *UpdateRegistrationFlowWithCodeMethod
	UpdateRegistrationFlowWithOidcMethod *UpdateRegistrationFlowWithOidcMethod
	UpdateRegistrationFlowWithPasskeyMethod *UpdateRegistrationFlowWithPasskeyMethod
	UpdateRegistrationFlowWithPasswordMethod *UpdateRegistrationFlowWithPasswordMethod
	UpdateRegistrationFlowWithProfileMethod *UpdateRegistrationFlowWithProfileMethod
	UpdateRegistrationFlowWithWebAuthnMethod *UpdateRegistrationFlowWithWebAuthnMethod
}

// UpdateRegistrationFlowWithCodeMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithCodeMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithCodeMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithCodeMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithCodeMethod: v,
	}
}

// UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithOidcMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithOidcMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithOidcMethod: v,
	}
}

// UpdateRegistrationFlowWithPasskeyMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithPasskeyMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithPasskeyMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithPasskeyMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithPasskeyMethod: v,
	}
}

// UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithPasswordMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithPasswordMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithPasswordMethod: v,
	}
}

// UpdateRegistrationFlowWithProfileMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithProfileMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithProfileMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithProfileMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithProfileMethod: v,
	}
}

// UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithWebAuthnMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithWebAuthnMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithWebAuthnMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRegistrationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'code'
	if jsonDict["method"] == "code" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithCodeMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithCodeMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithCodeMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithCodeMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithCodeMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'oidc'
	if jsonDict["method"] == "oidc" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithOidcMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithOidcMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithOidcMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithOidcMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithOidcMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'passkey'
	if jsonDict["method"] == "passkey" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasskeyMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasskeyMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasskeyMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasskeyMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasskeyMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password'
	if jsonDict["method"] == "password" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasswordMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasswordMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasswordMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasswordMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasswordMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'profile'
	if jsonDict["method"] == "profile" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithProfileMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithProfileMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithProfileMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithProfileMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithProfileMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["method"] == "webauthn" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithWebAuthnMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithWebAuthnMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithWebAuthnMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithWebAuthnMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithCodeMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithCodeMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithCodeMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithCodeMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithCodeMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithCodeMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithCodeMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithOidcMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithOidcMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithOidcMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithOidcMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithOidcMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithOidcMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithOidcMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithPasskeyMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithPasskeyMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasskeyMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasskeyMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasskeyMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasskeyMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasskeyMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithPasswordMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithPasswordMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasswordMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasswordMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasswordMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasswordMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasswordMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithProfileMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithProfileMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithProfileMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithProfileMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithProfileMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithProfileMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithProfileMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithWebAuthnMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithWebAuthnMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithWebAuthnMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithWebAuthnMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithWebAuthnMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithWebAuthnMethod: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	if src.UpdateRegistrationFlowWithCodeMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithCodeMethod)
	}

	if src.UpdateRegistrationFlowWithOidcMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithOidcMethod)
	}

	if src.UpdateRegistrationFlowWithPasskeyMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithPasskeyMethod)
	}

	if src.UpdateRegistrationFlowWithPasswordMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithPasswordMethod)
	}

	if src.UpdateRegistrationFlowWithProfileMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithProfileMethod)
	}

	if src.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithWebAuthnMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRegistrationFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateRegistrationFlowWithCodeMethod != nil {
		return obj.UpdateRegistrationFlowWithCodeMethod
	}

	if obj.UpdateRegistrationFlowWithOidcMethod != nil {
		return obj.UpdateRegistrationFlowWithOidcMethod
	}

	if obj.UpdateRegistrationFlowWithPasskeyMethod != nil {
		return obj.UpdateRegistrationFlowWithPasskeyMethod
	}

	if obj.UpdateRegistrationFlowWithPasswordMethod != nil {
		return obj.UpdateRegistrationFlowWithPasswordMethod
	}

	if obj.UpdateRegistrationFlowWithProfileMethod != nil {
		return obj.UpdateRegistrationFlowWithProfileMethod
	}

	if obj.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return obj.UpdateRegistrationFlowWithWebAuthnMethod
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRegistrationFlowBody struct {
	value *UpdateRegistrationFlowBody
	isSet bool
}

func (v NullableUpdateRegistrationFlowBody) Get() *UpdateRegistrationFlowBody {
	return v.value
}

func (v *NullableUpdateRegistrationFlowBody) Set(val *UpdateRegistrationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowBody(val *UpdateRegistrationFlowBody) *NullableUpdateRegistrationFlowBody {
	return &NullableUpdateRegistrationFlowBody{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


