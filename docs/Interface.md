# Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Vdcs** | Pointer to [**[]VirtualDeviceContext**](VirtualDeviceContext.md) |  | [optional] 
**Module** | Pointer to [**NullableBriefModule**](BriefModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**InterfaceType**](InterfaceType.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NullableNestedInterface**](NestedInterface.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterface**](NestedInterface.md) |  | [optional] 
**Lag** | Pointer to [**NullableNestedInterface**](NestedInterface.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] [readonly] 
**PrimaryMacAddress** | Pointer to [**NullableBriefMACAddress**](BriefMACAddress.md) |  | [optional] 
**MacAddresses** | Pointer to [**[]BriefMACAddress**](BriefMACAddress.md) |  | [optional] [readonly] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Duplex** | Pointer to [**NullableInterfaceDuplex**](InterfaceDuplex.md) |  | [optional] 
**Wwn** | Pointer to **NullableString** |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**InterfaceMode**](InterfaceMode.md) |  | [optional] 
**RfRole** | Pointer to [**InterfaceRfRole**](InterfaceRfRole.md) |  | [optional] 
**RfChannel** | Pointer to [**InterfaceRfChannel**](InterfaceRfChannel.md) |  | [optional] 
**PoeMode** | Pointer to [**InterfacePoeMode**](InterfacePoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**InterfacePoeType**](InterfacePoeType.md) |  | [optional] 
**RfChannelFrequency** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**RfChannelWidth** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**TxPower** | Pointer to **NullableInt32** |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBriefVLAN**](BriefVLAN.md) |  | [optional] 
**TaggedVlans** | Pointer to [**[]VLAN**](VLAN.md) |  | [optional] 
**QinqSvlan** | Pointer to [**NullableBriefVLAN**](BriefVLAN.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableBriefVLANTranslationPolicy**](BriefVLANTranslationPolicy.md) |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | Pointer to [**NullableBriefCable**](BriefCable.md) |  | [optional] [readonly] 
**CableEnd** | Pointer to **string** |  | [optional] [readonly] 
**WirelessLink** | Pointer to [**NullableNestedWirelessLink**](NestedWirelessLink.md) |  | [optional] [readonly] 
**LinkPeers** | **[]interface{}** |  | [readonly] 
**LinkPeersType** | Pointer to **NullableString** | Return the type of the peer link terminations, or None. | [optional] [readonly] 
**WirelessLans** | Pointer to [**[]WirelessLAN**](WirelessLAN.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBriefVRF**](BriefVRF.md) |  | [optional] 
**L2vpnTermination** | Pointer to [**NullableBriefL2VPNTermination**](BriefL2VPNTermination.md) |  | [optional] [readonly] 
**ConnectedEndpoints** | Pointer to **[]interface{}** |  | [optional] [readonly] 
**ConnectedEndpointsType** | Pointer to **NullableString** |  | [optional] [readonly] 
**ConnectedEndpointsReachable** | **bool** |  | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CountIpaddresses** | **int32** |  | [readonly] 
**CountFhrpGroups** | **int32** |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewInterface

`func NewInterface(id int32, url string, display string, device BriefDevice, name string, type_ InterfaceType, linkPeers []interface{}, connectedEndpointsReachable bool, countIpaddresses int32, countFhrpGroups int32, occupied bool, ) *Interface`

NewInterface instantiates a new Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceWithDefaults

`func NewInterfaceWithDefaults() *Interface`

NewInterfaceWithDefaults instantiates a new Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Interface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interface) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Interface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Interface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Interface) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *Interface) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *Interface) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *Interface) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *Interface) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *Interface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Interface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Interface) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *Interface) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Interface) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Interface) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetVdcs

`func (o *Interface) GetVdcs() []VirtualDeviceContext`

GetVdcs returns the Vdcs field if non-nil, zero value otherwise.

### GetVdcsOk

`func (o *Interface) GetVdcsOk() (*[]VirtualDeviceContext, bool)`

GetVdcsOk returns a tuple with the Vdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcs

`func (o *Interface) SetVdcs(v []VirtualDeviceContext)`

