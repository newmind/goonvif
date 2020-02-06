package Device

import (
	"gitlab.markany.wm/external/goonvif/xsd"
	"gitlab.markany.wm/external/goonvif/xsd/onvif"
)

type Service struct {
	Namespace xsd.AnyURI
	XAddr     xsd.AnyURI
	Capabilities
	Version onvif.OnvifVersion
}

type Capabilities struct {
	Any string
}

type DeviceServiceCapabilities struct {
	Network  NetworkCapabilities
	Security SecurityCapabilities
	System   SystemCapabilities
	Misc     MiscCapabilities
}

type NetworkCapabilities struct {
	IPFilter            xsd.Boolean `xml:"IPFilter,attr"`
	ZeroConfiguration   xsd.Boolean `xml:"ZeroConfiguration,attr"`
	IPVersion6          xsd.Boolean `xml:"IPVersion6,attr"`
	DynDNS              xsd.Boolean `xml:"DynDNS,attr"`
	Dot11Configuration  xsd.Boolean `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int         `xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP    xsd.Boolean `xml:"HostnameFromDHCP,attr"`
	NTP                 int         `xml:"NTP,attr"`
	DHCPv6              xsd.Boolean `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS1_0               xsd.Boolean    `xml:"TLS1_0,attr"`
	TLS1_1               xsd.Boolean    `xml:"TLS1_1,attr"`
	TLS1_2               xsd.Boolean    `xml:"TLS1_2,attr"`
	OnboardKeyGeneration xsd.Boolean    `xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig   xsd.Boolean    `xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy  xsd.Boolean    `xml:"DefaultAccessPolicy,attr"`
	Dot1X                xsd.Boolean    `xml:"Dot1X,attr"`
	RemoteUserHandling   xsd.Boolean    `xml:"RemoteUserHandling,attr"`
	X_509Token           xsd.Boolean    `xml:"X_509Token,attr"`
	SAMLToken            xsd.Boolean    `xml:"SAMLToken,attr"`
	KerberosToken        xsd.Boolean    `xml:"KerberosToken,attr"`
	UsernameToken        xsd.Boolean    `xml:"UsernameToken,attr"`
	HttpDigest           xsd.Boolean    `xml:"HttpDigest,attr"`
	RELToken             xsd.Boolean    `xml:"RELToken,attr"`
	SupportedEAPMethods  EAPMethodTypes `xml:"SupportedEAPMethods,attr"`
	MaxUsers             int            `xml:"MaxUsers,attr"`
	MaxUserNameLength    int            `xml:"MaxUserNameLength,attr"`
	MaxPasswordLength    int            `xml:"MaxPasswordLength,attr"`
}

type EAPMethodTypes struct {
	Types []int
}

type SystemCapabilities struct {
	DiscoveryResolve         xsd.Boolean          `xml:"DiscoveryResolve,attr"`
	DiscoveryBye             xsd.Boolean          `xml:"DiscoveryBye,attr"`
	RemoteDiscovery          xsd.Boolean          `xml:"RemoteDiscovery,attr"`
	SystemBackup             xsd.Boolean          `xml:"SystemBackup,attr"`
	SystemLogging            xsd.Boolean          `xml:"SystemLogging,attr"`
	FirmwareUpgrade          xsd.Boolean          `xml:"FirmwareUpgrade,attr"`
	HttpFirmwareUpgrade      xsd.Boolean          `xml:"HttpFirmwareUpgrade,attr"`
	HttpSystemBackup         xsd.Boolean          `xml:"HttpSystemBackup,attr"`
	HttpSystemLogging        xsd.Boolean          `xml:"HttpSystemLogging,attr"`
	HttpSupportInformation   xsd.Boolean          `xml:"HttpSupportInformation,attr"`
	StorageConfiguration     xsd.Boolean          `xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations int                  `xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries       int                  `xml:"GeoLocationEntries,attr"`
	AutoGeo                  onvif.StringAttrList `xml:"AutoGeo,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands onvif.StringAttrList `xml:"AuxiliaryCommands,attr"`
}

