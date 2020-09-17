# IntegrationOrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | REKKI&#39;s item id, for REKKI internal reference | [optional] 
**Name** | Pointer to **string** | item name as defined on the customer product list | [optional] 
**Price** | Pointer to **string** | the item price as set in REKKI | [optional] 
**PriceCents** | Pointer to **int32** | the item price as set in REKKI, in cents | [optional] 
**ProductCode** | Pointer to **string** | product code that maps to the supplier catalog, suppliers can modify the product code for the future orders for this customer at https://supplier.rekki.com | [optional] 
**Quantity** | Pointer to **float32** | quantity | [optional] 
**Spec** | Pointer to **string** | details/notes provided by the supplier for the item | [optional] 
**Units** | Pointer to **string** | item unit as defined on the customer product list | [optional] 

## Methods

### NewIntegrationOrderItem

`func NewIntegrationOrderItem() *IntegrationOrderItem`

NewIntegrationOrderItem instantiates a new IntegrationOrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationOrderItemWithDefaults

`func NewIntegrationOrderItemWithDefaults() *IntegrationOrderItem`

NewIntegrationOrderItemWithDefaults instantiates a new IntegrationOrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationOrderItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationOrderItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationOrderItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationOrderItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationOrderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationOrderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationOrderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationOrderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *IntegrationOrderItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *IntegrationOrderItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *IntegrationOrderItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *IntegrationOrderItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceCents

`func (o *IntegrationOrderItem) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *IntegrationOrderItem) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *IntegrationOrderItem) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *IntegrationOrderItem) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetProductCode

`func (o *IntegrationOrderItem) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *IntegrationOrderItem) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *IntegrationOrderItem) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *IntegrationOrderItem) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetQuantity

`func (o *IntegrationOrderItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *IntegrationOrderItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *IntegrationOrderItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *IntegrationOrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSpec

`func (o *IntegrationOrderItem) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IntegrationOrderItem) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IntegrationOrderItem) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IntegrationOrderItem) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetUnits

`func (o *IntegrationOrderItem) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *IntegrationOrderItem) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *IntegrationOrderItem) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *IntegrationOrderItem) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


