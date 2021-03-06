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

// CatalogOrderCutoffTimes struct for CatalogOrderCutoffTimes
type CatalogOrderCutoffTimes struct {
	// Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day.
	Fri *int32 `json:"fri,omitempty"`
	// Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day.
	Mon *int32 `json:"mon,omitempty"`
	// Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day.
	Thu *int32 `json:"thu,omitempty"`
	// Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day.
	Tue *int32 `json:"tue,omitempty"`
	// Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day.
	Wed *int32 `json:"wed,omitempty"`
}

// NewCatalogOrderCutoffTimes instantiates a new CatalogOrderCutoffTimes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogOrderCutoffTimes() *CatalogOrderCutoffTimes {
	this := CatalogOrderCutoffTimes{}
	return &this
}

// NewCatalogOrderCutoffTimesWithDefaults instantiates a new CatalogOrderCutoffTimes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogOrderCutoffTimesWithDefaults() *CatalogOrderCutoffTimes {
	this := CatalogOrderCutoffTimes{}
	return &this
}

// GetFri returns the Fri field value if set, zero value otherwise.
func (o *CatalogOrderCutoffTimes) GetFri() int32 {
	if o == nil || o.Fri == nil {
		var ret int32
		return ret
	}
	return *o.Fri
}

// GetFriOk returns a tuple with the Fri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogOrderCutoffTimes) GetFriOk() (*int32, bool) {
	if o == nil || o.Fri == nil {
		return nil, false
	}
	return o.Fri, true
}

// HasFri returns a boolean if a field has been set.
func (o *CatalogOrderCutoffTimes) HasFri() bool {
	if o != nil && o.Fri != nil {
		return true
	}

	return false
}

// SetFri gets a reference to the given int32 and assigns it to the Fri field.
func (o *CatalogOrderCutoffTimes) SetFri(v int32) {
	o.Fri = &v
}

// GetMon returns the Mon field value if set, zero value otherwise.
func (o *CatalogOrderCutoffTimes) GetMon() int32 {
	if o == nil || o.Mon == nil {
		var ret int32
		return ret
	}
	return *o.Mon
}

// GetMonOk returns a tuple with the Mon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogOrderCutoffTimes) GetMonOk() (*int32, bool) {
	if o == nil || o.Mon == nil {
		return nil, false
	}
	return o.Mon, true
}

// HasMon returns a boolean if a field has been set.
func (o *CatalogOrderCutoffTimes) HasMon() bool {
	if o != nil && o.Mon != nil {
		return true
	}

	return false
}

// SetMon gets a reference to the given int32 and assigns it to the Mon field.
func (o *CatalogOrderCutoffTimes) SetMon(v int32) {
	o.Mon = &v
}

// GetThu returns the Thu field value if set, zero value otherwise.
func (o *CatalogOrderCutoffTimes) GetThu() int32 {
	if o == nil || o.Thu == nil {
		var ret int32
		return ret
	}
	return *o.Thu
}

// GetThuOk returns a tuple with the Thu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogOrderCutoffTimes) GetThuOk() (*int32, bool) {
	if o == nil || o.Thu == nil {
		return nil, false
	}
	return o.Thu, true
}

// HasThu returns a boolean if a field has been set.
func (o *CatalogOrderCutoffTimes) HasThu() bool {
	if o != nil && o.Thu != nil {
		return true
	}

	return false
}

// SetThu gets a reference to the given int32 and assigns it to the Thu field.
func (o *CatalogOrderCutoffTimes) SetThu(v int32) {
	o.Thu = &v
}

// GetTue returns the Tue field value if set, zero value otherwise.
func (o *CatalogOrderCutoffTimes) GetTue() int32 {
	if o == nil || o.Tue == nil {
		var ret int32
		return ret
	}
	return *o.Tue
}

// GetTueOk returns a tuple with the Tue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogOrderCutoffTimes) GetTueOk() (*int32, bool) {
	if o == nil || o.Tue == nil {
		return nil, false
	}
	return o.Tue, true
}

// HasTue returns a boolean if a field has been set.
func (o *CatalogOrderCutoffTimes) HasTue() bool {
	if o != nil && o.Tue != nil {
		return true
	}

	return false
}

// SetTue gets a reference to the given int32 and assigns it to the Tue field.
func (o *CatalogOrderCutoffTimes) SetTue(v int32) {
	o.Tue = &v
}

// GetWed returns the Wed field value if set, zero value otherwise.
func (o *CatalogOrderCutoffTimes) GetWed() int32 {
	if o == nil || o.Wed == nil {
		var ret int32
		return ret
	}
	return *o.Wed
}

// GetWedOk returns a tuple with the Wed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogOrderCutoffTimes) GetWedOk() (*int32, bool) {
	if o == nil || o.Wed == nil {
		return nil, false
	}
	return o.Wed, true
}

// HasWed returns a boolean if a field has been set.
func (o *CatalogOrderCutoffTimes) HasWed() bool {
	if o != nil && o.Wed != nil {
		return true
	}

	return false
}

// SetWed gets a reference to the given int32 and assigns it to the Wed field.
func (o *CatalogOrderCutoffTimes) SetWed(v int32) {
	o.Wed = &v
}

func (o CatalogOrderCutoffTimes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fri != nil {
		toSerialize["fri"] = o.Fri
	}
	if o.Mon != nil {
		toSerialize["mon"] = o.Mon
	}
	if o.Thu != nil {
		toSerialize["thu"] = o.Thu
	}
	if o.Tue != nil {
		toSerialize["tue"] = o.Tue
	}
	if o.Wed != nil {
		toSerialize["wed"] = o.Wed
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogOrderCutoffTimes struct {
	value *CatalogOrderCutoffTimes
	isSet bool
}

func (v NullableCatalogOrderCutoffTimes) Get() *CatalogOrderCutoffTimes {
	return v.value
}

func (v *NullableCatalogOrderCutoffTimes) Set(val *CatalogOrderCutoffTimes) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogOrderCutoffTimes) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogOrderCutoffTimes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogOrderCutoffTimes(val *CatalogOrderCutoffTimes) *NullableCatalogOrderCutoffTimes {
	return &NullableCatalogOrderCutoffTimes{value: val, isSet: true}
}

func (v NullableCatalogOrderCutoffTimes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogOrderCutoffTimes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


