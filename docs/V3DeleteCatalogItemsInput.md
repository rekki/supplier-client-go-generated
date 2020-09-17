# V3DeleteCatalogItemsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]int32** | Array of catalog item ids to delete | [optional] 

## Methods

### NewV3DeleteCatalogItemsInput

`func NewV3DeleteCatalogItemsInput() *V3DeleteCatalogItemsInput`

NewV3DeleteCatalogItemsInput instantiates a new V3DeleteCatalogItemsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3DeleteCatalogItemsInputWithDefaults

`func NewV3DeleteCatalogItemsInputWithDefaults() *V3DeleteCatalogItemsInput`

NewV3DeleteCatalogItemsInputWithDefaults instantiates a new V3DeleteCatalogItemsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V3DeleteCatalogItemsInput) GetItems() []int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V3DeleteCatalogItemsInput) GetItemsOk() (*[]int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V3DeleteCatalogItemsInput) SetItems(v []int32)`

SetItems sets Items field to given value.

### HasItems

`func (o *V3DeleteCatalogItemsInput) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


