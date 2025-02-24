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
)

// checks if the UiNodeMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiNodeMeta{}

// UiNodeMeta This might include a label and other information that can optionally be used to render UIs.
type UiNodeMeta struct {
	Label *UiText `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiNodeMeta UiNodeMeta

// NewUiNodeMeta instantiates a new UiNodeMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeMeta() *UiNodeMeta {
	this := UiNodeMeta{}
	return &this
}

// NewUiNodeMetaWithDefaults instantiates a new UiNodeMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeMetaWithDefaults() *UiNodeMeta {
	this := UiNodeMeta{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UiNodeMeta) GetLabel() UiText {
	if o == nil || IsNil(o.Label) {
		var ret UiText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeMeta) GetLabelOk() (*UiText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UiNodeMeta) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given UiText and assigns it to the Label field.
func (o *UiNodeMeta) SetLabel(v UiText) {
	o.Label = &v
}

func (o UiNodeMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiNodeMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiNodeMeta) UnmarshalJSON(data []byte) (err error) {
	varUiNodeMeta := _UiNodeMeta{}

	err = json.Unmarshal(data, &varUiNodeMeta)

	if err != nil {
		return err
	}

	*o = UiNodeMeta(varUiNodeMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiNodeMeta struct {
	value *UiNodeMeta
	isSet bool
}

func (v NullableUiNodeMeta) Get() *UiNodeMeta {
	return v.value
}

func (v *NullableUiNodeMeta) Set(val *UiNodeMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeMeta(val *UiNodeMeta) *NullableUiNodeMeta {
	return &NullableUiNodeMeta{value: val, isSet: true}
}

func (v NullableUiNodeMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


