/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model{}

// Model struct for Model
type Model struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Created int32 `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel(id string, object string, created int32, ownedBy string) *Model {
	this := Model{}
	this.Id = id
	this.Object = object
	this.Created = created
	this.OwnedBy = ownedBy
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetId returns the Id field value
func (o *Model) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Model) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Model) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Model) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Model) SetObject(v string) {
	o.Object = v
}

// GetCreated returns the Created field value
func (o *Model) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Model) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Model) SetCreated(v int32) {
	o.Created = v
}

// GetOwnedBy returns the OwnedBy field value
func (o *Model) GetOwnedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value
// and a boolean to check if the value has been set.
func (o *Model) GetOwnedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedBy, true
}

// SetOwnedBy sets field value
func (o *Model) SetOwnedBy(v string) {
	o.OwnedBy = v
}

func (o Model) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created"] = o.Created
	toSerialize["owned_by"] = o.OwnedBy
	return toSerialize, nil
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


