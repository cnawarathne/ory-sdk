/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.39-alpha.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IntrospectedOAuth2Token Introspection contains an access token's session data as specified by [IETF RFC 7662](https://tools.ietf.org/html/rfc7662)
type IntrospectedOAuth2Token struct {
	// Active is a boolean indicator of whether or not the presented token is currently active.  The specifics of a token's \"active\" state will vary depending on the implementation of the authorization server and the information it keeps about its tokens, but a \"true\" value return for the \"active\" property will generally indicate that a given token has been issued by this authorization server, has not been revoked by the resource owner, and is within its given time window of validity (e.g., after its issuance time and before its expiration time).
	Active bool `json:"active"`
	// Audience contains a list of the token's intended audiences.
	Aud []string `json:"aud,omitempty"`
	// ID is aclient identifier for the OAuth 2.0 client that requested this token.
	ClientId *string `json:"client_id,omitempty"`
	// Expires at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp *int64 `json:"exp,omitempty"`
	// Extra is arbitrary data set by the session.
	Ext map[string]interface{} `json:"ext,omitempty"`
	// Issued at is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token was originally issued.
	Iat *int64 `json:"iat,omitempty"`
	// IssuerURL is a string representing the issuer of this token
	Iss *string `json:"iss,omitempty"`
	// NotBefore is an integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this token is not to be used before.
	Nbf *int64 `json:"nbf,omitempty"`
	// ObfuscatedSubject is set when the subject identifier algorithm was set to \"pairwise\" during authorization. It is the `sub` value of the ID Token that was issued.
	ObfuscatedSubject *string `json:"obfuscated_subject,omitempty"`
	// Scope is a JSON string containing a space-separated list of scopes associated with this token.
	Scope *string `json:"scope,omitempty"`
	// Subject of the token, as defined in JWT [RFC7519]. Usually a machine-readable identifier of the resource owner who authorized this token.
	Sub *string `json:"sub,omitempty"`
	// TokenType is the introspected token's type, typically `Bearer`.
	TokenType *string `json:"token_type,omitempty"`
	// TokenUse is the introspected token's use, for example `access_token` or `refresh_token`.
	TokenUse *string `json:"token_use,omitempty"`
	// Username is a human-readable identifier for the resource owner who authorized this token.
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntrospectedOAuth2Token IntrospectedOAuth2Token

// NewIntrospectedOAuth2Token instantiates a new IntrospectedOAuth2Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntrospectedOAuth2Token(active bool) *IntrospectedOAuth2Token {
	this := IntrospectedOAuth2Token{}
	this.Active = active
	return &this
}

// NewIntrospectedOAuth2TokenWithDefaults instantiates a new IntrospectedOAuth2Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntrospectedOAuth2TokenWithDefaults() *IntrospectedOAuth2Token {
	this := IntrospectedOAuth2Token{}
	return &this
}