SetVdcs sets Vdcs field to given value.

### HasVdcs

`func (o *Interface) HasVdcs() bool`

HasVdcs returns a boolean if a field has been set.

### GetModule

`func (o *Interface) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Interface) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Interface) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *Interface) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *Interface) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *Interface) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *Interface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Interface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Interface) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Interface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Interface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Interface) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Interface) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *Interface) GetType() InterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interface) GetTypeOk() (*InterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interface) SetType(v InterfaceType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *Interface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Interface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Interface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Interface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *Interface) GetParent() NestedInterface`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Interface) GetParentOk() (*NestedInterface, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Interface) SetParent(v NestedInterface)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Interface) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *Interface) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Interface) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *Interface) GetBridge() NestedInterface`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *Interface) GetBridgeOk() (*NestedInterface, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *Interface) SetBridge(v NestedInterface)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *Interface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *Interface) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *Interface) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *Interface) GetLag() NestedInterface`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *Interface) GetLagOk() (*NestedInterface, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *Interface) SetLag(v NestedInterface)`

SetLag sets Lag field to given value.

### HasLag

`func (o *Interface) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *Interface) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *Interface) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetMtu

`func (o *Interface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Interface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Interface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Interface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *Interface) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *Interface) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *Interface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Interface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Interface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Interface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *Interface) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *Interface) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetPrimaryMacAddress

`func (o *Interface) GetPrimaryMacAddress() BriefMACAddress`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *Interface) GetPrimaryMacAddressOk() (*BriefMACAddress, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *Interface) SetPrimaryMacAddress(v BriefMACAddress)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *Interface) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *Interface) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *Interface) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetMacAddresses

`func (o *Interface) GetMacAddresses() []BriefMACAddress`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *Interface) GetMacAddressesOk() (*[]BriefMACAddress, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *Interface) SetMacAddresses(v []BriefMACAddress)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *Interface) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.

### SetMacAddressesNil

`func (o *Interface) SetMacAddressesNil(b bool)`

 SetMacAddressesNil sets the value for MacAddresses to be an explicit nil

### UnsetMacAddresses
`func (o *Interface) UnsetMacAddresses()`

UnsetMacAddresses ensures that no value is present for MacAddresses, not even an explicit nil
### GetSpeed

`func (o *Interface) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Interface) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Interface) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Interface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *Interface) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *Interface) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDuplex

`func (o *Interface) GetDuplex() InterfaceDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *Interface) GetDuplexOk() (*InterfaceDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *Interface) SetDuplex(v InterfaceDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *Interface) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### SetDuplexNil

`func (o *Interface) SetDuplexNil(b bool)`

 SetDuplexNil sets the value for Duplex to be an explicit nil

### UnsetDuplex
`func (o *Interface) UnsetDuplex()`

UnsetDuplex ensures that no value is present for Duplex, not even an explicit nil
### GetWwn

`func (o *Interface) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *Interface) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *Interface) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *Interface) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### SetWwnNil

`func (o *Interface) SetWwnNil(b bool)`

 SetWwnNil sets the value for Wwn to be an explicit nil

### UnsetWwn
`func (o *Interface) UnsetWwn()`

UnsetWwn ensures that no value is present for Wwn, not even an explicit nil
### GetMgmtOnly

`func (o *Interface) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *Interface) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *Interface) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *Interface) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *Interface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *Interface) GetMode() InterfaceMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Interface) GetModeOk() (*InterfaceMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Interface) SetMode(v InterfaceMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Interface) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRfRole

`func (o *Interface) GetRfRole() InterfaceRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *Interface) GetRfRoleOk() (*InterfaceRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *Interface) SetRfRole(v InterfaceRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *Interface) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### GetRfChannel

`func (o *Interface) GetRfChannel() InterfaceRfChannel`

GetRfChannel returns the RfChannel field if non-nil, zero value otherwise.

### GetRfChannelOk

`func (o *Interface) GetRfChannelOk() (*InterfaceRfChannel, bool)`