type StorageConfiguration struct {
	onvif.DeviceEntity
	Data StorageConfigurationData `xml:"http://www.onvif.org/ver10/device/wsdl Data"`
}

type StorageConfigurationData struct {
	Type       xsd.String     `xml:"type,attr" json:"type"`
	LocalPath  xsd.AnyURI     `xml:"http://www.onvif.org/ver10/device/wsdl LocalPath"`
	StorageUri xsd.AnyURI     `xml:"http://www.onvif.org/ver10/device/wsdl StorageUri"`
	User       UserCredential `xml:"http://www.onvif.org/ver10/device/wsdl User"`
	Extension  xsd.AnyURI     `xml:"http://www.onvif.org/ver10/device/wsdl Extension"`
}

type UserCredential struct {
	UserName  xsd.String  `xml:"http://www.onvif.org/ver10/device/wsdl UserName"`
	Password  xsd.String  `xml:"http://www.onvif.org/ver10/device/wsdl Password"`
	Extension xsd.AnyType `xml:"http://www.onvif.org/ver10/device/wsdl Extension"`
}

//Device main types

type GetServices struct {
	XMLName           string      `xml:"http://www.onvif.org/ver10/device/wsdl GetServices": json"-"`
	IncludeCapability xsd.Boolean `xml:"http://www.onvif.org/ver10/device/wsdl IncludeCapability"`
}

type GetServicesResponse struct {
	Service Service
}

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetServiceCapabilities": json"-"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities DeviceServiceCapabilities
}

type GetDeviceInformation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDeviceInformation": json"-"`
}

type GetDeviceInformationResponse struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareId      string
}

type SetSystemDateAndTime struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemDateAndTime": json"-"`
	DateTimeType    onvif.SetDateTimeType `xml:"http://www.onvif.org/ver10/device/wsdl DateTimeType"`
	DaylightSavings xsd.Boolean           `xml:"http://www.onvif.org/ver10/device/wsdl DaylightSavings"`
	TimeZone        onvif.TimeZone        `xml:"http://www.onvif.org/ver10/device/wsdl TimeZone"`
	UTCDateTime     onvif.DateTime        `xml:"http://www.onvif.org/ver10/device/wsdl UTCDateTime"`
}

type SetSystemDateAndTimeResponse struct {
}

type GetSystemDateAndTime struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemDateAndTime": json"-"`
}

type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime onvif.SystemDateTime
}

type SetSystemFactoryDefault struct {
	XMLName        string                   `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemFactoryDefault": json"-"`
	FactoryDefault onvif.FactoryDefaultType `xml:"http://www.onvif.org/ver10/device/wsdl FactoryDefault"`
}

type SetSystemFactoryDefaultResponse struct {
}

type UpgradeSystemFirmware struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/device/wsdl UpgradeSystemFirmware": json"-"`
	Firmware onvif.AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Firmware"`
}

type UpgradeSystemFirmwareResponse struct {
	Message string
}

type SystemReboot struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl SystemReboot": json"-"`
}

type SystemRebootResponse struct {
	Message string
}

//TODO: one or more repetitions
type RestoreSystem struct {
	XMLName     string           `xml:"http://www.onvif.org/ver10/device/wsdl RestoreSystem": json"-"`
	BackupFiles onvif.BackupFile `xml:"http://www.onvif.org/ver10/device/wsdl BackupFiles"`
}

type RestoreSystemResponse struct {
}

type GetSystemBackup struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemBackup": json"-"`
}

type GetSystemBackupResponse struct {
	BackupFiles onvif.BackupFile
}

type GetSystemLog struct {
	XMLName string              `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemLog": json"-"`
	LogType onvif.SystemLogType `xml:"http://www.onvif.org/ver10/device/wsdl LogType"`
}

type GetSystemLogResponse struct {
	SystemLog onvif.SystemLog
}

