/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.9
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateSubscriptionBody Update Subscription Request Body
type UpdateSubscriptionBody struct {
	//  monthly Monthly yearly Yearly
	Interval string `json:"interval"`
	Plan string `json:"plan"`
	ReturnTo *string `json:"return_to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSubscriptionBody UpdateSubscriptionBody

// NewUpdateSubscriptionBody instantiates a new UpdateSubscriptionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionBody(interval string, plan string) *UpdateSubscriptionBody {
	this := UpdateSubscriptionBody{}
	this.Interval = interval
	this.Plan = plan
	return &this
}

// NewUpdateSubscriptionBodyWithDefaults instantiates a new UpdateSubscriptionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionBodyWithDefaults() *UpdateSubscriptionBody {
	this := UpdateSubscriptionBody{}
	return &this
}

// GetInterval returns the Interval field value
func (o *UpdateSubscriptionBody) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionBody) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *UpdateSubscriptionBody) SetInterval(v string) {
	o.Interval = v
}

// GetPlan returns the Plan field value
func (o *UpdateSubscriptionBody) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionBody) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *UpdateSubscriptionBody) SetPlan(v string) {
	o.Plan = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *UpdateSubscriptionBody) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionBody) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *UpdateSubscriptionBody) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *UpdateSubscriptionBody) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o UpdateSubscriptionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateSubscriptionBody) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateSubscriptionBody := _UpdateSubscriptionBody{}

	if err = json.Unmarshal(bytes, &varUpdateSubscriptionBody); err == nil {
		*o = UpdateSubscriptionBody(varUpdateSubscriptionBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "interval")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "return_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSubscriptionBody struct {
	value *UpdateSubscriptionBody
	isSet bool
}

func (v NullableUpdateSubscriptionBody) Get() *UpdateSubscriptionBody {
	return v.value
}

func (v *NullableUpdateSubscriptionBody) Set(val *UpdateSubscriptionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionBody(val *UpdateSubscriptionBody) *NullableUpdateSubscriptionBody {
	return &NullableUpdateSubscriptionBody{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


