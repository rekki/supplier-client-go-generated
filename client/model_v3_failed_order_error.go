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

// V3FailedOrderError struct for V3FailedOrderError
type V3FailedOrderError struct {
	Error *string `json:"error,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
}

// NewV3FailedOrderError instantiates a new V3FailedOrderError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3FailedOrderError() *V3FailedOrderError {
	this := V3FailedOrderError{}
	return &this
}

// NewV3FailedOrderErrorWithDefaults instantiates a new V3FailedOrderError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3FailedOrderErrorWithDefaults() *V3FailedOrderError {
	this := V3FailedOrderError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V3FailedOrderError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3FailedOrderError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V3FailedOrderError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V3FailedOrderError) SetError(v string) {
	o.Error = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *V3FailedOrderError) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3FailedOrderError) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *V3FailedOrderError) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *V3FailedOrderError) SetOrderId(v string) {
	o.OrderId = &v
}

func (o V3FailedOrderError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}
	return json.Marshal(toSerialize)
}

type NullableV3FailedOrderError struct {
	value *V3FailedOrderError
	isSet bool
}

func (v NullableV3FailedOrderError) Get() *V3FailedOrderError {
	return v.value
}

func (v *NullableV3FailedOrderError) Set(val *V3FailedOrderError) {
	v.value = val
	v.isSet = true
}

func (v NullableV3FailedOrderError) IsSet() bool {
	return v.isSet
}

func (v *NullableV3FailedOrderError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3FailedOrderError(val *V3FailedOrderError) *NullableV3FailedOrderError {
	return &NullableV3FailedOrderError{value: val, isSet: true}
}

func (v NullableV3FailedOrderError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3FailedOrderError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


