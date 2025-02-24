/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.17.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PatchIdentitiesBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchIdentitiesBody{}

// PatchIdentitiesBody Patch Identities Body
type PatchIdentitiesBody struct {
	// Identities holds the list of patches to apply  required
	Identities []IdentityPatch `json:"identities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchIdentitiesBody PatchIdentitiesBody

// NewPatchIdentitiesBody instantiates a new PatchIdentitiesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchIdentitiesBody() *PatchIdentitiesBody {
	this := PatchIdentitiesBody{}
	return &this
}

// NewPatchIdentitiesBodyWithDefaults instantiates a new PatchIdentitiesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchIdentitiesBodyWithDefaults() *PatchIdentitiesBody {
	this := PatchIdentitiesBody{}
	return &this
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *PatchIdentitiesBody) GetIdentities() []IdentityPatch {
	if o == nil || IsNil(o.Identities) {
		var ret []IdentityPatch
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchIdentitiesBody) GetIdentitiesOk() ([]IdentityPatch, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *PatchIdentitiesBody) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []IdentityPatch and assigns it to the Identities field.
func (o *PatchIdentitiesBody) SetIdentities(v []IdentityPatch) {
	o.Identities = v
}

func (o PatchIdentitiesBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchIdentitiesBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchIdentitiesBody) UnmarshalJSON(data []byte) (err error) {
	varPatchIdentitiesBody := _PatchIdentitiesBody{}

	err = json.Unmarshal(data, &varPatchIdentitiesBody)

	if err != nil {
		return err
	}

	*o = PatchIdentitiesBody(varPatchIdentitiesBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchIdentitiesBody struct {
	value *PatchIdentitiesBody
	isSet bool
}

func (v NullablePatchIdentitiesBody) Get() *PatchIdentitiesBody {
	return v.value
}

func (v *NullablePatchIdentitiesBody) Set(val *PatchIdentitiesBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchIdentitiesBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchIdentitiesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchIdentitiesBody(val *PatchIdentitiesBody) *NullablePatchIdentitiesBody {
	return &NullablePatchIdentitiesBody{value: val, isSet: true}
}

func (v NullablePatchIdentitiesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchIdentitiesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


