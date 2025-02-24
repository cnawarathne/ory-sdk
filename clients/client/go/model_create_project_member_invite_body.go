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
)

// checks if the CreateProjectMemberInviteBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectMemberInviteBody{}

// CreateProjectMemberInviteBody Create Project MemberInvite Request Body
type CreateProjectMemberInviteBody struct {
	// A email to invite
	InviteeEmail *string `json:"invitee_email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateProjectMemberInviteBody CreateProjectMemberInviteBody

// NewCreateProjectMemberInviteBody instantiates a new CreateProjectMemberInviteBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectMemberInviteBody() *CreateProjectMemberInviteBody {
	this := CreateProjectMemberInviteBody{}
	return &this
}

// NewCreateProjectMemberInviteBodyWithDefaults instantiates a new CreateProjectMemberInviteBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectMemberInviteBodyWithDefaults() *CreateProjectMemberInviteBody {
	this := CreateProjectMemberInviteBody{}
	return &this
}

// GetInviteeEmail returns the InviteeEmail field value if set, zero value otherwise.
func (o *CreateProjectMemberInviteBody) GetInviteeEmail() string {
	if o == nil || IsNil(o.InviteeEmail) {
		var ret string
		return ret
	}
	return *o.InviteeEmail
}

// GetInviteeEmailOk returns a tuple with the InviteeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectMemberInviteBody) GetInviteeEmailOk() (*string, bool) {
	if o == nil || IsNil(o.InviteeEmail) {
		return nil, false
	}
	return o.InviteeEmail, true
}

// HasInviteeEmail returns a boolean if a field has been set.
func (o *CreateProjectMemberInviteBody) HasInviteeEmail() bool {
	if o != nil && !IsNil(o.InviteeEmail) {
		return true
	}

	return false
}

// SetInviteeEmail gets a reference to the given string and assigns it to the InviteeEmail field.
func (o *CreateProjectMemberInviteBody) SetInviteeEmail(v string) {
	o.InviteeEmail = &v
}

func (o CreateProjectMemberInviteBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectMemberInviteBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InviteeEmail) {
		toSerialize["invitee_email"] = o.InviteeEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProjectMemberInviteBody) UnmarshalJSON(data []byte) (err error) {
	varCreateProjectMemberInviteBody := _CreateProjectMemberInviteBody{}

	err = json.Unmarshal(data, &varCreateProjectMemberInviteBody)

	if err != nil {
		return err
	}

	*o = CreateProjectMemberInviteBody(varCreateProjectMemberInviteBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invitee_email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProjectMemberInviteBody struct {
	value *CreateProjectMemberInviteBody
	isSet bool
}

func (v NullableCreateProjectMemberInviteBody) Get() *CreateProjectMemberInviteBody {
	return v.value
}

func (v *NullableCreateProjectMemberInviteBody) Set(val *CreateProjectMemberInviteBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectMemberInviteBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectMemberInviteBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectMemberInviteBody(val *CreateProjectMemberInviteBody) *NullableCreateProjectMemberInviteBody {
	return &NullableCreateProjectMemberInviteBody{value: val, isSet: true}
}

func (v NullableCreateProjectMemberInviteBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectMemberInviteBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