type GetSystemSupportInformation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemSupportInformation": json"-"`
}

type GetSystemSupportInformationResponse struct {
	SupportInformation onvif.SupportInformation
}

type GetScopes struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetScopes": json"-"`
}

type GetScopesResponse struct {
	Scopes onvif.Scope
}

//TODO: one or more scopes
type SetScopes struct {
	XMLName string     `xml:"http://www.onvif.org/ver10/device/wsdl SetScopes": json"-"`
	Scopes  xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl Scopes"`
}

type SetScopesResponse struct {
}

//TODO: list of scopes
type AddScopes struct {
	XMLName   string     `xml:"http://www.onvif.org/ver10/device/wsdl AddScopes": json"-"`
	ScopeItem xsd.AnyURI `xml:"http://www.onvif.org/ver10/device/wsdl ScopeItem"`
}

type AddScopesResponse struct {
}

//TODO: One or more repetitions
type RemoveScopes struct {
	XMLName   string     `xml:"http://www.onvif.org/ver10/device/wsdl RemoveScopes": json"-"`
	ScopeItem xsd.AnyURI `xml:"http://www.onvif.org/ver10/schema ScopeItem"`
}

type RemoveScopesResponse struct {
	ScopeItem xsd.AnyURI
}

type GetDiscoveryMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDiscoveryMode": json"-"`
}

type GetDiscoveryModeResponse struct {
	DiscoveryMode onvif.DiscoveryMode
}

type SetDiscoveryMode struct {
	XMLName       string              `xml:"http://www.onvif.org/ver10/device/wsdl SetDiscoveryMode": json"-"`
	DiscoveryMode onvif.DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl DiscoveryMode"`
}

type SetDiscoveryModeResponse struct {
}

type GetRemoteDiscoveryMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteDiscoveryMode": json"-"`
}

type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode onvif.DiscoveryMode
}

type SetRemoteDiscoveryMode struct {
	XMLName             string              `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteDiscoveryMode": json"-"`
	RemoteDiscoveryMode onvif.DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl RemoteDiscoveryMode"`
}

type SetRemoteDiscoveryModeResponse struct {
}

type GetDPAddresses struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDPAddresses": json"-"`
}

type GetDPAddressesResponse struct {
	DPAddress onvif.NetworkHost
}

type SetDPAddresses struct {
	XMLName   string            `xml:"http://www.onvif.org/ver10/device/wsdl SetDPAddresses": json"-"`
	DPAddress onvif.NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl DPAddress"`
}

type SetDPAddressesResponse struct {
}

type GetEndpointReference struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetEndpointReference": json"-"`
}

type GetEndpointReferenceResponse struct {
	GUID string
}

type GetRemoteUser struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteUser": json"-"`
}

type GetRemoteUserResponse struct {
	RemoteUser onvif.RemoteUser
}

type SetRemoteUser struct {
	XMLName    string           `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteUser": json"-"`
	RemoteUser onvif.RemoteUser `xml:"http://www.onvif.org/ver10/device/wsdl RemoteUser"`
}

type SetRemoteUserResponse struct {
}

type GetUsers struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetUsers": json"-"`
}

type GetUsersResponse struct {
	User onvif.User
}

//TODO: List of users
type CreateUsers struct {
	XMLName string     `xml:"http://www.onvif.org/ver10/device/wsdl CreateUsers": json"-"`
	User    onvif.User `xml:"http://www.onvif.org/ver10/device/wsdl User,omitempty"`
}

type CreateUsersResponse struct {
}

//TODO: one or more Username
type DeleteUsers struct {
	XMLName  xsd.String `xml:"http://www.onvif.org/ver10/device/wsdl DeleteUsers": json"-"`
	Username xsd.String `xml:"http://www.onvif.org/ver10/device/wsdl Username"`
}

type DeleteUsersResponse struct {
}

