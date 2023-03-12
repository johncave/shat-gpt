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

// checks if the DeleteFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFileResponse{}

// DeleteFileResponse struct for DeleteFileResponse
type DeleteFileResponse struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Deleted bool `json:"deleted"`
}

// NewDeleteFileResponse instantiates a new DeleteFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFileResponse(id string, object string, deleted bool) *DeleteFileResponse {
	this := DeleteFileResponse{}
	this.Id = id
	this.Object = object
	this.Deleted = deleted
	return &this
}

// NewDeleteFileResponseWithDefaults instantiates a new DeleteFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFileResponseWithDefaults() *DeleteFileResponse {
	this := DeleteFileResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeleteFileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteFileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteFileResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *DeleteFileResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DeleteFileResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DeleteFileResponse) SetObject(v string) {
	o.Object = v
}

// GetDeleted returns the Deleted field value
func (o *DeleteFileResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteFileResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteFileResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o DeleteFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

type NullableDeleteFileResponse struct {
	value *DeleteFileResponse
	isSet bool
}

func (v NullableDeleteFileResponse) Get() *DeleteFileResponse {
	return v.value
}

func (v *NullableDeleteFileResponse) Set(val *DeleteFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFileResponse(val *DeleteFileResponse) *NullableDeleteFileResponse {
	return &NullableDeleteFileResponse{value: val, isSet: true}
}

func (v NullableDeleteFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

