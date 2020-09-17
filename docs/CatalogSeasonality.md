# CatalogSeasonality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **string** | The end date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD. | [optional] 
**StartDate** | Pointer to **string** | The start date when the item is in season. In ISO 8601 calendar date format YYYY-MM-DD. | [optional] 

## Methods

### NewCatalogSeasonality

`func NewCatalogSeasonality() *CatalogSeasonality`

NewCatalogSeasonality instantiates a new CatalogSeasonality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogSeasonalityWithDefaults

`func NewCatalogSeasonalityWithDefaults() *CatalogSeasonality`

NewCatalogSeasonalityWithDefaults instantiates a new CatalogSeasonality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *CatalogSeasonality) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CatalogSeasonality) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CatalogSeasonality) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CatalogSeasonality) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CatalogSeasonality) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CatalogSeasonality) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CatalogSeasonality) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CatalogSeasonality) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


