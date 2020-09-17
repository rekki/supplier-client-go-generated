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

// V3Allergens struct for V3Allergens
type V3Allergens struct {
	// List of symptoms for the allergy.
	Symptoms *[]string `json:"symptoms,omitempty"`
	// Type of allergy. For example \"contains peanuts\" or \"may contain peanuts\".
	Type *string `json:"type,omitempty"`
}

// NewV3Allergens instantiates a new V3Allergens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3Allergens() *V3Allergens {
	this := V3Allergens{}
	return &this
}

// NewV3AllergensWithDefaults instantiates a new V3Allergens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3AllergensWithDefaults() *V3Allergens {
	this := V3Allergens{}
	return &this
}

// GetSymptoms returns the Symptoms field value if set, zero value otherwise.
func (o *V3Allergens) GetSymptoms() []string {
	if o == nil || o.Symptoms == nil {
		var ret []string
		return ret
	}
	return *o.Symptoms
}

// GetSymptomsOk returns a tuple with the Symptoms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Allergens) GetSymptomsOk() (*[]string, bool) {
	if o == nil || o.Symptoms == nil {
		return nil, false
	}
	return o.Symptoms, true
}

// HasSymptoms returns a boolean if a field has been set.
func (o *V3Allergens) HasSymptoms() bool {
	if o != nil && o.Symptoms != nil {
		return true
	}

	return false
}

// SetSymptoms gets a reference to the given []string and assigns it to the Symptoms field.
func (o *V3Allergens) SetSymptoms(v []string) {
	o.Symptoms = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V3Allergens) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Allergens) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V3Allergens) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V3Allergens) SetType(v string) {
	o.Type = &v
}

func (o V3Allergens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symptoms != nil {
		toSerialize["symptoms"] = o.Symptoms
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableV3Allergens struct {
	value *V3Allergens
	isSet bool
}

func (v NullableV3Allergens) Get() *V3Allergens {
	return v.value
}

func (v *NullableV3Allergens) Set(val *V3Allergens) {
	v.value = val
	v.isSet = true
}

func (v NullableV3Allergens) IsSet() bool {
	return v.isSet
}

func (v *NullableV3Allergens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3Allergens(val *V3Allergens) *NullableV3Allergens {
	return &NullableV3Allergens{value: val, isSet: true}
}

func (v NullableV3Allergens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3Allergens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


