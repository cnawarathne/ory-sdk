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

// checks if the IdentityCredentialsCodeAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCredentialsCodeAddress{}

// IdentityCredentialsCodeAddress struct for IdentityCredentialsCodeAddress
type IdentityCredentialsCodeAddress struct {
	// The address for this code
	Address *string `json:"address,omitempty"`
	Channel *string `json:"channel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityCredentialsCodeAddress IdentityCredentialsCodeAddress

// NewIdentityCredentialsCodeAddress instantiates a new IdentityCredentialsCodeAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentialsCodeAddress() *IdentityCredentialsCodeAddress {
	this := IdentityCredentialsCodeAddress{}
	return &this
}

// NewIdentityCredentialsCodeAddressWithDefaults instantiates a new IdentityCredentialsCodeAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialsCodeAddressWithDefaults() *IdentityCredentialsCodeAddress {
	this := IdentityCredentialsCodeAddress{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IdentityCredentialsCodeAddress) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialsCodeAddress) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IdentityCredentialsCodeAddress) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IdentityCredentialsCodeAddress) SetAddress(v string) {
	o.Address = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *IdentityCredentialsCodeAddress) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialsCodeAddress) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *IdentityCredentialsCodeAddress) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *IdentityCredentialsCodeAddress) SetChannel(v string) {
	o.Channel = &v
}

func (o IdentityCredentialsCodeAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCredentialsCodeAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityCredentialsCodeAddress) UnmarshalJSON(data []byte) (err error) {
	varIdentityCredentialsCodeAddress := _IdentityCredentialsCodeAddress{}

	err = json.Unmarshal(data, &varIdentityCredentialsCodeAddress)

	if err != nil {
		return err
	}

	*o = IdentityCredentialsCodeAddress(varIdentityCredentialsCodeAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "channel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityCredentialsCodeAddress struct {
	value *IdentityCredentialsCodeAddress
	isSet bool
}

func (v NullableIdentityCredentialsCodeAddress) Get() *IdentityCredentialsCodeAddress {
	return v.value
}

func (v *NullableIdentityCredentialsCodeAddress) Set(val *IdentityCredentialsCodeAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialsCodeAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialsCodeAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialsCodeAddress(val *IdentityCredentialsCodeAddress) *NullableIdentityCredentialsCodeAddress {
	return &NullableIdentityCredentialsCodeAddress{value: val, isSet: true}
}

func (v NullableIdentityCredentialsCodeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialsCodeAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


