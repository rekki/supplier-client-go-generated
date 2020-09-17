# V3OrderListInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **string** | Filters orders created at or after the given unix timestamp (in seconds) | [optional] 
**SkipIntegrated** | Pointer to **bool** | Only fetch orders that were not marked as integrated | [optional] [default to false]

## Methods

### NewV3OrderListInput

`func NewV3OrderListInput() *V3OrderListInput`

NewV3OrderListInput instantiates a new V3OrderListInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3OrderListInputWithDefaults

`func NewV3OrderListInputWithDefaults() *V3OrderListInput`

NewV3OrderListInputWithDefaults instantiates a new V3OrderListInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *V3OrderListInput) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *V3OrderListInput) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *V3OrderListInput) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *V3OrderListInput) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetSkipIntegrated

`func (o *V3OrderListInput) GetSkipIntegrated() bool`

GetSkipIntegrated returns the SkipIntegrated field if non-nil, zero value otherwise.

### GetSkipIntegratedOk

`func (o *V3OrderListInput) GetSkipIntegratedOk() (*bool, bool)`

GetSkipIntegratedOk returns a tuple with the SkipIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntegrated

`func (o *V3OrderListInput) SetSkipIntegrated(v bool)`

SetSkipIntegrated sets SkipIntegrated field to given value.

### HasSkipIntegrated

`func (o *V3OrderListInput) HasSkipIntegrated() bool`

HasSkipIntegrated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