// GetActive returns the Active field value
func (o *IntrospectedOAuth2Token) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *IntrospectedOAuth2Token) SetActive(v bool) {
	o.Active = v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetAud() []string {
	if o == nil || o.Aud == nil {
		var ret []string
		return ret
	}
	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetAudOk() ([]string, bool) {
	if o == nil || o.Aud == nil {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasAud() bool {
	if o != nil && o.Aud != nil {
		return true
	}

	return false
}

// SetAud gets a reference to the given []string and assigns it to the Aud field.
func (o *IntrospectedOAuth2Token) SetAud(v []string) {
	o.Aud = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IntrospectedOAuth2Token) SetClientId(v string) {
	o.ClientId = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetExp() int64 {
	if o == nil || o.Exp == nil {
		var ret int64
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetExpOk() (*int64, bool) {
	if o == nil || o.Exp == nil {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasExp() bool {
	if o != nil && o.Exp != nil {
		return true
	}

	return false
}

// SetExp gets a reference to the given int64 and assigns it to the Exp field.
func (o *IntrospectedOAuth2Token) SetExp(v int64) {
	o.Exp = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetExt() map[string]interface{} {
	if o == nil || o.Ext == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetExtOk() (map[string]interface{}, bool) {
	if o == nil || o.Ext == nil {
		return nil, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasExt() bool {
	if o != nil && o.Ext != nil {
		return true
	}

	return false
}

// SetExt gets a reference to the given map[string]interface{} and assigns it to the Ext field.
func (o *IntrospectedOAuth2Token) SetExt(v map[string]interface{}) {
	o.Ext = v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetIat() int64 {
	if o == nil || o.Iat == nil {
		var ret int64
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetIatOk() (*int64, bool) {
	if o == nil || o.Iat == nil {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasIat() bool {
	if o != nil && o.Iat != nil {
		return true
	}

	return false
}

// SetIat gets a reference to the given int64 and assigns it to the Iat field.
func (o *IntrospectedOAuth2Token) SetIat(v int64) {
	o.Iat = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *IntrospectedOAuth2Token) SetIss(v string) {
	o.Iss = &v
}

// GetNbf returns the Nbf field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetNbf() int64 {
	if o == nil || o.Nbf == nil {
		var ret int64
		return ret
	}
	return *o.Nbf
}

// GetNbfOk returns a tuple with the Nbf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetNbfOk() (*int64, bool) {
	if o == nil || o.Nbf == nil {
		return nil, false
	}
	return o.Nbf, true
}

// HasNbf returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasNbf() bool {
	if o != nil && o.Nbf != nil {
		return true
	}

	return false
}

// SetNbf gets a reference to the given int64 and assigns it to the Nbf field.
func (o *IntrospectedOAuth2Token) SetNbf(v int64) {
	o.Nbf = &v
}

// GetObfuscatedSubject returns the ObfuscatedSubject field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetObfuscatedSubject() string {
	if o == nil || o.ObfuscatedSubject == nil {
		var ret string
		return ret
	}
	return *o.ObfuscatedSubject
}

// GetObfuscatedSubjectOk returns a tuple with the ObfuscatedSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetObfuscatedSubjectOk() (*string, bool) {
	if o == nil || o.ObfuscatedSubject == nil {
		return nil, false
	}
	return o.ObfuscatedSubject, true
}

// HasObfuscatedSubject returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasObfuscatedSubject() bool {
	if o != nil && o.ObfuscatedSubject != nil {
		return true
	}

	return false
}

// SetObfuscatedSubject gets a reference to the given string and assigns it to the ObfuscatedSubject field.
func (o *IntrospectedOAuth2Token) SetObfuscatedSubject(v string) {
	o.ObfuscatedSubject = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *IntrospectedOAuth2Token) SetScope(v string) {
	o.Scope = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetSub() string {
	if o == nil || o.Sub == nil {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetSubOk() (*string, bool) {
	if o == nil || o.Sub == nil {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasSub() bool {
	if o != nil && o.Sub != nil {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *IntrospectedOAuth2Token) SetSub(v string) {
	o.Sub = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *IntrospectedOAuth2Token) SetTokenType(v string) {
	o.TokenType = &v
}

// GetTokenUse returns the TokenUse field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetTokenUse() string {
	if o == nil || o.TokenUse == nil {
		var ret string
		return ret
	}
	return *o.TokenUse
}

// GetTokenUseOk returns a tuple with the TokenUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetTokenUseOk() (*string, bool) {
	if o == nil || o.TokenUse == nil {
		return nil, false
	}
	return o.TokenUse, true
}

// HasTokenUse returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasTokenUse() bool {
	if o != nil && o.TokenUse != nil {
		return true
	}

	return false
}

// SetTokenUse gets a reference to the given string and assigns it to the TokenUse field.
func (o *IntrospectedOAuth2Token) SetTokenUse(v string) {
	o.TokenUse = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IntrospectedOAuth2Token) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntrospectedOAuth2Token) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IntrospectedOAuth2Token) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IntrospectedOAuth2Token) SetUsername(v string) {
	o.Username = &v
}

func (o IntrospectedOAuth2Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	if o.Ext != nil {
		toSerialize["ext"] = o.Ext
	}
	if o.Iat != nil {
		toSerialize["iat"] = o.Iat
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.Nbf != nil {
		toSerialize["nbf"] = o.Nbf
	}
	if o.ObfuscatedSubject != nil {
		toSerialize["obfuscated_subject"] = o.ObfuscatedSubject
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.TokenUse != nil {
		toSerialize["token_use"] = o.TokenUse
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IntrospectedOAuth2Token) UnmarshalJSON(bytes []byte) (err error) {
	varIntrospectedOAuth2Token := _IntrospectedOAuth2Token{}

	if err = json.Unmarshal(bytes, &varIntrospectedOAuth2Token); err == nil {
		*o = IntrospectedOAuth2Token(varIntrospectedOAuth2Token)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "aud")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "exp")
		delete(additionalProperties, "ext")
		delete(additionalProperties, "iat")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "nbf")
		delete(additionalProperties, "obfuscated_subject")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "sub")
		delete(additionalProperties, "token_type")
		delete(additionalProperties, "token_use")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntrospectedOAuth2Token struct {
	value *IntrospectedOAuth2Token
	isSet bool
}

func (v NullableIntrospectedOAuth2Token) Get() *IntrospectedOAuth2Token {
	return v.value
}

func (v *NullableIntrospectedOAuth2Token) Set(val *IntrospectedOAuth2Token) {
	v.value = val
	v.isSet = true
}

func (v NullableIntrospectedOAuth2Token) IsSet() bool {
	return v.isSet
}

func (v *NullableIntrospectedOAuth2Token) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntrospectedOAuth2Token(val *IntrospectedOAuth2Token) *NullableIntrospectedOAuth2Token {
	return &NullableIntrospectedOAuth2Token{value: val, isSet: true}
}

func (v NullableIntrospectedOAuth2Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntrospectedOAuth2Token) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


