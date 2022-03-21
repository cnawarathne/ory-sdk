/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.141
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceSettingsFlowWithTotpMethodBody struct for SubmitSelfServiceSettingsFlowWithTotpMethodBody
type SubmitSelfServiceSettingsFlowWithTotpMethodBody struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"totp\" when trying to add, update, or remove a totp pairing.
	Method string `json:"method"`
	// ValidationTOTP must contain a valid TOTP based on the
	TotpCode *string `json:"totp_code,omitempty"`
	// UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.
	TotpUnlink *bool `json:"totp_unlink,omitempty"`
}

// NewSubmitSelfServiceSettingsFlowWithTotpMethodBody instantiates a new SubmitSelfServiceSettingsFlowWithTotpMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithTotpMethodBody(method string) *SubmitSelfServiceSettingsFlowWithTotpMethodBody {
	this := SubmitSelfServiceSettingsFlowWithTotpMethodBody{}
	this.Method = method
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithTotpMethodBodyWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithTotpMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithTotpMethodBodyWithDefaults() *SubmitSelfServiceSettingsFlowWithTotpMethodBody {
	this := SubmitSelfServiceSettingsFlowWithTotpMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetTotpCode returns the TotpCode field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetTotpCode() string {
	if o == nil || o.TotpCode == nil {
		var ret string
		return ret
	}
	return *o.TotpCode
}

// GetTotpCodeOk returns a tuple with the TotpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetTotpCodeOk() (*string, bool) {
	if o == nil || o.TotpCode == nil {
		return nil, false
	}
	return o.TotpCode, true
}

// HasTotpCode returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) HasTotpCode() bool {
	if o != nil && o.TotpCode != nil {
		return true
	}

	return false
}

// SetTotpCode gets a reference to the given string and assigns it to the TotpCode field.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) SetTotpCode(v string) {
	o.TotpCode = &v
}

// GetTotpUnlink returns the TotpUnlink field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetTotpUnlink() bool {
	if o == nil || o.TotpUnlink == nil {
		var ret bool
		return ret
	}
	return *o.TotpUnlink
}

// GetTotpUnlinkOk returns a tuple with the TotpUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) GetTotpUnlinkOk() (*bool, bool) {
	if o == nil || o.TotpUnlink == nil {
		return nil, false
	}
	return o.TotpUnlink, true
}

// HasTotpUnlink returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) HasTotpUnlink() bool {
	if o != nil && o.TotpUnlink != nil {
		return true
	}

	return false
}

// SetTotpUnlink gets a reference to the given bool and assigns it to the TotpUnlink field.
func (o *SubmitSelfServiceSettingsFlowWithTotpMethodBody) SetTotpUnlink(v bool) {
	o.TotpUnlink = &v
}

func (o SubmitSelfServiceSettingsFlowWithTotpMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.TotpCode != nil {
		toSerialize["totp_code"] = o.TotpCode
	}
	if o.TotpUnlink != nil {
		toSerialize["totp_unlink"] = o.TotpUnlink
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody struct {
	value *SubmitSelfServiceSettingsFlowWithTotpMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) Get() *SubmitSelfServiceSettingsFlowWithTotpMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) Set(val *SubmitSelfServiceSettingsFlowWithTotpMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithTotpMethodBody(val *SubmitSelfServiceSettingsFlowWithTotpMethodBody) *NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody {
	return &NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTotpMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


