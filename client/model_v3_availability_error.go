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

// V3AvailabilityError struct for V3AvailabilityError
type V3AvailabilityError struct {
	Error *string `json:"error,omitempty"`
	Item *V3StockItemState `json:"item,omitempty"`
}

// NewV3AvailabilityError instantiates a new V3AvailabilityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3AvailabilityError() *V3AvailabilityError {
	this := V3AvailabilityError{}
	return &this
}

// NewV3AvailabilityErrorWithDefaults instantiates a new V3AvailabilityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AvailabilityErrorWithDefaults() *V3AvailabilityError {
	this := V3AvailabilityError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V3AvailabilityError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AvailabilityError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V3AvailabilityError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V3AvailabilityError) SetError(v string) {
	o.Error = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *V3AvailabilityError) GetItem() V3StockItemState {
	if o == nil || o.Item == nil {
		var ret V3StockItemState
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3AvailabilityError) GetItemOk() (*V3StockItemState, bool) {
	if o == nil || o.Item == nil {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *V3AvailabilityError) HasItem() bool {
	if o != nil && o.Item != nil {
		return true
	}

	return false
}

// SetItem gets a reference to the given V3StockItemState and assigns it to the Item field.
func (o *V3AvailabilityError) SetItem(v V3StockItemState) {
	o.Item = &v
}

func (o V3AvailabilityError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Item != nil {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableV3AvailabilityError struct {
	value *V3AvailabilityError
	isSet bool
}

func (v NullableV3AvailabilityError) Get() *V3AvailabilityError {
	return v.value
}

func (v *NullableV3AvailabilityError) Set(val *V3AvailabilityError) {
	v.value = val
	v.isSet = true
}

func (v NullableV3AvailabilityError) IsSet() bool {
	return v.isSet
}

func (v *NullableV3AvailabilityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3AvailabilityError(val *V3AvailabilityError) *NullableV3AvailabilityError {
	return &NullableV3AvailabilityError{value: val, isSet: true}
}

func (v NullableV3AvailabilityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3AvailabilityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


