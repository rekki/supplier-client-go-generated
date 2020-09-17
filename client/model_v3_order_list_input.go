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

// V3OrderListInput struct for V3OrderListInput
type V3OrderListInput struct {
	// Filters orders created at or after the given unix timestamp (in seconds)
	Since *string `json:"since,omitempty"`
	// Only fetch orders that were not marked as integrated
	SkipIntegrated *bool `json:"skip_integrated,omitempty"`
}

// NewV3OrderListInput instantiates a new V3OrderListInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3OrderListInput() *V3OrderListInput {
	this := V3OrderListInput{}
	var skipIntegrated bool = false
	this.SkipIntegrated = &skipIntegrated
	return &this
}

// NewV3OrderListInputWithDefaults instantiates a new V3OrderListInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3OrderListInputWithDefaults() *V3OrderListInput {
	this := V3OrderListInput{}
	var skipIntegrated bool = false
	this.SkipIntegrated = &skipIntegrated
	return &this
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *V3OrderListInput) GetSince() string {
	if o == nil || o.Since == nil {
		var ret string
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3OrderListInput) GetSinceOk() (*string, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *V3OrderListInput) HasSince() bool {
	if o != nil && o.Since != nil {
		return true
	}

	return false
}

// SetSince gets a reference to the given string and assigns it to the Since field.
func (o *V3OrderListInput) SetSince(v string) {
	o.Since = &v
}

// GetSkipIntegrated returns the SkipIntegrated field value if set, zero value otherwise.
func (o *V3OrderListInput) GetSkipIntegrated() bool {
	if o == nil || o.SkipIntegrated == nil {
		var ret bool
		return ret
	}
	return *o.SkipIntegrated
}

// GetSkipIntegratedOk returns a tuple with the SkipIntegrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3OrderListInput) GetSkipIntegratedOk() (*bool, bool) {
	if o == nil || o.SkipIntegrated == nil {
		return nil, false
	}
	return o.SkipIntegrated, true
}

// HasSkipIntegrated returns a boolean if a field has been set.
func (o *V3OrderListInput) HasSkipIntegrated() bool {
	if o != nil && o.SkipIntegrated != nil {
		return true
	}

	return false
}

// SetSkipIntegrated gets a reference to the given bool and assigns it to the SkipIntegrated field.
func (o *V3OrderListInput) SetSkipIntegrated(v bool) {
	o.SkipIntegrated = &v
}

func (o V3OrderListInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.SkipIntegrated != nil {
		toSerialize["skip_integrated"] = o.SkipIntegrated
	}
	return json.Marshal(toSerialize)
}

type NullableV3OrderListInput struct {
	value *V3OrderListInput
	isSet bool
}

func (v NullableV3OrderListInput) Get() *V3OrderListInput {
	return v.value
}

func (v *NullableV3OrderListInput) Set(val *V3OrderListInput) {
	v.value = val
	v.isSet = true
}

func (v NullableV3OrderListInput) IsSet() bool {
	return v.isSet
}

func (v *NullableV3OrderListInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3OrderListInput(val *V3OrderListInput) *NullableV3OrderListInput {
	return &NullableV3OrderListInput{value: val, isSet: true}
}

func (v NullableV3OrderListInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3OrderListInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