type SetUser struct {
	XMLName string     `xml:"http://www.onvif.org/ver10/device/wsdl SetUser": json"-"`
	User    onvif.User `xml:"http://www.onvif.org/ver10/device/wsdl User"`
}

type SetUserResponse struct {
}

type GetWsdlUrl struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetWsdlUrl": json"-"`
}

type GetWsdlUrlResponse struct {
	WsdlUrl xsd.AnyURI
}

type GetCapabilities struct {
	XMLName  string                   `xml:"http://www.onvif.org/ver10/device/wsdl GetCapabilities": json"-"`
	Category onvif.CapabilityCategory `xml:"http://www.onvif.org/ver10/device/wsdl Category"`
}

type GetCapabilitiesResponse struct {
	Capabilities onvif.Capabilities
}

type GetHostname struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetHostname": json"-"`
}

type GetHostnameResponse struct {
	HostnameInformation onvif.HostnameInformation
}

type SetHostname struct {
	XMLName string    `xml:"http://www.onvif.org/ver10/device/wsdl SetHostname": json"-"`
	Name    xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl Name"`
}

type SetHostnameResponse struct {
}

type SetHostnameFromDHCP struct {
	XMLName  string      `xml:"http://www.onvif.org/ver10/device/wsdl SetHostnameFromDHCP": json"-"`
	FromDHCP xsd.Boolean `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
}

type SetHostnameFromDHCPResponse struct {
	RebootNeeded xsd.Boolean
}

type GetDNS struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDNS": json"-"`
}

type GetDNSResponse struct {
	DNSInformation onvif.DNSInformation
}

type SetDNS struct {
	XMLName      string          `xml:"http://www.onvif.org/ver10/device/wsdl SetDNS": json"-"`
	FromDHCP     xsd.Boolean     `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
	SearchDomain xsd.Token       `xml:"http://www.onvif.org/ver10/device/wsdl SearchDomain"`
	DNSManual    onvif.IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl DNSManual"`
}

type SetDNSResponse struct {
}

type GetNTP struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNTP": json"-"`
}

type GetNTPResponse struct {
	NTPInformation onvif.NTPInformation
}

type SetNTP struct {
	XMLName   string            `xml:"http://www.onvif.org/ver10/device/wsdl SetNTP": json"-"`
	FromDHCP  xsd.Boolean       `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP"`
	NTPManual onvif.NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl NTPManual"`
}

type SetNTPResponse struct {
}

type GetDynamicDNS struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDynamicDNS": json"-"`
}

type GetDynamicDNSResponse struct {
	DynamicDNSInformation onvif.DynamicDNSInformation
}

type SetDynamicDNS struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl SetDynamicDNS": json"-"`
	Type    onvif.DynamicDNSType `xml:"http://www.onvif.org/ver10/device/wsdl Type"`
	Name    onvif.DNSName        `xml:"http://www.onvif.org/ver10/device/wsdl Name"`
	TTL     xsd.Duration         `xml:"http://www.onvif.org/ver10/device/wsdl TTL"`
}

type SetDynamicDNSResponse struct {
}

type GetNetworkInterfaces struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkInterfaces": json"-"`
}

type GetNetworkInterfacesResponse struct {
	NetworkInterfaces onvif.NetworkInterface
}

type SetNetworkInterfaces struct {
	XMLName          string                                 `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkInterfaces": json"-"`
	InterfaceToken   onvif.ReferenceToken                   `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
	NetworkInterface onvif.NetworkInterfaceSetConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl NetworkInterface"`
}

type SetNetworkInterfacesResponse struct {
	RebootNeeded xsd.Boolean
}

type GetNetworkProtocols struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkProtocols": json"-"`
}

type GetNetworkProtocolsResponse struct {
	NetworkProtocols onvif.NetworkProtocol
}

type SetNetworkProtocols struct {
	XMLName          string                `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkProtocols": json"-"`
	NetworkProtocols onvif.NetworkProtocol `xml:"http://www.onvif.org/ver10/device/wsdl NetworkProtocols"`
}

