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
	"fmt"
)

// checks if the UpdateRegistrationFlowWithProfileMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRegistrationFlowWithProfileMethod{}

// UpdateRegistrationFlowWithProfileMethod Update Registration Flow with Profile Method
type UpdateRegistrationFlowWithProfileMethod struct {
	// The Anti-CSRF Token  This token is only required when performing browser flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to profile when trying to update a profile.
	Method string `json:"method"`
	// Screen requests navigation to a previous screen.  This must be set to credential-selection to go back to the credential selection screen. credential-selection RegistrationScreenCredentialSelection nolint:gosec // not a credential previous RegistrationScreenPrevious
	Screen *string `json:"screen,omitempty"`
	// Traits  The identity's traits.
	Traits map[string]interface{} `json:"traits"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRegistrationFlowWithProfileMethod UpdateRegistrationFlowWithProfileMethod

// NewUpdateRegistrationFlowWithProfileMethod instantiates a new UpdateRegistrationFlowWithProfileMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegistrationFlowWithProfileMethod(method string, traits map[string]interface{}) *UpdateRegistrationFlowWithProfileMethod {
	this := UpdateRegistrationFlowWithProfileMethod{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewUpdateRegistrationFlowWithProfileMethodWithDefaults instantiates a new UpdateRegistrationFlowWithProfileMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegistrationFlowWithProfileMethodWithDefaults() *UpdateRegistrationFlowWithProfileMethod {
	this := UpdateRegistrationFlowWithProfileMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithProfileMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRegistrationFlowWithProfileMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateRegistrationFlowWithProfileMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRegistrationFlowWithProfileMethod) SetMethod(v string) {
	o.Method = v
}

// GetScreen returns the Screen field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithProfileMethod) GetScreen() string {
	if o == nil || IsNil(o.Screen) {
		var ret string
		return ret
	}
	return *o.Screen
}

// GetScreenOk returns a tuple with the Screen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) GetScreenOk() (*string, bool) {
	if o == nil || IsNil(o.Screen) {
		return nil, false
	}
	return o.Screen, true
}

// HasScreen returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) HasScreen() bool {
	if o != nil && !IsNil(o.Screen) {
		return true
	}

	return false
}

// SetScreen gets a reference to the given string and assigns it to the Screen field.
func (o *UpdateRegistrationFlowWithProfileMethod) SetScreen(v string) {
	o.Screen = &v
}

// GetTraits returns the Traits field value
func (o *UpdateRegistrationFlowWithProfileMethod) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *UpdateRegistrationFlowWithProfileMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithProfileMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithProfileMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRegistrationFlowWithProfileMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateRegistrationFlowWithProfileMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRegistrationFlowWithProfileMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.Screen) {
		toSerialize["screen"] = o.Screen
	}
	toSerialize["traits"] = o.Traits
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRegistrationFlowWithProfileMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"traits",
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

	varUpdateRegistrationFlowWithProfileMethod := _UpdateRegistrationFlowWithProfileMethod{}

	err = json.Unmarshal(data, &varUpdateRegistrationFlowWithProfileMethod)

	if err != nil {
		return err
	}

	*o = UpdateRegistrationFlowWithProfileMethod(varUpdateRegistrationFlowWithProfileMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "screen")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRegistrationFlowWithProfileMethod struct {
	value *UpdateRegistrationFlowWithProfileMethod
	isSet bool
}

func (v NullableUpdateRegistrationFlowWithProfileMethod) Get() *UpdateRegistrationFlowWithProfileMethod {
	return v.value
}

func (v *NullableUpdateRegistrationFlowWithProfileMethod) Set(val *UpdateRegistrationFlowWithProfileMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowWithProfileMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowWithProfileMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowWithProfileMethod(val *UpdateRegistrationFlowWithProfileMethod) *NullableUpdateRegistrationFlowWithProfileMethod {
	return &NullableUpdateRegistrationFlowWithProfileMethod{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowWithProfileMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowWithProfileMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


