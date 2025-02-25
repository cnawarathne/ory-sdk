/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.3.8
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ContinueWithRecoveryUiFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueWithRecoveryUiFlow{}

// ContinueWithRecoveryUiFlow struct for ContinueWithRecoveryUiFlow
type ContinueWithRecoveryUiFlow struct {
	// The ID of the recovery flow
	Id string `json:"id"`
	// The URL of the recovery flow  If this value is set, redirect the user's browser to this URL. This value is typically unset for native clients / API flows.
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinueWithRecoveryUiFlow ContinueWithRecoveryUiFlow

// NewContinueWithRecoveryUiFlow instantiates a new ContinueWithRecoveryUiFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithRecoveryUiFlow(id string) *ContinueWithRecoveryUiFlow {
	this := ContinueWithRecoveryUiFlow{}
	this.Id = id
	return &this
}

// NewContinueWithRecoveryUiFlowWithDefaults instantiates a new ContinueWithRecoveryUiFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithRecoveryUiFlowWithDefaults() *ContinueWithRecoveryUiFlow {
	this := ContinueWithRecoveryUiFlow{}
	return &this
}

// GetId returns the Id field value
func (o *ContinueWithRecoveryUiFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContinueWithRecoveryUiFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContinueWithRecoveryUiFlow) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContinueWithRecoveryUiFlow) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinueWithRecoveryUiFlow) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContinueWithRecoveryUiFlow) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContinueWithRecoveryUiFlow) SetUrl(v string) {
	o.Url = &v
}

func (o ContinueWithRecoveryUiFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueWithRecoveryUiFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContinueWithRecoveryUiFlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varContinueWithRecoveryUiFlow := _ContinueWithRecoveryUiFlow{}

	err = json.Unmarshal(data, &varContinueWithRecoveryUiFlow)

	if err != nil {
		return err
	}

	*o = ContinueWithRecoveryUiFlow(varContinueWithRecoveryUiFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContinueWithRecoveryUiFlow struct {
	value *ContinueWithRecoveryUiFlow
	isSet bool
}

func (v NullableContinueWithRecoveryUiFlow) Get() *ContinueWithRecoveryUiFlow {
	return v.value
}

func (v *NullableContinueWithRecoveryUiFlow) Set(val *ContinueWithRecoveryUiFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithRecoveryUiFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithRecoveryUiFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithRecoveryUiFlow(val *ContinueWithRecoveryUiFlow) *NullableContinueWithRecoveryUiFlow {
	return &NullableContinueWithRecoveryUiFlow{value: val, isSet: true}
}

func (v NullableContinueWithRecoveryUiFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithRecoveryUiFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


