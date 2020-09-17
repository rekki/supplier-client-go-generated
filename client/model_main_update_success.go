/*
 * Rekki.com Supplier API
 *
 * The base URL for all API endpoints is https://api.rekki.com  Api key value contains of word Bearer together with api key that you can get from integrations@rekki.com 
 *
 * API version: 
 * Contact: integrations@rekki.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MainUpdateSuccess struct for MainUpdateSuccess
type MainUpdateSuccess struct {
	// number of items actually updated
	Affected *int32 `json:"affected,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewMainUpdateSuccess instantiates a new MainUpdateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainUpdateSuccess() *MainUpdateSuccess {
	this := MainUpdateSuccess{}
	return &this
}

// NewMainUpdateSuccessWithDefaults instantiates a new MainUpdateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainUpdateSuccessWithDefaults() *MainUpdateSuccess {
	this := MainUpdateSuccess{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *MainUpdateSuccess) GetAffected() int32 {
	if o == nil || o.Affected == nil {
		var ret int32
		return ret
	}
	return *o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainUpdateSuccess) GetAffectedOk() (*int32, bool) {
	if o == nil || o.Affected == nil {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *MainUpdateSuccess) HasAffected() bool {
	if o != nil && o.Affected != nil {
		return true
	}

	return false
}

// SetAffected gets a reference to the given int32 and assigns it to the Affected field.
func (o *MainUpdateSuccess) SetAffected(v int32) {
	o.Affected = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MainUpdateSuccess) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainUpdateSuccess) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MainUpdateSuccess) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *MainUpdateSuccess) SetSuccess(v bool) {
	o.Success = &v
}

func (o MainUpdateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Affected != nil {
		toSerialize["affected"] = o.Affected
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableMainUpdateSuccess struct {
	value *MainUpdateSuccess
	isSet bool
}

func (v NullableMainUpdateSuccess) Get() *MainUpdateSuccess {
	return v.value
}

func (v *NullableMainUpdateSuccess) Set(val *MainUpdateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableMainUpdateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableMainUpdateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainUpdateSuccess(val *MainUpdateSuccess) *NullableMainUpdateSuccess {
	return &NullableMainUpdateSuccess{value: val, isSet: true}
}

func (v NullableMainUpdateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainUpdateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


