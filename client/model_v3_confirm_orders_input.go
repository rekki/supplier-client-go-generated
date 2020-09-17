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

// V3ConfirmOrdersInput struct for V3ConfirmOrdersInput
type V3ConfirmOrdersInput struct {
	// Array of References of the orders to confirm, required. Order refs are discoverable when listing orders.
	Orders *[]string `json:"orders,omitempty"`
}

// NewV3ConfirmOrdersInput instantiates a new V3ConfirmOrdersInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3ConfirmOrdersInput() *V3ConfirmOrdersInput {
	this := V3ConfirmOrdersInput{}
	return &this
}

// NewV3ConfirmOrdersInputWithDefaults instantiates a new V3ConfirmOrdersInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ConfirmOrdersInputWithDefaults() *V3ConfirmOrdersInput {
	this := V3ConfirmOrdersInput{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *V3ConfirmOrdersInput) GetOrders() []string {
	if o == nil || o.Orders == nil {
		var ret []string
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3ConfirmOrdersInput) GetOrdersOk() (*[]string, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *V3ConfirmOrdersInput) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *V3ConfirmOrdersInput) SetOrders(v []string) {
	o.Orders = &v
}

func (o V3ConfirmOrdersInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableV3ConfirmOrdersInput struct {
	value *V3ConfirmOrdersInput
	isSet bool
}

func (v NullableV3ConfirmOrdersInput) Get() *V3ConfirmOrdersInput {
	return v.value
}

func (v *NullableV3ConfirmOrdersInput) Set(val *V3ConfirmOrdersInput) {
	v.value = val
	v.isSet = true
}

func (v NullableV3ConfirmOrdersInput) IsSet() bool {
	return v.isSet
}

func (v *NullableV3ConfirmOrdersInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3ConfirmOrdersInput(val *V3ConfirmOrdersInput) *NullableV3ConfirmOrdersInput {
	return &NullableV3ConfirmOrdersInput{value: val, isSet: true}
}

func (v NullableV3ConfirmOrdersInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3ConfirmOrdersInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