GetRfChannelOk returns a tuple with the RfChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannel

`func (o *Interface) SetRfChannel(v InterfaceRfChannel)`

SetRfChannel sets RfChannel field to given value.

### HasRfChannel

`func (o *Interface) HasRfChannel() bool`

HasRfChannel returns a boolean if a field has been set.

### GetPoeMode

`func (o *Interface) GetPoeMode() InterfacePoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *Interface) GetPoeModeOk() (*InterfacePoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *Interface) SetPoeMode(v InterfacePoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *Interface) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPoeType

`func (o *Interface) GetPoeType() InterfacePoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *Interface) GetPoeTypeOk() (*InterfacePoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *Interface) SetPoeType(v InterfacePoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *Interface) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### GetRfChannelFrequency

`func (o *Interface) GetRfChannelFrequency() float64`

GetRfChannelFrequency returns the RfChannelFrequency field if non-nil, zero value otherwise.

### GetRfChannelFrequencyOk

`func (o *Interface) GetRfChannelFrequencyOk() (*float64, bool)`

GetRfChannelFrequencyOk returns a tuple with the RfChannelFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelFrequency

`func (o *Interface) SetRfChannelFrequency(v float64)`

SetRfChannelFrequency sets RfChannelFrequency field to given value.

### HasRfChannelFrequency

`func (o *Interface) HasRfChannelFrequency() bool`

HasRfChannelFrequency returns a boolean if a field has been set.

### SetRfChannelFrequencyNil

`func (o *Interface) SetRfChannelFrequencyNil(b bool)`

 SetRfChannelFrequencyNil sets the value for RfChannelFrequency to be an explicit nil

### UnsetRfChannelFrequency
`func (o *Interface) UnsetRfChannelFrequency()`

UnsetRfChannelFrequency ensures that no value is present for RfChannelFrequency, not even an explicit nil
### GetRfChannelWidth

`func (o *Interface) GetRfChannelWidth() float64`

GetRfChannelWidth returns the RfChannelWidth field if non-nil, zero value otherwise.

### GetRfChannelWidthOk

`func (o *Interface) GetRfChannelWidthOk() (*float64, bool)`

GetRfChannelWidthOk returns a tuple with the RfChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelWidth

`func (o *Interface) SetRfChannelWidth(v float64)`

SetRfChannelWidth sets RfChannelWidth field to given value.

### HasRfChannelWidth

`func (o *Interface) HasRfChannelWidth() bool`

HasRfChannelWidth returns a boolean if a field has been set.

### SetRfChannelWidthNil

`func (o *Interface) SetRfChannelWidthNil(b bool)`

 SetRfChannelWidthNil sets the value for RfChannelWidth to be an explicit nil

### UnsetRfChannelWidth
`func (o *Interface) UnsetRfChannelWidth()`

UnsetRfChannelWidth ensures that no value is present for RfChannelWidth, not even an explicit nil
### GetTxPower

`func (o *Interface) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *Interface) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *Interface) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *Interface) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### SetTxPowerNil

`func (o *Interface) SetTxPowerNil(b bool)`

 SetTxPowerNil sets the value for TxPower to be an explicit nil

### UnsetTxPower
`func (o *Interface) UnsetTxPower()`

UnsetTxPower ensures that no value is present for TxPower, not even an explicit nil
### GetUntaggedVlan

`func (o *Interface) GetUntaggedVlan() BriefVLAN`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *Interface) GetUntaggedVlanOk() (*BriefVLAN, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *Interface) SetUntaggedVlan(v BriefVLAN)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *Interface) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *Interface) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *Interface) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *Interface) GetTaggedVlans() []VLAN`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *Interface) GetTaggedVlansOk() (*[]VLAN, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *Interface) SetTaggedVlans(v []VLAN)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *Interface) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *Interface) GetQinqSvlan() BriefVLAN`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *Interface) GetQinqSvlanOk() (*BriefVLAN, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *Interface) SetQinqSvlan(v BriefVLAN)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *Interface) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *Interface) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *Interface) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *Interface) GetVlanTranslationPolicy() BriefVLANTranslationPolicy`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *Interface) GetVlanTranslationPolicyOk() (*BriefVLANTranslationPolicy, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *Interface) SetVlanTranslationPolicy(v BriefVLANTranslationPolicy)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *Interface) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *Interface) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *Interface) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetMarkConnected

