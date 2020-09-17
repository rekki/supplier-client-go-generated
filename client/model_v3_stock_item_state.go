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

// V3StockItemState struct for V3StockItemState
type V3StockItemState struct {
	// only possible values are `in_stock` or `out_of_stock`
	Availability string `json:"availability"`
	ProductCode string `json:"product_code"`
}

// NewV3StockItemState instantiates a new V3StockItemState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3StockItemState(availability string, productCode string, ) *V3StockItemState {
	this := V3StockItemState{}
	this.Availability = availability
	this.ProductCode = productCode
	return &this
}

// NewV3StockItemStateWithDefaults instantiates a new V3StockItemState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3StockItemStateWithDefaults() *V3StockItemState {
	this := V3StockItemState{}
	return &this
}

// GetAvailability returns the Availability field value
func (o *V3StockItemState) GetAvailability() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *V3StockItemState) GetAvailabilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *V3StockItemState) SetAvailability(v string) {
	o.Availability = v
}

// GetProductCode returns the ProductCode field value
func (o *V3StockItemState) GetProductCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *V3StockItemState) GetProductCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *V3StockItemState) SetProductCode(v string) {
	o.ProductCode = v
}

func (o V3StockItemState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["availability"] = o.Availability
	}
	if true {
		toSerialize["product_code"] = o.ProductCode
	}
	return json.Marshal(toSerialize)
}

type NullableV3StockItemState struct {
	value *V3StockItemState
	isSet bool
}

func (v NullableV3StockItemState) Get() *V3StockItemState {
	return v.value
}

func (v *NullableV3StockItemState) Set(val *V3StockItemState) {
	v.value = val
	v.isSet = true
}

func (v NullableV3StockItemState) IsSet() bool {
	return v.isSet
}

func (v *NullableV3StockItemState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3StockItemState(val *V3StockItemState) *NullableV3StockItemState {
	return &NullableV3StockItemState{value: val, isSet: true}
}

func (v NullableV3StockItemState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3StockItemState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