type SetNetworkProtocolsResponse struct {
}

type GetNetworkDefaultGateway struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkDefaultGateway": json"-"`
}

type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway onvif.NetworkGateway
}

type SetNetworkDefaultGateway struct {
	XMLName     string            `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkDefaultGateway": json"-"`
	IPv4Address onvif.IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address"`
	IPv6Address onvif.IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address"`
}

type SetNetworkDefaultGatewayResponse struct {
}

type GetZeroConfiguration struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetZeroConfiguration": json"-"`
}

type GetZeroConfigurationResponse struct {
	ZeroConfiguration onvif.NetworkZeroConfiguration
}

type SetZeroConfiguration struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl SetZeroConfiguration": json"-"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
	Enabled        xsd.Boolean          `xml:"http://www.onvif.org/ver10/device/wsdl Enabled"`
}

type SetZeroConfigurationResponse struct {
}

type GetIPAddressFilter struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetIPAddressFilter": json"-"`
}

type GetIPAddressFilterResponse struct {
	IPAddressFilter onvif.IPAddressFilter
}

type SetIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl SetIPAddressFilter": json"-"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter"`
}

type SetIPAddressFilterResponse struct {
}

//This operation adds an IP filter address to a device.
//If the device supports device access control based on
//IP filtering rules (denied or accepted ranges of IP addresses),
//the device shall support adding of IP filtering addresses through
//the AddIPAddressFilter command.
type AddIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl AddIPAddressFilter": json"-"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter"`
}

type AddIPAddressFilterResponse struct {
}

type RemoveIPAddressFilter struct {
	XMLName         string                `xml:"http://www.onvif.org/ver10/device/wsdl RemoveIPAddressFilter": json"-"`
	IPAddressFilter onvif.IPAddressFilter `xml:"http://www.onvif.org/ver10/schema IPAddressFilter"`
}

type RemoveIPAddressFilterResponse struct {
}

type GetAccessPolicy struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetAccessPolicy": json"-"`
}

type GetAccessPolicyResponse struct {
	PolicyFile onvif.BinaryData
}

type SetAccessPolicy struct {
	XMLName    string           `xml:"http://www.onvif.org/ver10/device/wsdl SetAccessPolicy": json"-"`
	PolicyFile onvif.BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl PolicyFile"`
}

type SetAccessPolicyResponse struct {
}

type CreateCertificate struct {
	XMLName        string       `xml:"http://www.onvif.org/ver10/device/wsdl CreateCertificate": json"-"`
	CertificateID  xsd.Token    `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`
	Subject        string       `xml:"http://www.onvif.org/ver10/device/wsdl Subject,omitempty"`
	ValidNotBefore xsd.DateTime `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotBefore,omitempty"`
	ValidNotAfter  xsd.DateTime `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotAfter,omitempty"`
}

type CreateCertificateResponse struct {
	NvtCertificate onvif.Certificate
}

type GetCertificates struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificates": json"-"`
}

type GetCertificatesResponse struct {
	NvtCertificate onvif.Certificate
}

type GetCertificatesStatus struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificatesStatus": json"-"`
}

type GetCertificatesStatusResponse struct {
	CertificateStatus onvif.CertificateStatus
}

type SetCertificatesStatus struct {
	XMLName           string                  `xml:"http://www.onvif.org/ver10/device/wsdl SetCertificatesStatus": json"-"`
	CertificateStatus onvif.CertificateStatus `xml:"http://www.onvif.org/ver10/device/wsdl CertificateStatus"`
}

type SetCertificatesStatusResponse struct {
}

//TODO: List of CertificateID
type DeleteCertificates struct {
	XMLName       string    `xml:"http://www.onvif.org/ver10/device/wsdl DeleteCertificates": json"-"`
	CertificateID xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
}

type DeleteCertificatesResponse struct {
}

