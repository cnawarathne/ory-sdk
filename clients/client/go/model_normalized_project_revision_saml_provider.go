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
	"time"
)

// checks if the NormalizedProjectRevisionSAMLProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NormalizedProjectRevisionSAMLProvider{}

// NormalizedProjectRevisionSAMLProvider struct for NormalizedProjectRevisionSAMLProvider
type NormalizedProjectRevisionSAMLProvider struct {
	// ClientID is the application's Client ID.
	ClientId *string `json:"client_id,omitempty"`
	ClientSecret NullableString `json:"client_secret,omitempty"`
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	// Label represents an optional label which can be used in the UI generation.
	Label *string `json:"label,omitempty"`
	// Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider's data (e.g. GitHub or Google profile information) to hydrate the identity's data.
	MapperUrl *string `json:"mapper_url,omitempty"`
	OrganizationId NullableString `json:"organization_id,omitempty"`
	// The Revision's ID this schema belongs to
	ProjectRevisionId *string `json:"project_revision_id,omitempty"`
	// ID is the provider's ID
	ProviderId *string `json:"provider_id,omitempty"`
	// RawIDPMetadataXML is the raw XML metadata of the IDP.
	RawIdpMetadataXml *string `json:"raw_idp_metadata_xml,omitempty"`
	// State indicates the state of the provider  Only providers with state `enabled` will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled
	State *string `json:"state,omitempty"`
	// Last Time Project's Revision was Updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NormalizedProjectRevisionSAMLProvider NormalizedProjectRevisionSAMLProvider

// NewNormalizedProjectRevisionSAMLProvider instantiates a new NormalizedProjectRevisionSAMLProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionSAMLProvider() *NormalizedProjectRevisionSAMLProvider {
	this := NormalizedProjectRevisionSAMLProvider{}
	return &this
}

// NewNormalizedProjectRevisionSAMLProviderWithDefaults instantiates a new NormalizedProjectRevisionSAMLProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionSAMLProviderWithDefaults() *NormalizedProjectRevisionSAMLProvider {
	this := NormalizedProjectRevisionSAMLProvider{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *NormalizedProjectRevisionSAMLProvider) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionSAMLProvider) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionSAMLProvider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *NormalizedProjectRevisionSAMLProvider) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *NormalizedProjectRevisionSAMLProvider) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *NormalizedProjectRevisionSAMLProvider) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionSAMLProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalizedProjectRevisionSAMLProvider) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NormalizedProjectRevisionSAMLProvider) SetLabel(v string) {
	o.Label = &v
}

// GetMapperUrl returns the MapperUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetMapperUrl() string {
	if o == nil || IsNil(o.MapperUrl) {
		var ret string
		return ret
	}
	return *o.MapperUrl
}

// GetMapperUrlOk returns a tuple with the MapperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetMapperUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MapperUrl) {
		return nil, false
	}
	return o.MapperUrl, true
}

// HasMapperUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasMapperUrl() bool {
	if o != nil && !IsNil(o.MapperUrl) {
		return true
	}

	return false
}

// SetMapperUrl gets a reference to the given string and assigns it to the MapperUrl field.
func (o *NormalizedProjectRevisionSAMLProvider) SetMapperUrl(v string) {
	o.MapperUrl = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionSAMLProvider) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionSAMLProvider) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableString and assigns it to the OrganizationId field.
func (o *NormalizedProjectRevisionSAMLProvider) SetOrganizationId(v string) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *NormalizedProjectRevisionSAMLProvider) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *NormalizedProjectRevisionSAMLProvider) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetProjectRevisionId() string {
	if o == nil || IsNil(o.ProjectRevisionId) {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectRevisionId) {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasProjectRevisionId() bool {
	if o != nil && !IsNil(o.ProjectRevisionId) {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *NormalizedProjectRevisionSAMLProvider) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *NormalizedProjectRevisionSAMLProvider) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetRawIdpMetadataXml returns the RawIdpMetadataXml field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetRawIdpMetadataXml() string {
	if o == nil || IsNil(o.RawIdpMetadataXml) {
		var ret string
		return ret
	}
	return *o.RawIdpMetadataXml
}

// GetRawIdpMetadataXmlOk returns a tuple with the RawIdpMetadataXml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetRawIdpMetadataXmlOk() (*string, bool) {
	if o == nil || IsNil(o.RawIdpMetadataXml) {
		return nil, false
	}
	return o.RawIdpMetadataXml, true
}

// HasRawIdpMetadataXml returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasRawIdpMetadataXml() bool {
	if o != nil && !IsNil(o.RawIdpMetadataXml) {
		return true
	}

	return false
}

// SetRawIdpMetadataXml gets a reference to the given string and assigns it to the RawIdpMetadataXml field.
func (o *NormalizedProjectRevisionSAMLProvider) SetRawIdpMetadataXml(v string) {
	o.RawIdpMetadataXml = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NormalizedProjectRevisionSAMLProvider) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionSAMLProvider) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionSAMLProvider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionSAMLProvider) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionSAMLProvider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NormalizedProjectRevisionSAMLProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NormalizedProjectRevisionSAMLProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.MapperUrl) {
		toSerialize["mapper_url"] = o.MapperUrl
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organization_id"] = o.OrganizationId.Get()
	}
	if !IsNil(o.ProjectRevisionId) {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if !IsNil(o.ProviderId) {
		toSerialize["provider_id"] = o.ProviderId
	}
	if !IsNil(o.RawIdpMetadataXml) {
		toSerialize["raw_idp_metadata_xml"] = o.RawIdpMetadataXml
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NormalizedProjectRevisionSAMLProvider) UnmarshalJSON(data []byte) (err error) {
	varNormalizedProjectRevisionSAMLProvider := _NormalizedProjectRevisionSAMLProvider{}

	err = json.Unmarshal(data, &varNormalizedProjectRevisionSAMLProvider)

	if err != nil {
		return err
	}

	*o = NormalizedProjectRevisionSAMLProvider(varNormalizedProjectRevisionSAMLProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "mapper_url")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "project_revision_id")
		delete(additionalProperties, "provider_id")
		delete(additionalProperties, "raw_idp_metadata_xml")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNormalizedProjectRevisionSAMLProvider struct {
	value *NormalizedProjectRevisionSAMLProvider
	isSet bool
}

func (v NullableNormalizedProjectRevisionSAMLProvider) Get() *NormalizedProjectRevisionSAMLProvider {
	return v.value
}

func (v *NullableNormalizedProjectRevisionSAMLProvider) Set(val *NormalizedProjectRevisionSAMLProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionSAMLProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionSAMLProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionSAMLProvider(val *NormalizedProjectRevisionSAMLProvider) *NullableNormalizedProjectRevisionSAMLProvider {
	return &NullableNormalizedProjectRevisionSAMLProvider{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionSAMLProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionSAMLProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


