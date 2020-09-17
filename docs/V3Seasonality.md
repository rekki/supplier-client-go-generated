# V3Seasonality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **string** | The end date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD. | [optional] 
**StartDate** | Pointer to **string** | The start date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD. | [optional] 

## Methods

### NewV3Seasonality

`func NewV3Seasonality() *V3Seasonality`

NewV3Seasonality instantiates a new V3Seasonality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SeasonalityWithDefaults

`func NewV3SeasonalityWithDefaults() *V3Seasonality`

NewV3SeasonalityWithDefaults instantiates a new V3Seasonality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *V3Seasonality) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *V3Seasonality) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *V3Seasonality) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *V3Seasonality) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *V3Seasonality) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *V3Seasonality) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *V3Seasonality) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *V3Seasonality) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


