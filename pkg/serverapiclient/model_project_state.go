/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverapiclient

import (
	"encoding/json"
)

// checks if the ProjectState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectState{}

// ProjectState struct for ProjectState
type ProjectState struct {
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Uptime    *int32  `json:"uptime,omitempty"`
}

// NewProjectState instantiates a new ProjectState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectState() *ProjectState {
	this := ProjectState{}
	return &this
}

// NewProjectStateWithDefaults instantiates a new ProjectState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectStateWithDefaults() *ProjectState {
	this := ProjectState{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectState) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectState) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectState) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ProjectState) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *ProjectState) GetUptime() int32 {
	if o == nil || IsNil(o.Uptime) {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectState) GetUptimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *ProjectState) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *ProjectState) SetUptime(v int32) {
	o.Uptime = &v
}

func (o ProjectState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	return toSerialize, nil
}

type NullableProjectState struct {
	value *ProjectState
	isSet bool
}

func (v NullableProjectState) Get() *ProjectState {
	return v.value
}

func (v *NullableProjectState) Set(val *ProjectState) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectState) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectState(val *ProjectState) *NullableProjectState {
	return &NullableProjectState{value: val, isSet: true}
}

func (v NullableProjectState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}