# CircuitGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**Group** | [**BriefCircuitGroup**](BriefCircuitGroup.md) |  | 
**MemberType** | **string** |  | 
**MemberId** | **int64** |  | 
**Member** | Pointer to **interface{}** |  | [optional] [readonly] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriority**](BriefCircuitGroupAssignmentSerializerPriority.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewCircuitGroupAssignment

`func NewCircuitGroupAssignment(id int32, url string, display string, group BriefCircuitGroup, memberType string, memberId int64, ) *CircuitGroupAssignment`

NewCircuitGroupAssignment instantiates a new CircuitGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitGroupAssignmentWithDefaults

`func NewCircuitGroupAssignmentWithDefaults() *CircuitGroupAssignment`

NewCircuitGroupAssignmentWithDefaults instantiates a new CircuitGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CircuitGroupAssignment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CircuitGroupAssignment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CircuitGroupAssignment) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CircuitGroupAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CircuitGroupAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CircuitGroupAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CircuitGroupAssignment) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CircuitGroupAssignment) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CircuitGroupAssignment) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *CircuitGroupAssignment) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CircuitGroupAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CircuitGroupAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CircuitGroupAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetGroup

`func (o *CircuitGroupAssignment) GetGroup() BriefCircuitGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CircuitGroupAssignment) GetGroupOk() (*BriefCircuitGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CircuitGroupAssignment) SetGroup(v BriefCircuitGroup)`

SetGroup sets Group field to given value.


### GetMemberType

`func (o *CircuitGroupAssignment) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *CircuitGroupAssignment) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *CircuitGroupAssignment) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.


### GetMemberId

`func (o *CircuitGroupAssignment) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *CircuitGroupAssignment) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *CircuitGroupAssignment) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.


### GetMember

`func (o *CircuitGroupAssignment) GetMember() interface{}`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *CircuitGroupAssignment) GetMemberOk() (*interface{}, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *CircuitGroupAssignment) SetMember(v interface{})`

SetMember sets Member field to given value.

### HasMember

`func (o *CircuitGroupAssignment) HasMember() bool`

HasMember returns a boolean if a field has been set.

### SetMemberNil

`func (o *CircuitGroupAssignment) SetMemberNil(b bool)`

 SetMemberNil sets the value for Member to be an explicit nil

### UnsetMember
`func (o *CircuitGroupAssignment) UnsetMember()`

UnsetMember ensures that no value is present for Member, not even an explicit nil
### GetPriority

`func (o *CircuitGroupAssignment) GetPriority() BriefCircuitGroupAssignmentSerializerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CircuitGroupAssignment) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CircuitGroupAssignment) SetPriority(v BriefCircuitGroupAssignmentSerializerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CircuitGroupAssignment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *CircuitGroupAssignment) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitGroupAssignment) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitGroupAssignment) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitGroupAssignment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreated

`func (o *CircuitGroupAssignment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CircuitGroupAssignment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CircuitGroupAssignment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CircuitGroupAssignment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *CircuitGroupAssignment) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CircuitGroupAssignment) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CircuitGroupAssignment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CircuitGroupAssignment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CircuitGroupAssignment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CircuitGroupAssignment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *CircuitGroupAssignment) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CircuitGroupAssignment) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