//TODO: Откуда onvif:data = cid:21312413412
type GetPkcs10Request struct {
	XMLName       string           `xml:"http://www.onvif.org/ver10/device/wsdl GetPkcs10Request": json"-"`
	CertificateID xsd.Token        `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
	Subject       xsd.String       `xml:"http://www.onvif.org/ver10/device/wsdl Subject"`
	Attributes    onvif.BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl Attributes"`
}

type GetPkcs10RequestResponse struct {
	Pkcs10Request onvif.BinaryData
}

//TODO: one or more NTVCertificate
type LoadCertificates struct {
	XMLName        string            `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificates": json"-"`
	NVTCertificate onvif.Certificate `xml:"http://www.onvif.org/ver10/device/wsdl NVTCertificate"`
}

type LoadCertificatesResponse struct {
}

type GetClientCertificateMode struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetClientCertificateMode": json"-"`
}

type GetClientCertificateModeResponse struct {
	Enabled xsd.Boolean
}

type SetClientCertificateMode struct {
	XMLName string      `xml:"http://www.onvif.org/ver10/device/wsdl SetClientCertificateMode": json"-"`
	Enabled xsd.Boolean `xml:"http://www.onvif.org/ver10/device/wsdl Enabled"`
}

type SetClientCertificateModeResponse struct {
}

type GetRelayOutputs struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetRelayOutputs": json"-"`
}

type GetRelayOutputsResponse struct {
	RelayOutputs onvif.RelayOutput
}

type SetRelayOutputSettings struct {
	XMLName          string                    `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputSettings": json"-"`
	RelayOutputToken onvif.ReferenceToken      `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken"`
	Properties       onvif.RelayOutputSettings `xml:"http://www.onvif.org/ver10/device/wsdl Properties"`
}

type SetRelayOutputSettingsResponse struct {
}

type SetRelayOutputState struct {
	XMLName          string                  `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputState": json"-"`
	RelayOutputToken onvif.ReferenceToken    `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken"`
	LogicalState     onvif.RelayLogicalState `xml:"http://www.onvif.org/ver10/device/wsdl LogicalState"`
}

type SetRelayOutputStateResponse struct {
}

type SendAuxiliaryCommand struct {
	XMLName          string              `xml:"http://www.onvif.org/ver10/device/wsdl SendAuxiliaryCommand": json"-"`
	AuxiliaryCommand onvif.AuxiliaryData `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommand"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse onvif.AuxiliaryData
}

type GetCACertificates struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetCACertificates": json"-"`
}

type GetCACertificatesResponse struct {
	CACertificate onvif.Certificate
}

//TODO: one or more CertificateWithPrivateKey
type LoadCertificateWithPrivateKey struct {
	XMLName                   string                          `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificateWithPrivateKey": json"-"`
	CertificateWithPrivateKey onvif.CertificateWithPrivateKey `xml:"http://www.onvif.org/ver10/device/wsdl CertificateWithPrivateKey"`
}

type LoadCertificateWithPrivateKeyResponse struct {
}

type GetCertificateInformation struct {
	XMLName       string    `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificateInformation": json"-"`
	CertificateID xsd.Token `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID"`
}

type GetCertificateInformationResponse struct {
	CertificateInformation onvif.CertificateInformation
}

type LoadCACertificates struct {
	XMLName       string            `xml:"http://www.onvif.org/ver10/device/wsdl LoadCACertificates": json"-"`
	CACertificate onvif.Certificate `xml:"http://www.onvif.org/ver10/device/wsdl CACertificate"`
}

type LoadCACertificatesResponse struct {
}

type CreateDot1XConfiguration struct {
	XMLName            string                   `xml:"http://www.onvif.org/ver10/device/wsdl CreateDot1XConfiguration": json"-"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration"`
}

type CreateDot1XConfigurationResponse struct {
}

type SetDot1XConfiguration struct {
	XMLName            string                   `xml:"http://www.onvif.org/ver10/device/wsdl SetDot1XConfiguration": json"-"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration"`
}