`func (o *Interface) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *Interface) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *Interface) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *Interface) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *Interface) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *Interface) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *Interface) SetCable(v BriefCable)`

SetCable sets Cable field to given value.

### HasCable

`func (o *Interface) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *Interface) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *Interface) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *Interface) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *Interface) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *Interface) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.

### HasCableEnd

`func (o *Interface) HasCableEnd() bool`

HasCableEnd returns a boolean if a field has been set.

### GetWirelessLink

`func (o *Interface) GetWirelessLink() NestedWirelessLink`

GetWirelessLink returns the WirelessLink field if non-nil, zero value otherwise.

### GetWirelessLinkOk

`func (o *Interface) GetWirelessLinkOk() (*NestedWirelessLink, bool)`

GetWirelessLinkOk returns a tuple with the WirelessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLink

`func (o *Interface) SetWirelessLink(v NestedWirelessLink)`

SetWirelessLink sets WirelessLink field to given value.

### HasWirelessLink

`func (o *Interface) HasWirelessLink() bool`

HasWirelessLink returns a boolean if a field has been set.

### SetWirelessLinkNil

`func (o *Interface) SetWirelessLinkNil(b bool)`

 SetWirelessLinkNil sets the value for WirelessLink to be an explicit nil

### UnsetWirelessLink
`func (o *Interface) UnsetWirelessLink()`

UnsetWirelessLink ensures that no value is present for WirelessLink, not even an explicit nil
### GetLinkPeers

`func (o *Interface) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *Interface) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *Interface) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *Interface) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *Interface) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *Interface) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.

### HasLinkPeersType

`func (o *Interface) HasLinkPeersType() bool`

HasLinkPeersType returns a boolean if a field has been set.

### SetLinkPeersTypeNil

