# PatchedWritableDeviceWithConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to [**DeviceBayTemplateRequestDeviceType**](DeviceBayTemplateRequestDeviceType.md) |  | [optional] 
**Role** | Pointer to [**DeviceWithConfigContextRequestRole**](DeviceWithConfigContextRequestRole.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableDeviceTypeRequestDefaultPlatform**](DeviceTypeRequestDefaultPlatform.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | Pointer to [**DeviceWithConfigContextRequestSite**](DeviceWithConfigContextRequestSite.md) |  | [optional] 
**Location** | Pointer to [**NullableDeviceWithConfigContextRequestLocation**](DeviceWithConfigContextRequestLocation.md) |  | [optional] 
**Rack** | Pointer to [**NullableDeviceWithConfigContextRequestRack**](DeviceWithConfigContextRequestRack.md) |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to [**NullableRackFace1**](RackFace1.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Status** | Pointer to [**DeviceStatusValue**](DeviceStatusValue.md) |  | [optional] 
**Airflow** | Pointer to [**NullableDeviceTypeRequestAirflow**](DeviceTypeRequestAirflow.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**OobIp** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**Cluster** | Pointer to [**NullableDeviceWithConfigContextRequestCluster**](DeviceWithConfigContextRequestCluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableDeviceWithConfigContextRequestVirtualChassis**](DeviceWithConfigContextRequestVirtualChassis.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableDeviceRoleRequestConfigTemplate**](DeviceRoleRequestConfigTemplate.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableDeviceWithConfigContextRequest

`func NewPatchedWritableDeviceWithConfigContextRequest() *PatchedWritableDeviceWithConfigContextRequest`

NewPatchedWritableDeviceWithConfigContextRequest instantiates a new PatchedWritableDeviceWithConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeviceWithConfigContextRequestWithDefaults

`func NewPatchedWritableDeviceWithConfigContextRequestWithDefaults() *PatchedWritableDeviceWithConfigContextRequest`

NewPatchedWritableDeviceWithConfigContextRequestWithDefaults instantiates a new PatchedWritableDeviceWithConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceType() DeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceTypeOk() (*DeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetDeviceType(v DeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRole() DeviceWithConfigContextRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRoleOk() (*DeviceWithConfigContextRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetRole(v DeviceWithConfigContextRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPlatform() DeviceTypeRequestDefaultPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPlatformOk() (*DeviceTypeRequestDefaultPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPlatform(v DeviceTypeRequestDefaultPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSite() DeviceWithConfigContextRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSiteOk() (*DeviceWithConfigContextRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetSite(v DeviceWithConfigContextRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocation() DeviceWithConfigContextRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocationOk() (*DeviceWithConfigContextRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocation(v DeviceWithConfigContextRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRack() DeviceWithConfigContextRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRackOk() (*DeviceWithConfigContextRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetRack(v DeviceWithConfigContextRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetFace() RackFace1`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetFaceOk() (*RackFace1, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetFace(v RackFace1)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### SetFaceNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetFaceNil(b bool)`

 SetFaceNil sets the value for Face to be an explicit nil

### UnsetFace
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetFace()`

UnsetFace ensures that no value is present for Face, not even an explicit nil
### GetLatitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetStatus() DeviceStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetStatusOk() (*DeviceStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetStatus(v DeviceStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAirflow() DeviceTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAirflowOk() (*DeviceTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAirflow(v DeviceTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp4() DeviceWithConfigContextRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp4Ok() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp4(v DeviceWithConfigContextRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp6() DeviceWithConfigContextRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp6Ok() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp6(v DeviceWithConfigContextRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetOobIp() DeviceWithConfigContextRequestPrimaryIp4`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetOobIpOk() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetOobIp(v DeviceWithConfigContextRequestPrimaryIp4)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCluster() DeviceWithConfigContextRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetClusterOk() (*DeviceWithConfigContextRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetCluster(v DeviceWithConfigContextRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVirtualChassis() DeviceWithConfigContextRequestVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVirtualChassisOk() (*DeviceWithConfigContextRequestVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVirtualChassis(v DeviceWithConfigContextRequestVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetConfigTemplate() DeviceRoleRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetConfigTemplateOk() (*DeviceRoleRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetConfigTemplate(v DeviceRoleRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


