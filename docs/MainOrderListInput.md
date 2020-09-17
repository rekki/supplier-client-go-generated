# MainOrderListInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **int32** | Filters orders created at or after the given unix timestamp (in seconds) | [optional] 

## Methods

### NewMainOrderListInput

`func NewMainOrderListInput() *MainOrderListInput`

NewMainOrderListInput instantiates a new MainOrderListInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainOrderListInputWithDefaults

`func NewMainOrderListInputWithDefaults() *MainOrderListInput`

NewMainOrderListInputWithDefaults instantiates a new MainOrderListInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *MainOrderListInput) GetSince() int32`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *MainOrderListInput) GetSinceOk() (*int32, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *MainOrderListInput) SetSince(v int32)`

SetSince sets Since field to given value.

### HasSince

`func (o *MainOrderListInput) HasSince() bool`

HasSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


