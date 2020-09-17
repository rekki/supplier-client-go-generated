# MainLogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | possible values are: \&quot;panic\&quot;, \&quot;error\&quot;, \&quot;fatal\&quot;, \&quot;warn\&quot;, \&quot;warning\&quot;, \&quot;info\&quot;, \&quot;debug\&quot;, \&quot;trace\&quot; | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewMainLogMessage

`func NewMainLogMessage() *MainLogMessage`

NewMainLogMessage instantiates a new MainLogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainLogMessageWithDefaults

`func NewMainLogMessageWithDefaults() *MainLogMessage`

NewMainLogMessageWithDefaults instantiates a new MainLogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *MainLogMessage) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MainLogMessage) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MainLogMessage) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MainLogMessage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *MainLogMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MainLogMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MainLogMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MainLogMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


