/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// EmailTemplateData Contains the data of the email template, including the subject and body in HTML and plaintext variants
type EmailTemplateData struct {
	Body EmailTemplateDataBody `json:"body"`
	Subject string `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _EmailTemplateData EmailTemplateData

// NewEmailTemplateData instantiates a new EmailTemplateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateData(body EmailTemplateDataBody, subject string) *EmailTemplateData {
	this := EmailTemplateData{}
	this.Body = body
	this.Subject = subject
	return &this
}

// NewEmailTemplateDataWithDefaults instantiates a new EmailTemplateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateDataWithDefaults() *EmailTemplateData {
	this := EmailTemplateData{}
	return &this
}

// GetBody returns the Body field value
func (o *EmailTemplateData) GetBody() EmailTemplateDataBody {
	if o == nil {
		var ret EmailTemplateDataBody
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateData) GetBodyOk() (*EmailTemplateDataBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailTemplateData) SetBody(v EmailTemplateDataBody) {
	o.Body = v
}

// GetSubject returns the Subject field value
func (o *EmailTemplateData) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateData) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailTemplateData) SetSubject(v string) {
	o.Subject = v
}

func (o EmailTemplateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailTemplateData) UnmarshalJSON(bytes []byte) (err error) {
	varEmailTemplateData := _EmailTemplateData{}

	if err = json.Unmarshal(bytes, &varEmailTemplateData); err == nil {
		*o = EmailTemplateData(varEmailTemplateData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailTemplateData struct {
	value *EmailTemplateData
	isSet bool
}

func (v NullableEmailTemplateData) Get() *EmailTemplateData {
	return v.value
}

func (v *NullableEmailTemplateData) Set(val *EmailTemplateData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateData(val *EmailTemplateData) *NullableEmailTemplateData {
	return &NullableEmailTemplateData{value: val, isSet: true}
}

func (v NullableEmailTemplateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


