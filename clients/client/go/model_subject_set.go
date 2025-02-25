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
	"fmt"
)

// checks if the SubjectSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectSet{}

// SubjectSet struct for SubjectSet
type SubjectSet struct {
	// Namespace of the Subject Set
	Namespace string `json:"namespace"`
	// Object of the Subject Set
	Object string `json:"object"`
	// Relation of the Subject Set
	Relation string `json:"relation"`
	AdditionalProperties map[string]interface{}
}

type _SubjectSet SubjectSet

// NewSubjectSet instantiates a new SubjectSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectSet(namespace string, object string, relation string) *SubjectSet {
	this := SubjectSet{}
	this.Namespace = namespace
	this.Object = object
	this.Relation = relation
	return &this
}

// NewSubjectSetWithDefaults instantiates a new SubjectSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectSetWithDefaults() *SubjectSet {
	this := SubjectSet{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *SubjectSet) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *SubjectSet) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *SubjectSet) SetNamespace(v string) {
	o.Namespace = v
}

// GetObject returns the Object field value
func (o *SubjectSet) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SubjectSet) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SubjectSet) SetObject(v string) {
	o.Object = v
}

// GetRelation returns the Relation field value
func (o *SubjectSet) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *SubjectSet) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *SubjectSet) SetRelation(v string) {
	o.Relation = v
}

func (o SubjectSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["namespace"] = o.Namespace
	toSerialize["object"] = o.Object
	toSerialize["relation"] = o.Relation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubjectSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"namespace",
		"object",
		"relation",
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

	varSubjectSet := _SubjectSet{}

	err = json.Unmarshal(data, &varSubjectSet)

	if err != nil {
		return err
	}

	*o = SubjectSet(varSubjectSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "object")
		delete(additionalProperties, "relation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubjectSet struct {
	value *SubjectSet
	isSet bool
}

func (v NullableSubjectSet) Get() *SubjectSet {
	return v.value
}

func (v *NullableSubjectSet) Set(val *SubjectSet) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectSet) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectSet(val *SubjectSet) *NullableSubjectSet {
	return &NullableSubjectSet{value: val, isSet: true}
}

func (v NullableSubjectSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


