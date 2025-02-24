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

// checks if the ListWorkspaceProjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWorkspaceProjects{}

// ListWorkspaceProjects struct for ListWorkspaceProjects
type ListWorkspaceProjects struct {
	HasNextPage bool `json:"has_next_page"`
	NextPage string `json:"next_page"`
	Projects []ProjectMetadata `json:"projects"`
	AdditionalProperties map[string]interface{}
}

type _ListWorkspaceProjects ListWorkspaceProjects

// NewListWorkspaceProjects instantiates a new ListWorkspaceProjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceProjects(hasNextPage bool, nextPage string, projects []ProjectMetadata) *ListWorkspaceProjects {
	this := ListWorkspaceProjects{}
	this.HasNextPage = hasNextPage
	this.NextPage = nextPage
	this.Projects = projects
	return &this
}

// NewListWorkspaceProjectsWithDefaults instantiates a new ListWorkspaceProjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceProjectsWithDefaults() *ListWorkspaceProjects {
	this := ListWorkspaceProjects{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *ListWorkspaceProjects) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaceProjects) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *ListWorkspaceProjects) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetNextPage returns the NextPage field value
func (o *ListWorkspaceProjects) GetNextPage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaceProjects) GetNextPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPage, true
}

// SetNextPage sets field value
func (o *ListWorkspaceProjects) SetNextPage(v string) {
	o.NextPage = v
}

// GetProjects returns the Projects field value
func (o *ListWorkspaceProjects) GetProjects() []ProjectMetadata {
	if o == nil {
		var ret []ProjectMetadata
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ListWorkspaceProjects) GetProjectsOk() ([]ProjectMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ListWorkspaceProjects) SetProjects(v []ProjectMetadata) {
	o.Projects = v
}

func (o ListWorkspaceProjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWorkspaceProjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_next_page"] = o.HasNextPage
	toSerialize["next_page"] = o.NextPage
	toSerialize["projects"] = o.Projects

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWorkspaceProjects) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_next_page",
		"next_page",
		"projects",
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

	varListWorkspaceProjects := _ListWorkspaceProjects{}

	err = json.Unmarshal(data, &varListWorkspaceProjects)

	if err != nil {
		return err
	}

	*o = ListWorkspaceProjects(varListWorkspaceProjects)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "has_next_page")
		delete(additionalProperties, "next_page")
		delete(additionalProperties, "projects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWorkspaceProjects struct {
	value *ListWorkspaceProjects
	isSet bool
}

func (v NullableListWorkspaceProjects) Get() *ListWorkspaceProjects {
	return v.value
}

func (v *NullableListWorkspaceProjects) Set(val *ListWorkspaceProjects) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceProjects) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceProjects(val *ListWorkspaceProjects) *NullableListWorkspaceProjects {
	return &NullableListWorkspaceProjects{value: val, isSet: true}
}

func (v NullableListWorkspaceProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


