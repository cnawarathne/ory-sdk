/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.17.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the RecoveryIdentityAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryIdentityAddress{}

// RecoveryIdentityAddress struct for RecoveryIdentityAddress
type RecoveryIdentityAddress struct {
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id string `json:"id"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Value string `json:"value"`
	Via string `json:"via"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryIdentityAddress RecoveryIdentityAddress

// NewRecoveryIdentityAddress instantiates a new RecoveryIdentityAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryIdentityAddress(id string, value string, via string) *RecoveryIdentityAddress {
	this := RecoveryIdentityAddress{}
	this.Id = id
	this.Value = value
	this.Via = via
	return &this
}

// NewRecoveryIdentityAddressWithDefaults instantiates a new RecoveryIdentityAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryIdentityAddressWithDefaults() *RecoveryIdentityAddress {
	this := RecoveryIdentityAddress{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RecoveryIdentityAddress) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryIdentityAddress) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RecoveryIdentityAddress) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RecoveryIdentityAddress) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *RecoveryIdentityAddress) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecoveryIdentityAddress) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecoveryIdentityAddress) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RecoveryIdentityAddress) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryIdentityAddress) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RecoveryIdentityAddress) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RecoveryIdentityAddress) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value
func (o *RecoveryIdentityAddress) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RecoveryIdentityAddress) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RecoveryIdentityAddress) SetValue(v string) {
	o.Value = v
}

// GetVia returns the Via field value
func (o *RecoveryIdentityAddress) GetVia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Via
}

// GetViaOk returns a tuple with the Via field value
// and a boolean to check if the value has been set.
func (o *RecoveryIdentityAddress) GetViaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Via, true
}

// SetVia sets field value
func (o *RecoveryIdentityAddress) SetVia(v string) {
	o.Via = v
}

func (o RecoveryIdentityAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryIdentityAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["value"] = o.Value
	toSerialize["via"] = o.Via

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryIdentityAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"value",
		"via",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRecoveryIdentityAddress := _RecoveryIdentityAddress{}

	err = json.Unmarshal(data, &varRecoveryIdentityAddress)

	if err != nil {
		return err
	}

	*o = RecoveryIdentityAddress(varRecoveryIdentityAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "value")
		delete(additionalProperties, "via")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryIdentityAddress struct {
	value *RecoveryIdentityAddress
	isSet bool
}

func (v NullableRecoveryIdentityAddress) Get() *RecoveryIdentityAddress {
	return v.value
}

func (v *NullableRecoveryIdentityAddress) Set(val *RecoveryIdentityAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryIdentityAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryIdentityAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryIdentityAddress(val *RecoveryIdentityAddress) *NullableRecoveryIdentityAddress {
	return &NullableRecoveryIdentityAddress{value: val, isSet: true}
}

func (v NullableRecoveryIdentityAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryIdentityAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