`func (o *Interface) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *Interface) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetWirelessLans

`func (o *Interface) GetWirelessLans() []WirelessLAN`

GetWirelessLans returns the WirelessLans field if non-nil, zero value otherwise.

### GetWirelessLansOk

`func (o *Interface) GetWirelessLansOk() (*[]WirelessLAN, bool)`

GetWirelessLansOk returns a tuple with the WirelessLans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLans

`func (o *Interface) SetWirelessLans(v []WirelessLAN)`

SetWirelessLans sets WirelessLans field to given value.

### HasWirelessLans

`func (o *Interface) HasWirelessLans() bool`

HasWirelessLans returns a boolean if a field has been set.

### GetVrf

`func (o *Interface) GetVrf() BriefVRF`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *Interface) GetVrfOk() (*BriefVRF, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *Interface) SetVrf(v BriefVRF)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *Interface) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *Interface) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *Interface) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetL2vpnTermination

`func (o *Interface) GetL2vpnTermination() BriefL2VPNTermination`

GetL2vpnTermination returns the L2vpnTermination field if non-nil, zero value otherwise.

### GetL2vpnTerminationOk

`func (o *Interface) GetL2vpnTerminationOk() (*BriefL2VPNTermination, bool)`

GetL2vpnTerminationOk returns a tuple with the L2vpnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2vpnTermination

`func (o *Interface) SetL2vpnTermination(v BriefL2VPNTermination)`

SetL2vpnTermination sets L2vpnTermination field to given value.

### HasL2vpnTermination

`func (o *Interface) HasL2vpnTermination() bool`

HasL2vpnTermination returns a boolean if a field has been set.

### SetL2vpnTerminationNil

`func (o *Interface) SetL2vpnTerminationNil(b bool)`

 SetL2vpnTerminationNil sets the value for L2vpnTermination to be an explicit nil

### UnsetL2vpnTermination
`func (o *Interface) UnsetL2vpnTermination()`

UnsetL2vpnTermination ensures that no value is present for L2vpnTermination, not even an explicit nil
### GetConnectedEndpoints

`func (o *Interface) GetConnectedEndpoints() []interface{}`

GetConnectedEndpoints returns the ConnectedEndpoints field if non-nil, zero value otherwise.

### GetConnectedEndpointsOk

`func (o *Interface) GetConnectedEndpointsOk() (*[]interface{}, bool)`

GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoints

`func (o *Interface) SetConnectedEndpoints(v []interface{})`

SetConnectedEndpoints sets ConnectedEndpoints field to given value.

### HasConnectedEndpoints

`func (o *Interface) HasConnectedEndpoints() bool`

HasConnectedEndpoints returns a boolean if a field has been set.

### SetConnectedEndpointsNil

`func (o *Interface) SetConnectedEndpointsNil(b bool)`

 SetConnectedEndpointsNil sets the value for ConnectedEndpoints to be an explicit nil

### UnsetConnectedEndpoints
`func (o *Interface) UnsetConnectedEndpoints()`

UnsetConnectedEndpoints ensures that no value is present for ConnectedEndpoints, not even an explicit nil
### GetConnectedEndpointsType

`func (o *Interface) GetConnectedEndpointsType() string`

GetConnectedEndpointsType returns the ConnectedEndpointsType field if non-nil, zero value otherwise.

### GetConnectedEndpointsTypeOk

`func (o *Interface) GetConnectedEndpointsTypeOk() (*string, bool)`

GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsType

`func (o *Interface) SetConnectedEndpointsType(v string)`

SetConnectedEndpointsType sets ConnectedEndpointsType field to given value.

### HasConnectedEndpointsType

`func (o *Interface) HasConnectedEndpointsType() bool`

HasConnectedEndpointsType returns a boolean if a field has been set.

### SetConnectedEndpointsTypeNil

`func (o *Interface) SetConnectedEndpointsTypeNil(b bool)`

 SetConnectedEndpointsTypeNil sets the value for ConnectedEndpointsType to be an explicit nil

### UnsetConnectedEndpointsType
`func (o *Interface) UnsetConnectedEndpointsType()`

UnsetConnectedEndpointsType ensures that no value is present for ConnectedEndpointsType, not even an explicit nil
### GetConnectedEndpointsReachable

`func (o *Interface) GetConnectedEndpointsReachable() bool`

GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointsReachableOk

`func (o *Interface) GetConnectedEndpointsReachableOk() (*bool, bool)`

GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsReachable

`func (o *Interface) SetConnectedEndpointsReachable(v bool)`

SetConnectedEndpointsReachable sets ConnectedEndpointsReachable field to given value.


### GetTags

`func (o *Interface) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Interface) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Interface) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Interface) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Interface) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Interface) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Interface) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Interface) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Interface) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Interface) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Interface) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Interface) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Interface) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Interface) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Interface) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Interface) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Interface) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Interface) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *Interface) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Interface) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetCountIpaddresses

`func (o *Interface) GetCountIpaddresses() int32`

GetCountIpaddresses returns the CountIpaddresses field if non-nil, zero value otherwise.

### GetCountIpaddressesOk

`func (o *Interface) GetCountIpaddressesOk() (*int32, bool)`

GetCountIpaddressesOk returns a tuple with the CountIpaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountIpaddresses

`func (o *Interface) SetCountIpaddresses(v int32)`

SetCountIpaddresses sets CountIpaddresses field to given value.


### GetCountFhrpGroups

`func (o *Interface) GetCountFhrpGroups() int32`

GetCountFhrpGroups returns the CountFhrpGroups field if non-nil, zero value otherwise.

### GetCountFhrpGroupsOk

`func (o *Interface) GetCountFhrpGroupsOk() (*int32, bool)`

GetCountFhrpGroupsOk returns a tuple with the CountFhrpGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFhrpGroups

`func (o *Interface) SetCountFhrpGroups(v int32)`

SetCountFhrpGroups sets CountFhrpGroups field to given value.


### GetOccupied

`func (o *Interface) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *Interface) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *Interface) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


