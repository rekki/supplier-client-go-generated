# V3OrderIntegrationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | Number of attempts made to push the integration | [optional] 
**Error** | Pointer to **string** | error message | [optional] 
**Reference** | Pointer to **string** | Order reference | [optional] 

## Methods

### NewV3OrderIntegrationError

`func NewV3OrderIntegrationError() *V3OrderIntegrationError`

NewV3OrderIntegrationError instantiates a new V3OrderIntegrationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3OrderIntegrationErrorWithDefaults

`func NewV3OrderIntegrationErrorWithDefaults() *V3OrderIntegrationError`

NewV3OrderIntegrationErrorWithDefaults instantiates a new V3OrderIntegrationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *V3OrderIntegrationError) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *V3OrderIntegrationError) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *V3OrderIntegrationError) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *V3OrderIntegrationError) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetError

`func (o *V3OrderIntegrationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V3OrderIntegrationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V3OrderIntegrationError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V3OrderIntegrationError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReference

`func (o *V3OrderIntegrationError) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *V3OrderIntegrationError) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *V3OrderIntegrationError) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *V3OrderIntegrationError) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


