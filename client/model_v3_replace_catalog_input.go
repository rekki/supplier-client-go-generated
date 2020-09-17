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

// V3ReplaceCatalogInput struct for V3ReplaceCatalogInput
type V3ReplaceCatalogInput struct {
	// items to insert
	Data *[]V3APISupplierCatalogItem `json:"data,omitempty"`
}

// NewV3ReplaceCatalogInput instantiates a new V3ReplaceCatalogInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3ReplaceCatalogInput() *V3ReplaceCatalogInput {
	this := V3ReplaceCatalogInput{}
	return &this
}

// NewV3ReplaceCatalogInputWithDefaults instantiates a new V3ReplaceCatalogInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3ReplaceCatalogInputWithDefaults() *V3ReplaceCatalogInput {
	this := V3ReplaceCatalogInput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *V3ReplaceCatalogInput) GetData() []V3APISupplierCatalogItem {
	if o == nil || o.Data == nil {
		var ret []V3APISupplierCatalogItem
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3ReplaceCatalogInput) GetDataOk() (*[]V3APISupplierCatalogItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *V3ReplaceCatalogInput) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []V3APISupplierCatalogItem and assigns it to the Data field.
func (o *V3ReplaceCatalogInput) SetData(v []V3APISupplierCatalogItem) {
	o.Data = &v
}

func (o V3ReplaceCatalogInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableV3ReplaceCatalogInput struct {
	value *V3ReplaceCatalogInput
	isSet bool
}

func (v NullableV3ReplaceCatalogInput) Get() *V3ReplaceCatalogInput {
	return v.value
}

func (v *NullableV3ReplaceCatalogInput) Set(val *V3ReplaceCatalogInput) {
	v.value = val
	v.isSet = true
}

func (v NullableV3ReplaceCatalogInput) IsSet() bool {
	return v.isSet
}

func (v *NullableV3ReplaceCatalogInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3ReplaceCatalogInput(val *V3ReplaceCatalogInput) *NullableV3ReplaceCatalogInput {
	return &NullableV3ReplaceCatalogInput{value: val, isSet: true}
}

func (v NullableV3ReplaceCatalogInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3ReplaceCatalogInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


