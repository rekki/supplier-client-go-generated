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

// IntegrationOrderItem struct for IntegrationOrderItem
type IntegrationOrderItem struct {
	// REKKI's item id, for REKKI internal reference
	Id *string `json:"id,omitempty"`
	// item name as defined on the customer product list
	Name *string `json:"name,omitempty"`
	// the item price as set in REKKI
	Price *string `json:"price,omitempty"`
	// the item price as set in REKKI, in cents
	PriceCents *int32 `json:"price_cents,omitempty"`
	// product code that maps to the supplier catalog, suppliers can modify the product code for the future orders for this customer at https://supplier.rekki.com
	ProductCode *string `json:"product_code,omitempty"`
	// quantity
	Quantity *float32 `json:"quantity,omitempty"`
	// details/notes provided by the supplier for the item
	Spec *string `json:"spec,omitempty"`
	// item unit as defined on the customer product list
	Units *string `json:"units,omitempty"`
}

// NewIntegrationOrderItem instantiates a new IntegrationOrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOrderItem() *IntegrationOrderItem {
	this := IntegrationOrderItem{}
	return &this
}

// NewIntegrationOrderItemWithDefaults instantiates a new IntegrationOrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOrderItemWithDefaults() *IntegrationOrderItem {
	this := IntegrationOrderItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationOrderItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationOrderItem) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetPriceOk() (*string, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *IntegrationOrderItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceCents returns the PriceCents field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetPriceCents() int32 {
	if o == nil || o.PriceCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceCents
}

// GetPriceCentsOk returns a tuple with the PriceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetPriceCentsOk() (*int32, bool) {
	if o == nil || o.PriceCents == nil {
		return nil, false
	}
	return o.PriceCents, true
}

// HasPriceCents returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasPriceCents() bool {
	if o != nil && o.PriceCents != nil {
		return true
	}

	return false
}

// SetPriceCents gets a reference to the given int32 and assigns it to the PriceCents field.
func (o *IntegrationOrderItem) SetPriceCents(v int32) {
	o.PriceCents = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetProductCode() string {
	if o == nil || o.ProductCode == nil {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetProductCodeOk() (*string, bool) {
	if o == nil || o.ProductCode == nil {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasProductCode() bool {
	if o != nil && o.ProductCode != nil {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *IntegrationOrderItem) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *IntegrationOrderItem) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetSpec() string {
	if o == nil || o.Spec == nil {
		var ret string
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetSpecOk() (*string, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given string and assigns it to the Spec field.
func (o *IntegrationOrderItem) SetSpec(v string) {
	o.Spec = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *IntegrationOrderItem) GetUnits() string {
	if o == nil || o.Units == nil {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOrderItem) GetUnitsOk() (*string, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *IntegrationOrderItem) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *IntegrationOrderItem) SetUnits(v string) {
	o.Units = &v
}

func (o IntegrationOrderItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.PriceCents != nil {
		toSerialize["price_cents"] = o.PriceCents
	}
	if o.ProductCode != nil {
		toSerialize["product_code"] = o.ProductCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationOrderItem struct {
	value *IntegrationOrderItem
	isSet bool
}

func (v NullableIntegrationOrderItem) Get() *IntegrationOrderItem {
	return v.value
}

func (v *NullableIntegrationOrderItem) Set(val *IntegrationOrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOrderItem(val *IntegrationOrderItem) *NullableIntegrationOrderItem {
	return &NullableIntegrationOrderItem{value: val, isSet: true}
}

func (v NullableIntegrationOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


