# CircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefCircuitRequestProvider**](BriefCircuitRequestProvider.md) |  | 
**ProviderAccount** | Pointer to [**NullableCircuitRequestProviderAccount**](CircuitRequestProviderAccount.md) |  | [optional] 
**Type** | [**CircuitRequestType**](CircuitRequestType.md) |  | 
**Status** | Pointer to [**CircuitStatusValue**](CircuitStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**TerminationDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** | Committed rate | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **NullableFloat64** |  | [optional] 
**DistanceUnit** | Pointer to [**NullableCircuitRequestDistanceUnit**](CircuitRequestDistanceUnit.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Assignments** | Pointer to [**[]BriefCircuitGroupAssignmentSerializerRequest**](BriefCircuitGroupAssignmentSerializerRequest.md) |  | [optional] 

## Methods

### NewCircuitRequest

`func NewCircuitRequest(cid string, provider BriefCircuitRequestProvider, type_ CircuitRequestType, ) *CircuitRequest`

NewCircuitRequest instantiates a new CircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitRequestWithDefaults

`func NewCircuitRequestWithDefaults() *CircuitRequest`

NewCircuitRequestWithDefaults instantiates a new CircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *CircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *CircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *CircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *CircuitRequest) GetProvider() BriefCircuitRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CircuitRequest) GetProviderOk() (*BriefCircuitRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CircuitRequest) SetProvider(v BriefCircuitRequestProvider)`

SetProvider sets Provider field to given value.


### GetProviderAccount

`func (o *CircuitRequest) GetProviderAccount() CircuitRequestProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *CircuitRequest) GetProviderAccountOk() (*CircuitRequestProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *CircuitRequest) SetProviderAccount(v CircuitRequestProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *CircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *CircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *CircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *CircuitRequest) GetType() CircuitRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CircuitRequest) GetTypeOk() (*CircuitRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CircuitRequest) SetType(v CircuitRequestType)`

SetType sets Type field to given value.


### GetStatus

`func (o *CircuitRequest) GetStatus() CircuitStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CircuitRequest) GetStatusOk() (*CircuitStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CircuitRequest) SetStatus(v CircuitStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *CircuitRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CircuitRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CircuitRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *CircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *CircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *CircuitRequest) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *CircuitRequest) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *CircuitRequest) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *CircuitRequest) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *CircuitRequest) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *CircuitRequest) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetTerminationDate

`func (o *CircuitRequest) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *CircuitRequest) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *CircuitRequest) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *CircuitRequest) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *CircuitRequest) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *CircuitRequest) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetCommitRate

`func (o *CircuitRequest) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *CircuitRequest) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *CircuitRequest) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *CircuitRequest) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *CircuitRequest) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *CircuitRequest) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *CircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDistance

`func (o *CircuitRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CircuitRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CircuitRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *CircuitRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *CircuitRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *CircuitRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *CircuitRequest) GetDistanceUnit() CircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *CircuitRequest) GetDistanceUnitOk() (*CircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *CircuitRequest) SetDistanceUnit(v CircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *CircuitRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *CircuitRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *CircuitRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetComments

`func (o *CircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *CircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetAssignments

`func (o *CircuitRequest) GetAssignments() []BriefCircuitGroupAssignmentSerializerRequest`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *CircuitRequest) GetAssignmentsOk() (*[]BriefCircuitGroupAssignmentSerializerRequest, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *CircuitRequest) SetAssignments(v []BriefCircuitGroupAssignmentSerializerRequest)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *CircuitRequest) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


