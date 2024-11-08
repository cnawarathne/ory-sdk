/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.15.10
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VerifiableCredentialResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifiableCredentialResponse{}

// VerifiableCredentialResponse struct for VerifiableCredentialResponse
type VerifiableCredentialResponse struct {
	CredentialDraft00 *string `json:"credential_draft_00,omitempty"`
	Format *string `json:"format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifiableCredentialResponse VerifiableCredentialResponse

// NewVerifiableCredentialResponse instantiates a new VerifiableCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiableCredentialResponse() *VerifiableCredentialResponse {
	this := VerifiableCredentialResponse{}
	return &this
}

// NewVerifiableCredentialResponseWithDefaults instantiates a new VerifiableCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiableCredentialResponseWithDefaults() *VerifiableCredentialResponse {
	this := VerifiableCredentialResponse{}
	return &this
}

// GetCredentialDraft00 returns the CredentialDraft00 field value if set, zero value otherwise.
func (o *VerifiableCredentialResponse) GetCredentialDraft00() string {
	if o == nil || IsNil(o.CredentialDraft00) {
		var ret string
		return ret
	}
	return *o.CredentialDraft00
}

// GetCredentialDraft00Ok returns a tuple with the CredentialDraft00 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialResponse) GetCredentialDraft00Ok() (*string, bool) {
	if o == nil || IsNil(o.CredentialDraft00) {
		return nil, false
	}
	return o.CredentialDraft00, true
}

// HasCredentialDraft00 returns a boolean if a field has been set.
func (o *VerifiableCredentialResponse) HasCredentialDraft00() bool {
	if o != nil && !IsNil(o.CredentialDraft00) {
		return true
	}

	return false
}

// SetCredentialDraft00 gets a reference to the given string and assigns it to the CredentialDraft00 field.
func (o *VerifiableCredentialResponse) SetCredentialDraft00(v string) {
	o.CredentialDraft00 = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *VerifiableCredentialResponse) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialResponse) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *VerifiableCredentialResponse) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *VerifiableCredentialResponse) SetFormat(v string) {
	o.Format = &v
}

func (o VerifiableCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifiableCredentialResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialDraft00) {
		toSerialize["credential_draft_00"] = o.CredentialDraft00
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifiableCredentialResponse) UnmarshalJSON(data []byte) (err error) {
	varVerifiableCredentialResponse := _VerifiableCredentialResponse{}

	err = json.Unmarshal(data, &varVerifiableCredentialResponse)

	if err != nil {
		return err
	}

	*o = VerifiableCredentialResponse(varVerifiableCredentialResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credential_draft_00")
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifiableCredentialResponse struct {
	value *VerifiableCredentialResponse
	isSet bool
}

func (v NullableVerifiableCredentialResponse) Get() *VerifiableCredentialResponse {
	return v.value
}

func (v *NullableVerifiableCredentialResponse) Set(val *VerifiableCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiableCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiableCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiableCredentialResponse(val *VerifiableCredentialResponse) *NullableVerifiableCredentialResponse {
	return &NullableVerifiableCredentialResponse{value: val, isSet: true}
}

func (v NullableVerifiableCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiableCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


