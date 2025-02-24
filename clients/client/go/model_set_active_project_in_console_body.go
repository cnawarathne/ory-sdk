/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.17.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the SetActiveProjectInConsoleBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetActiveProjectInConsoleBody{}

// SetActiveProjectInConsoleBody Set active project in the Ory Network Console Request Body
type SetActiveProjectInConsoleBody struct {
	// Project ID  The Project ID you want to set active.  format: uuid
	ProjectId string `json:"project_id"`
	AdditionalProperties map[string]interface{}
}

type _SetActiveProjectInConsoleBody SetActiveProjectInConsoleBody

// NewSetActiveProjectInConsoleBody instantiates a new SetActiveProjectInConsoleBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetActiveProjectInConsoleBody(projectId string) *SetActiveProjectInConsoleBody {
	this := SetActiveProjectInConsoleBody{}
	this.ProjectId = projectId
	return &this
}

// NewSetActiveProjectInConsoleBodyWithDefaults instantiates a new SetActiveProjectInConsoleBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetActiveProjectInConsoleBodyWithDefaults() *SetActiveProjectInConsoleBody {
	this := SetActiveProjectInConsoleBody{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *SetActiveProjectInConsoleBody) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SetActiveProjectInConsoleBody) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SetActiveProjectInConsoleBody) SetProjectId(v string) {
	o.ProjectId = v
}

func (o SetActiveProjectInConsoleBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetActiveProjectInConsoleBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetActiveProjectInConsoleBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varSetActiveProjectInConsoleBody := _SetActiveProjectInConsoleBody{}

	err = json.Unmarshal(data, &varSetActiveProjectInConsoleBody)

	if err != nil {
		return err
	}

	*o = SetActiveProjectInConsoleBody(varSetActiveProjectInConsoleBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetActiveProjectInConsoleBody struct {
	value *SetActiveProjectInConsoleBody
	isSet bool
}

func (v NullableSetActiveProjectInConsoleBody) Get() *SetActiveProjectInConsoleBody {
	return v.value
}

func (v *NullableSetActiveProjectInConsoleBody) Set(val *SetActiveProjectInConsoleBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetActiveProjectInConsoleBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetActiveProjectInConsoleBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetActiveProjectInConsoleBody(val *SetActiveProjectInConsoleBody) *NullableSetActiveProjectInConsoleBody {
	return &NullableSetActiveProjectInConsoleBody{value: val, isSet: true}
}

func (v NullableSetActiveProjectInConsoleBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetActiveProjectInConsoleBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


