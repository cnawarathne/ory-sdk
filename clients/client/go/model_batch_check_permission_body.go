/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the BatchCheckPermissionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchCheckPermissionBody{}

// BatchCheckPermissionBody Batch Check Permission Body
type BatchCheckPermissionBody struct {
	Tuples []Relationship `json:"tuples,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchCheckPermissionBody BatchCheckPermissionBody

// NewBatchCheckPermissionBody instantiates a new BatchCheckPermissionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCheckPermissionBody() *BatchCheckPermissionBody {
	this := BatchCheckPermissionBody{}
	return &this
}

// NewBatchCheckPermissionBodyWithDefaults instantiates a new BatchCheckPermissionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCheckPermissionBodyWithDefaults() *BatchCheckPermissionBody {
	this := BatchCheckPermissionBody{}
	return &this
}

// GetTuples returns the Tuples field value if set, zero value otherwise.
func (o *BatchCheckPermissionBody) GetTuples() []Relationship {
	if o == nil || IsNil(o.Tuples) {
		var ret []Relationship
		return ret
	}
	return o.Tuples
}

// GetTuplesOk returns a tuple with the Tuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCheckPermissionBody) GetTuplesOk() ([]Relationship, bool) {
	if o == nil || IsNil(o.Tuples) {
		return nil, false
	}
	return o.Tuples, true
}

// HasTuples returns a boolean if a field has been set.
func (o *BatchCheckPermissionBody) HasTuples() bool {
	if o != nil && !IsNil(o.Tuples) {
		return true
	}

	return false
}

// SetTuples gets a reference to the given []Relationship and assigns it to the Tuples field.
func (o *BatchCheckPermissionBody) SetTuples(v []Relationship) {
	o.Tuples = v
}

func (o BatchCheckPermissionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchCheckPermissionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tuples) {
		toSerialize["tuples"] = o.Tuples
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchCheckPermissionBody) UnmarshalJSON(data []byte) (err error) {
	varBatchCheckPermissionBody := _BatchCheckPermissionBody{}

	err = json.Unmarshal(data, &varBatchCheckPermissionBody)

	if err != nil {
		return err
	}

	*o = BatchCheckPermissionBody(varBatchCheckPermissionBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tuples")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchCheckPermissionBody struct {
	value *BatchCheckPermissionBody
	isSet bool
}

func (v NullableBatchCheckPermissionBody) Get() *BatchCheckPermissionBody {
	return v.value
}

func (v *NullableBatchCheckPermissionBody) Set(val *BatchCheckPermissionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCheckPermissionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCheckPermissionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCheckPermissionBody(val *BatchCheckPermissionBody) *NullableBatchCheckPermissionBody {
	return &NullableBatchCheckPermissionBody{value: val, isSet: true}
}

func (v NullableBatchCheckPermissionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCheckPermissionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