type SetDot1XConfigurationResponse struct {
}

type GetDot1XConfiguration struct {
	XMLName                 string               `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfiguration": json"-"`
	Dot1XConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken"`
}

type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration
}

type GetDot1XConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfigurations": json"-"`
}

type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration
}

//TODO: Zero or more Dot1XConfigurationToken
type DeleteDot1XConfiguration struct {
	XMLName                 string               `xml:"http://www.onvif.org/ver10/device/wsdl DeleteDot1XConfiguration": json"-"`
	Dot1XConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken"`
}

type DeleteDot1XConfigurationResponse struct {
}

type GetDot11Capabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Capabilities": json"-"`
}

type GetDot11CapabilitiesResponse struct {
	Capabilities onvif.Dot11Capabilities
}

type GetDot11Status struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Status": json"-"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
}

type GetDot11StatusResponse struct {
	Status onvif.Dot11Status
}

type ScanAvailableDot11Networks struct {
	XMLName        string               `xml:"http://www.onvif.org/ver10/device/wsdl ScanAvailableDot11Networks": json"-"`
	InterfaceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken"`
}

type ScanAvailableDot11NetworksResponse struct {
	Networks onvif.Dot11AvailableNetworks
}

type GetSystemUris struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemUris": json"-"`
}

type GetSystemUrisResponse struct {
	SystemLogUris   onvif.SystemLogUriList
	SupportInfoUri  xsd.AnyURI
	SystemBackupUri xsd.AnyURI
	Extension       xsd.AnyType
}

type StartFirmwareUpgrade struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl StartFirmwareUpgrade": json"-"`
}

type StartFirmwareUpgradeResponse struct {
	UploadUri        xsd.AnyURI
	UploadDelay      xsd.Duration
	ExpectedDownTime xsd.Duration
}

type StartSystemRestore struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl StartSystemRestore": json"-"`
}

type StartSystemRestoreResponse struct {
	UploadUri        xsd.AnyURI
	ExpectedDownTime xsd.Duration
}

type GetStorageConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfigurations": json"-"`
}

type GetStorageConfigurationsResponse struct {
	StorageConfigurations StorageConfiguration
}

type CreateStorageConfiguration struct {
	XMLName              string `xml:"http://www.onvif.org/ver10/device/wsdl CreateStorageConfiguration": json"-"`
	StorageConfiguration StorageConfigurationData
}

type CreateStorageConfigurationResponse struct {
	Token onvif.ReferenceToken
}

type GetStorageConfiguration struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfiguration": json"-"`
	Token   onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token"`
}

type GetStorageConfigurationResponse struct {
	StorageConfiguration StorageConfiguration
}

type SetStorageConfiguration struct {
	XMLName              string               `xml:"http://www.onvif.org/ver10/device/wsdl SetStorageConfiguration": json"-"`
	StorageConfiguration StorageConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfiguration"`
}

type SetStorageConfigurationResponse struct {
}

type DeleteStorageConfiguration struct {
	XMLName string               `xml:"http://www.onvif.org/ver10/device/wsdl DeleteStorageConfiguration": json"-"`
	Token   onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token"`
}

type DeleteStorageConfigurationResponse struct {
}

type GetGeoLocation struct {
	XMLName string `xml:"http://www.onvif.org/ver10/device/wsdl GetGeoLocation": json"-"`
}

type GetGeoLocationResponse struct {
	Location onvif.LocationEntity
}

//TODO: one or more Location
type SetGeoLocation struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/device/wsdl SetGeoLocation": json"-"`
	Location onvif.LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location"`
}

type SetGeoLocationResponse struct {
}

type DeleteGeoLocation struct {
	XMLName  string               `xml:"http://www.onvif.org/ver10/device/wsdl DeleteGeoLocation": json"-"`
	Location onvif.LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location"`
}

type DeleteGeoLocationResponse struct {
}
