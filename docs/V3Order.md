# V3Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmedAt** | Pointer to **string** | the time at which the supplier confirmed the order (via email, or via the REKKI supplier app, or via the REKKI API) | [optional] 
**ContactInfo** | Pointer to **string** | the phone number or email address for the person who placed the order | [optional] 
**ContactName** | Pointer to **string** | the full name of the person who placed the order | [optional] 
**CustomerAccountNo** | Pointer to **string** | the account number for customer within the supplier system, this can be setup in REKKI supplier app ( https://supplier.rekki.com ). | [optional] 
**DeliveryAddress** | Pointer to [**V3DeliveryAddress**](v3.DeliveryAddress.md) |  | [optional] 
**DeliveryOn** | Pointer to **string** | expected delivery date (when users place orders they specify for which day it is supposed to be delivered) | [optional] 
**InsertedAt** | Pointer to **string** | when was the order created by the customer (ISO date time) | [optional] 
**InsertedAtTs** | Pointer to **int32** | when was the order created by the customer (unix timestamp in seconds) | [optional] 
**Items** | Pointer to [**[]IntegrationOrderItem**](integration.OrderItem.md) | items requests in this order | [optional] 
**LocationName** | Pointer to **string** | the name of the location that placed the order, can be NULL | [optional] 
**Notes** | Pointer to **string** | defined by the user at the moment of making an order and usually refer to that specific order (e.g. \&quot;please send fresher tomatoes\&quot;) | [optional] 
**Reference** | Pointer to **string** | REKKI&#39;s order reference | [optional] 
**SupplierNotes** | Pointer to **string** | notes define by the user for the supplier, usually being common across orders (e.g.: \&quot;please use the side entrance for delivery\&quot;) | [optional] 

## Methods

### NewV3Order

`func NewV3Order() *V3Order`

NewV3Order instantiates a new V3Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3OrderWithDefaults

`func NewV3OrderWithDefaults() *V3Order`

NewV3OrderWithDefaults instantiates a new V3Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmedAt

`func (o *V3Order) GetConfirmedAt() string`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *V3Order) GetConfirmedAtOk() (*string, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *V3Order) SetConfirmedAt(v string)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *V3Order) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetContactInfo

`func (o *V3Order) GetContactInfo() string`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *V3Order) GetContactInfoOk() (*string, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *V3Order) SetContactInfo(v string)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *V3Order) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetContactName

`func (o *V3Order) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *V3Order) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *V3Order) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *V3Order) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetCustomerAccountNo

`func (o *V3Order) GetCustomerAccountNo() string`

GetCustomerAccountNo returns the CustomerAccountNo field if non-nil, zero value otherwise.

### GetCustomerAccountNoOk

`func (o *V3Order) GetCustomerAccountNoOk() (*string, bool)`

GetCustomerAccountNoOk returns a tuple with the CustomerAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountNo

`func (o *V3Order) SetCustomerAccountNo(v string)`

SetCustomerAccountNo sets CustomerAccountNo field to given value.

### HasCustomerAccountNo

`func (o *V3Order) HasCustomerAccountNo() bool`

HasCustomerAccountNo returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *V3Order) GetDeliveryAddress() V3DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *V3Order) GetDeliveryAddressOk() (*V3DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *V3Order) SetDeliveryAddress(v V3DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *V3Order) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryOn

`func (o *V3Order) GetDeliveryOn() string`

GetDeliveryOn returns the DeliveryOn field if non-nil, zero value otherwise.

### GetDeliveryOnOk

`func (o *V3Order) GetDeliveryOnOk() (*string, bool)`

GetDeliveryOnOk returns a tuple with the DeliveryOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOn

`func (o *V3Order) SetDeliveryOn(v string)`

SetDeliveryOn sets DeliveryOn field to given value.

### HasDeliveryOn

`func (o *V3Order) HasDeliveryOn() bool`

HasDeliveryOn returns a boolean if a field has been set.

### GetInsertedAt

`func (o *V3Order) GetInsertedAt() string`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *V3Order) GetInsertedAtOk() (*string, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *V3Order) SetInsertedAt(v string)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *V3Order) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.

### GetInsertedAtTs

`func (o *V3Order) GetInsertedAtTs() int32`

GetInsertedAtTs returns the InsertedAtTs field if non-nil, zero value otherwise.

### GetInsertedAtTsOk

`func (o *V3Order) GetInsertedAtTsOk() (*int32, bool)`

GetInsertedAtTsOk returns a tuple with the InsertedAtTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAtTs

`func (o *V3Order) SetInsertedAtTs(v int32)`

SetInsertedAtTs sets InsertedAtTs field to given value.

### HasInsertedAtTs

`func (o *V3Order) HasInsertedAtTs() bool`

HasInsertedAtTs returns a boolean if a field has been set.

### GetItems

`func (o *V3Order) GetItems() []IntegrationOrderItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V3Order) GetItemsOk() (*[]IntegrationOrderItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V3Order) SetItems(v []IntegrationOrderItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *V3Order) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLocationName

`func (o *V3Order) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *V3Order) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *V3Order) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *V3Order) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetNotes

`func (o *V3Order) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V3Order) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V3Order) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V3Order) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReference

`func (o *V3Order) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *V3Order) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *V3Order) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *V3Order) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSupplierNotes

`func (o *V3Order) GetSupplierNotes() string`

GetSupplierNotes returns the SupplierNotes field if non-nil, zero value otherwise.

### GetSupplierNotesOk

`func (o *V3Order) GetSupplierNotesOk() (*string, bool)`

GetSupplierNotesOk returns a tuple with the SupplierNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierNotes

`func (o *V3Order) SetSupplierNotes(v string)`

SetSupplierNotes sets SupplierNotes field to given value.

### HasSupplierNotes

`func (o *V3Order) HasSupplierNotes() bool`

HasSupplierNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


