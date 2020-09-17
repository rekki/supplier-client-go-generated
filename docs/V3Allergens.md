# V3Allergens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symptoms** | Pointer to **[]string** | List of symptoms for the allergy. | [optional] 
**Type** | Pointer to **string** | Type of allergy. For example \&quot;contains peanuts\&quot; or \&quot;may contain peanuts\&quot;. | [optional] 

## Methods

### NewV3Allergens

`func NewV3Allergens() *V3Allergens`

NewV3Allergens instantiates a new V3Allergens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AllergensWithDefaults

`func NewV3AllergensWithDefaults() *V3Allergens`

NewV3AllergensWithDefaults instantiates a new V3Allergens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymptoms

`func (o *V3Allergens) GetSymptoms() []string`

GetSymptoms returns the Symptoms field if non-nil, zero value otherwise.

### GetSymptomsOk

`func (o *V3Allergens) GetSymptomsOk() (*[]string, bool)`

GetSymptomsOk returns a tuple with the Symptoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymptoms

`func (o *V3Allergens) SetSymptoms(v []string)`

SetSymptoms sets Symptoms field to given value.

### HasSymptoms

`func (o *V3Allergens) HasSymptoms() bool`

HasSymptoms returns a boolean if a field has been set.

### GetType

`func (o *V3Allergens) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V3Allergens) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V3Allergens) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V3Allergens) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


