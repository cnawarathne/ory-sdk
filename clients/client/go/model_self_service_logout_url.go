/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.124
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SelfServiceLogoutUrl struct for SelfServiceLogoutUrl
type SelfServiceLogoutUrl struct {
	// LogoutToken can be used to perform logout using AJAX.
	LogoutToken string `json:"logout_token"`
	// LogoutURL can be opened in a browser to sign the user out.  format: uri
	LogoutUrl string `json:"logout_url"`
}

// NewSelfServiceLogoutUrl instantiates a new SelfServiceLogoutUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceLogoutUrl(logoutToken string, logoutUrl string) *SelfServiceLogoutUrl {
	this := SelfServiceLogoutUrl{}
	this.LogoutToken = logoutToken
	this.LogoutUrl = logoutUrl
	return &this
}

// NewSelfServiceLogoutUrlWithDefaults instantiates a new SelfServiceLogoutUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceLogoutUrlWithDefaults() *SelfServiceLogoutUrl {
	this := SelfServiceLogoutUrl{}
	return &this
}

// GetLogoutToken returns the LogoutToken field value
func (o *SelfServiceLogoutUrl) GetLogoutToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutToken
}

// GetLogoutTokenOk returns a tuple with the LogoutToken field value
// and a boolean to check if the value has been set.
func (o *SelfServiceLogoutUrl) GetLogoutTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoutToken, true
}

// SetLogoutToken sets field value
func (o *SelfServiceLogoutUrl) SetLogoutToken(v string) {
	o.LogoutToken = v
}

// GetLogoutUrl returns the LogoutUrl field value
func (o *SelfServiceLogoutUrl) GetLogoutUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value
// and a boolean to check if the value has been set.
func (o *SelfServiceLogoutUrl) GetLogoutUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogoutUrl, true
}

// SetLogoutUrl sets field value
func (o *SelfServiceLogoutUrl) SetLogoutUrl(v string) {
	o.LogoutUrl = v
}

func (o SelfServiceLogoutUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logout_token"] = o.LogoutToken
	}
	if true {
		toSerialize["logout_url"] = o.LogoutUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceLogoutUrl struct {
	value *SelfServiceLogoutUrl
	isSet bool
}

func (v NullableSelfServiceLogoutUrl) Get() *SelfServiceLogoutUrl {
	return v.value
}

func (v *NullableSelfServiceLogoutUrl) Set(val *SelfServiceLogoutUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceLogoutUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceLogoutUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceLogoutUrl(val *SelfServiceLogoutUrl) *NullableSelfServiceLogoutUrl {
	return &NullableSelfServiceLogoutUrl{value: val, isSet: true}
}

func (v NullableSelfServiceLogoutUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceLogoutUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


