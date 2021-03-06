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

// MainOrderListInput struct for MainOrderListInput
type MainOrderListInput struct {
	// Filters orders created at or after the given unix timestamp (in seconds)
	Since *int32 `json:"since,omitempty"`
}

// NewMainOrderListInput instantiates a new MainOrderListInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainOrderListInput() *MainOrderListInput {
	this := MainOrderListInput{}
	return &this
}

// NewMainOrderListInputWithDefaults instantiates a new MainOrderListInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainOrderListInputWithDefaults() *MainOrderListInput {
	this := MainOrderListInput{}
	return &this
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *MainOrderListInput) GetSince() int32 {
	if o == nil || o.Since == nil {
		var ret int32
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainOrderListInput) GetSinceOk() (*int32, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *MainOrderListInput) HasSince() bool {
	if o != nil && o.Since != nil {
		return true
	}

	return false
}

// SetSince gets a reference to the given int32 and assigns it to the Since field.
func (o *MainOrderListInput) SetSince(v int32) {
	o.Since = &v
}

func (o MainOrderListInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	return json.Marshal(toSerialize)
}

type NullableMainOrderListInput struct {
	value *MainOrderListInput
	isSet bool
}

func (v NullableMainOrderListInput) Get() *MainOrderListInput {
	return v.value
}

func (v *NullableMainOrderListInput) Set(val *MainOrderListInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMainOrderListInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMainOrderListInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainOrderListInput(val *MainOrderListInput) *NullableMainOrderListInput {
	return &NullableMainOrderListInput{value: val, isSet: true}
}

func (v NullableMainOrderListInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainOrderListInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


