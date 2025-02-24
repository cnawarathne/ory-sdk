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

// checks if the IdentityWithCredentialsOidcConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityWithCredentialsOidcConfig{}

// IdentityWithCredentialsOidcConfig struct for IdentityWithCredentialsOidcConfig
type IdentityWithCredentialsOidcConfig struct {
	Config *IdentityWithCredentialsPasswordConfig `json:"config,omitempty"`
	// A list of OpenID Connect Providers
	Providers []IdentityWithCredentialsOidcConfigProvider `json:"providers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityWithCredentialsOidcConfig IdentityWithCredentialsOidcConfig

// NewIdentityWithCredentialsOidcConfig instantiates a new IdentityWithCredentialsOidcConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsOidcConfig() *IdentityWithCredentialsOidcConfig {
	this := IdentityWithCredentialsOidcConfig{}
	return &this
}

// NewIdentityWithCredentialsOidcConfigWithDefaults instantiates a new IdentityWithCredentialsOidcConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsOidcConfigWithDefaults() *IdentityWithCredentialsOidcConfig {
	this := IdentityWithCredentialsOidcConfig{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentityWithCredentialsOidcConfig) GetConfig() IdentityWithCredentialsPasswordConfig {
	if o == nil || IsNil(o.Config) {
		var ret IdentityWithCredentialsPasswordConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsOidcConfig) GetConfigOk() (*IdentityWithCredentialsPasswordConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentityWithCredentialsOidcConfig) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentityWithCredentialsPasswordConfig and assigns it to the Config field.
func (o *IdentityWithCredentialsOidcConfig) SetConfig(v IdentityWithCredentialsPasswordConfig) {
	o.Config = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityWithCredentialsOidcConfig) GetProviders() []IdentityWithCredentialsOidcConfigProvider {
	if o == nil || IsNil(o.Providers) {
		var ret []IdentityWithCredentialsOidcConfigProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsOidcConfig) GetProvidersOk() ([]IdentityWithCredentialsOidcConfigProvider, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityWithCredentialsOidcConfig) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IdentityWithCredentialsOidcConfigProvider and assigns it to the Providers field.
func (o *IdentityWithCredentialsOidcConfig) SetProviders(v []IdentityWithCredentialsOidcConfigProvider) {
	o.Providers = v
}

func (o IdentityWithCredentialsOidcConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityWithCredentialsOidcConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityWithCredentialsOidcConfig) UnmarshalJSON(data []byte) (err error) {
	varIdentityWithCredentialsOidcConfig := _IdentityWithCredentialsOidcConfig{}

	err = json.Unmarshal(data, &varIdentityWithCredentialsOidcConfig)

	if err != nil {
		return err
	}

	*o = IdentityWithCredentialsOidcConfig(varIdentityWithCredentialsOidcConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "providers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityWithCredentialsOidcConfig struct {
	value *IdentityWithCredentialsOidcConfig
	isSet bool
}

func (v NullableIdentityWithCredentialsOidcConfig) Get() *IdentityWithCredentialsOidcConfig {
	return v.value
}

func (v *NullableIdentityWithCredentialsOidcConfig) Set(val *IdentityWithCredentialsOidcConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsOidcConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsOidcConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsOidcConfig(val *IdentityWithCredentialsOidcConfig) *NullableIdentityWithCredentialsOidcConfig {
	return &NullableIdentityWithCredentialsOidcConfig{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsOidcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsOidcConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


