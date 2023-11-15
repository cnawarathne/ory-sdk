/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PostCheckPermissionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCheckPermissionBody{}

// PostCheckPermissionBody Check Permission using Post Request Body
type PostCheckPermissionBody struct {
	// Namespace to query
	Namespace *string `json:"namespace,omitempty"`
	// Object to query
	Object *string `json:"object,omitempty"`
	// Relation to query
	Relation *string `json:"relation,omitempty"`
	// SubjectID to query  Either SubjectSet or SubjectID can be provided.
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostCheckPermissionBody PostCheckPermissionBody

// NewPostCheckPermissionBody instantiates a new PostCheckPermissionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCheckPermissionBody() *PostCheckPermissionBody {
	this := PostCheckPermissionBody{}
	return &this
}

// NewPostCheckPermissionBodyWithDefaults instantiates a new PostCheckPermissionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCheckPermissionBodyWithDefaults() *PostCheckPermissionBody {
	this := PostCheckPermissionBody{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *PostCheckPermissionBody) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCheckPermissionBody) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *PostCheckPermissionBody) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *PostCheckPermissionBody) SetNamespace(v string) {
	o.Namespace = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *PostCheckPermissionBody) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCheckPermissionBody) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *PostCheckPermissionBody) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *PostCheckPermissionBody) SetObject(v string) {
	o.Object = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *PostCheckPermissionBody) GetRelation() string {
	if o == nil || IsNil(o.Relation) {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCheckPermissionBody) GetRelationOk() (*string, bool) {
	if o == nil || IsNil(o.Relation) {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *PostCheckPermissionBody) HasRelation() bool {
	if o != nil && !IsNil(o.Relation) {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *PostCheckPermissionBody) SetRelation(v string) {
	o.Relation = &v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *PostCheckPermissionBody) GetSubjectId() string {
	if o == nil || IsNil(o.SubjectId) {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCheckPermissionBody) GetSubjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectId) {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *PostCheckPermissionBody) HasSubjectId() bool {
	if o != nil && !IsNil(o.SubjectId) {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *PostCheckPermissionBody) SetSubjectId(v string) {
	o.SubjectId = &v
}

// GetSubjectSet returns the SubjectSet field value if set, zero value otherwise.
func (o *PostCheckPermissionBody) GetSubjectSet() SubjectSet {
	if o == nil || IsNil(o.SubjectSet) {
		var ret SubjectSet
		return ret
	}
	return *o.SubjectSet
}

// GetSubjectSetOk returns a tuple with the SubjectSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCheckPermissionBody) GetSubjectSetOk() (*SubjectSet, bool) {
	if o == nil || IsNil(o.SubjectSet) {
		return nil, false
	}
	return o.SubjectSet, true
}

// HasSubjectSet returns a boolean if a field has been set.
func (o *PostCheckPermissionBody) HasSubjectSet() bool {
	if o != nil && !IsNil(o.SubjectSet) {
		return true
	}

	return false
}

// SetSubjectSet gets a reference to the given SubjectSet and assigns it to the SubjectSet field.
func (o *PostCheckPermissionBody) SetSubjectSet(v SubjectSet) {
	o.SubjectSet = &v
}

func (o PostCheckPermissionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCheckPermissionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Relation) {
		toSerialize["relation"] = o.Relation
	}
	if !IsNil(o.SubjectId) {
		toSerialize["subject_id"] = o.SubjectId
	}
	if !IsNil(o.SubjectSet) {
		toSerialize["subject_set"] = o.SubjectSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostCheckPermissionBody) UnmarshalJSON(bytes []byte) (err error) {
	varPostCheckPermissionBody := _PostCheckPermissionBody{}

	err = json.Unmarshal(bytes, &varPostCheckPermissionBody)

	if err != nil {
		return err
	}

	*o = PostCheckPermissionBody(varPostCheckPermissionBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "object")
		delete(additionalProperties, "relation")
		delete(additionalProperties, "subject_id")
		delete(additionalProperties, "subject_set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostCheckPermissionBody struct {
	value *PostCheckPermissionBody
	isSet bool
}

func (v NullablePostCheckPermissionBody) Get() *PostCheckPermissionBody {
	return v.value
}

func (v *NullablePostCheckPermissionBody) Set(val *PostCheckPermissionBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCheckPermissionBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCheckPermissionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCheckPermissionBody(val *PostCheckPermissionBody) *NullablePostCheckPermissionBody {
	return &NullablePostCheckPermissionBody{value: val, isSet: true}
}

func (v NullablePostCheckPermissionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCheckPermissionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


