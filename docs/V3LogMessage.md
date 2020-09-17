# V3LogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | possible values are: \&quot;panic\&quot;, \&quot;error\&quot;, \&quot;fatal\&quot;, \&quot;warn\&quot;, \&quot;warning\&quot;, \&quot;info\&quot;, \&quot;debug\&quot;, \&quot;trace\&quot; | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewV3LogMessage

`func NewV3LogMessage() *V3LogMessage`

NewV3LogMessage instantiates a new V3LogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3LogMessageWithDefaults

`func NewV3LogMessageWithDefaults() *V3LogMessage`

NewV3LogMessageWithDefaults instantiates a new V3LogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *V3LogMessage) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *V3LogMessage) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *V3LogMessage) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *V3LogMessage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *V3LogMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V3LogMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V3LogMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V3LogMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


