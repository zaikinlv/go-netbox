# WritableVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to [**NullablePatchedWritableVLANRequestSite**](PatchedWritableVLANRequestSite.md) |  | [optional] 
**Group** | Pointer to [**NullablePatchedWritableVLANRequestGroup**](PatchedWritableVLANRequestGroup.md) |  | [optional] 
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritableVLANRequestStatus**](PatchedWritableVLANRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableIPRangeRequestRole**](IPRangeRequestRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**QinqRole** | Pointer to [**NullableQInQRole**](QInQRole.md) |  | [optional] 
**QinqSvlan** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableVLANRequest

`func NewWritableVLANRequest(vid int32, name string, ) *WritableVLANRequest`

NewWritableVLANRequest instantiates a new WritableVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVLANRequestWithDefaults

`func NewWritableVLANRequestWithDefaults() *WritableVLANRequest`

NewWritableVLANRequestWithDefaults instantiates a new WritableVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *WritableVLANRequest) GetSite() PatchedWritableVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableVLANRequest) GetSiteOk() (*PatchedWritableVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableVLANRequest) SetSite(v PatchedWritableVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *WritableVLANRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *WritableVLANRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *WritableVLANRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *WritableVLANRequest) GetGroup() PatchedWritableVLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableVLANRequest) GetGroupOk() (*PatchedWritableVLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableVLANRequest) SetGroup(v PatchedWritableVLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableVLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableVLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableVLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *WritableVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *WritableVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *WritableVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *WritableVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVLANRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *WritableVLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableVLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableVLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableVLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableVLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableVLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *WritableVLANRequest) GetStatus() PatchedWritableVLANRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableVLANRequest) GetStatusOk() (*PatchedWritableVLANRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableVLANRequest) SetStatus(v PatchedWritableVLANRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableVLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *WritableVLANRequest) GetRole() IPRangeRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableVLANRequest) GetRoleOk() (*IPRangeRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableVLANRequest) SetRole(v IPRangeRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableVLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritableVLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritableVLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *WritableVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQinqRole

`func (o *WritableVLANRequest) GetQinqRole() QInQRole`

GetQinqRole returns the QinqRole field if non-nil, zero value otherwise.

### GetQinqRoleOk

`func (o *WritableVLANRequest) GetQinqRoleOk() (*QInQRole, bool)`

GetQinqRoleOk returns a tuple with the QinqRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqRole

`func (o *WritableVLANRequest) SetQinqRole(v QInQRole)`

SetQinqRole sets QinqRole field to given value.

### HasQinqRole

`func (o *WritableVLANRequest) HasQinqRole() bool`

HasQinqRole returns a boolean if a field has been set.

### SetQinqRoleNil

`func (o *WritableVLANRequest) SetQinqRoleNil(b bool)`

 SetQinqRoleNil sets the value for QinqRole to be an explicit nil

### UnsetQinqRole
`func (o *WritableVLANRequest) UnsetQinqRole()`

UnsetQinqRole ensures that no value is present for QinqRole, not even an explicit nil
### GetQinqSvlan

`func (o *WritableVLANRequest) GetQinqSvlan() int32`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *WritableVLANRequest) GetQinqSvlanOk() (*int32, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *WritableVLANRequest) SetQinqSvlan(v int32)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *WritableVLANRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *WritableVLANRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *WritableVLANRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetComments

`func (o *WritableVLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableVLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableVLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableVLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableVLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


