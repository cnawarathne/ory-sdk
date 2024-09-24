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

// checks if the ListEventStreams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEventStreams{}

// ListEventStreams Event Stream List
type ListEventStreams struct {
	EventStreams []EventStream `json:"event_streams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListEventStreams ListEventStreams

// NewListEventStreams instantiates a new ListEventStreams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEventStreams() *ListEventStreams {
	this := ListEventStreams{}
	return &this
}

// NewListEventStreamsWithDefaults instantiates a new ListEventStreams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEventStreamsWithDefaults() *ListEventStreams {
	this := ListEventStreams{}
	return &this
}

// GetEventStreams returns the EventStreams field value if set, zero value otherwise.
func (o *ListEventStreams) GetEventStreams() []EventStream {
	if o == nil || IsNil(o.EventStreams) {
		var ret []EventStream
		return ret
	}
	return o.EventStreams
}

// GetEventStreamsOk returns a tuple with the EventStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventStreams) GetEventStreamsOk() ([]EventStream, bool) {
	if o == nil || IsNil(o.EventStreams) {
		return nil, false
	}
	return o.EventStreams, true
}

// HasEventStreams returns a boolean if a field has been set.
func (o *ListEventStreams) HasEventStreams() bool {
	if o != nil && !IsNil(o.EventStreams) {
		return true
	}

	return false
}

// SetEventStreams gets a reference to the given []EventStream and assigns it to the EventStreams field.
func (o *ListEventStreams) SetEventStreams(v []EventStream) {
	o.EventStreams = v
}

func (o ListEventStreams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEventStreams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventStreams) {
		toSerialize["event_streams"] = o.EventStreams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListEventStreams) UnmarshalJSON(data []byte) (err error) {
	varListEventStreams := _ListEventStreams{}

	err = json.Unmarshal(data, &varListEventStreams)

	if err != nil {
		return err
	}

	*o = ListEventStreams(varListEventStreams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event_streams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListEventStreams struct {
	value *ListEventStreams
	isSet bool
}

func (v NullableListEventStreams) Get() *ListEventStreams {
	return v.value
}

func (v *NullableListEventStreams) Set(val *ListEventStreams) {
	v.value = val
	v.isSet = true
}

func (v NullableListEventStreams) IsSet() bool {
	return v.isSet
}

func (v *NullableListEventStreams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEventStreams(val *ListEventStreams) *NullableListEventStreams {
	return &NullableListEventStreams{value: val, isSet: true}
}

func (v NullableListEventStreams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEventStreams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


