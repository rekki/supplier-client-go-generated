# CatalogAllergen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symptoms** | Pointer to **[]string** | List of symptoms for the allergy. | [optional] 
**Type** | Pointer to **string** | Type of allergy. For example \&quot;contains peanuts\&quot; or \&quot;may contain peanuts\&quot;. | [optional] 

## Methods

### NewCatalogAllergen

`func NewCatalogAllergen() *CatalogAllergen`

NewCatalogAllergen instantiates a new CatalogAllergen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAllergenWithDefaults

`func NewCatalogAllergenWithDefaults() *CatalogAllergen`

NewCatalogAllergenWithDefaults instantiates a new CatalogAllergen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymptoms

`func (o *CatalogAllergen) GetSymptoms() []string`

GetSymptoms returns the Symptoms field if non-nil, zero value otherwise.

### GetSymptomsOk

`func (o *CatalogAllergen) GetSymptomsOk() (*[]string, bool)`

GetSymptomsOk returns a tuple with the Symptoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymptoms

`func (o *CatalogAllergen) SetSymptoms(v []string)`

SetSymptoms sets Symptoms field to given value.

### HasSymptoms

`func (o *CatalogAllergen) HasSymptoms() bool`

HasSymptoms returns a boolean if a field has been set.

### GetType

`func (o *CatalogAllergen) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogAllergen) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogAllergen) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogAllergen) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


