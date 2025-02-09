/*
Ory Hydra API

Documentation for all of Ory Hydra's APIs. 

API version: v2.4.0-alpha.1
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the TrustedOAuth2JwtGrantJsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustedOAuth2JwtGrantJsonWebKey{}

// TrustedOAuth2JwtGrantJsonWebKey OAuth2 JWT Bearer Grant Type Issuer Trusted JSON Web Key
type TrustedOAuth2JwtGrantJsonWebKey struct {
	// The \"key_id\" is key unique identifier (same as kid header in jws/jwt).
	Kid *string `json:"kid,omitempty"`
	// The \"set\" is basically a name for a group(set) of keys. Will be the same as \"issuer\" in grant.
	Set *string `json:"set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrustedOAuth2JwtGrantJsonWebKey TrustedOAuth2JwtGrantJsonWebKey

// NewTrustedOAuth2JwtGrantJsonWebKey instantiates a new TrustedOAuth2JwtGrantJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedOAuth2JwtGrantJsonWebKey() *TrustedOAuth2JwtGrantJsonWebKey {
	this := TrustedOAuth2JwtGrantJsonWebKey{}
	return &this
}

// NewTrustedOAuth2JwtGrantJsonWebKeyWithDefaults instantiates a new TrustedOAuth2JwtGrantJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedOAuth2JwtGrantJsonWebKeyWithDefaults() *TrustedOAuth2JwtGrantJsonWebKey {
	this := TrustedOAuth2JwtGrantJsonWebKey{}
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantJsonWebKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantJsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantJsonWebKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *TrustedOAuth2JwtGrantJsonWebKey) SetKid(v string) {
	o.Kid = &v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *TrustedOAuth2JwtGrantJsonWebKey) GetSet() string {
	if o == nil || IsNil(o.Set) {
		var ret string
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOAuth2JwtGrantJsonWebKey) GetSetOk() (*string, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *TrustedOAuth2JwtGrantJsonWebKey) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given string and assigns it to the Set field.
func (o *TrustedOAuth2JwtGrantJsonWebKey) SetSet(v string) {
	o.Set = &v
}

func (o TrustedOAuth2JwtGrantJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustedOAuth2JwtGrantJsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrustedOAuth2JwtGrantJsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varTrustedOAuth2JwtGrantJsonWebKey := _TrustedOAuth2JwtGrantJsonWebKey{}

	err = json.Unmarshal(data, &varTrustedOAuth2JwtGrantJsonWebKey)

	if err != nil {
		return err
	}

	*o = TrustedOAuth2JwtGrantJsonWebKey(varTrustedOAuth2JwtGrantJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kid")
		delete(additionalProperties, "set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrustedOAuth2JwtGrantJsonWebKey struct {
	value *TrustedOAuth2JwtGrantJsonWebKey
	isSet bool
}

func (v NullableTrustedOAuth2JwtGrantJsonWebKey) Get() *TrustedOAuth2JwtGrantJsonWebKey {
	return v.value
}

func (v *NullableTrustedOAuth2JwtGrantJsonWebKey) Set(val *TrustedOAuth2JwtGrantJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedOAuth2JwtGrantJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedOAuth2JwtGrantJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedOAuth2JwtGrantJsonWebKey(val *TrustedOAuth2JwtGrantJsonWebKey) *NullableTrustedOAuth2JwtGrantJsonWebKey {
	return &NullableTrustedOAuth2JwtGrantJsonWebKey{value: val, isSet: true}
}

func (v NullableTrustedOAuth2JwtGrantJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedOAuth2JwtGrantJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


