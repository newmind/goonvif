package PTZ

import (
	"gitlab.markany.wm/external/goonvif/xsd"
	"gitlab.markany.wm/external/goonvif/xsd/onvif"
)

type Capabilities struct {
	EFlip                       xsd.Boolean `swaggertype:"boolean" xml:"EFlip,attr"`
	Reverse                     xsd.Boolean `swaggertype:"boolean" xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `swaggertype:"boolean" xml:"GetCompatibleConfigurations,attr"`
	MoveStatus                  xsd.Boolean `swaggertype:"boolean" xml:"MoveStatus,attr"`
	StatusPosition              xsd.Boolean `swaggertype:"boolean" xml:"StatusPosition,attr"`
}

//PTZ main types

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver20/ptz/wsdl GetServiceCapabilities": json"-"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetNodes struct {
	XMLName string `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNodes": json"-"`
}

type GetNodesResponse struct {
	PTZNode []onvif.PTZNode
}

type GetNode struct {
	XMLName   string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNode": json"-"`
	NodeToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl NodeToken"`
}

type GetNodeResponse struct {
	PTZNode onvif.PTZNode
}

type GetConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfiguration": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetConfigurationResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}

type GetConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurations": json"-"`
}

type GetConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}

type SetConfiguration struct {
	XMLName          string                 `xml:"http://www.onvif.org/ver20/ptz/wsdl SetConfiguration": json"-"`
	PTZConfiguration onvif.PTZConfiguration `xml:"http://www.onvif.org/ver20/ptz/wsdl PTZConfiguration"`
	ForcePersistence xsd.Boolean            `swaggertype:"boolean" xml:"http://www.onvif.org/ver20/ptz/wsdl ForcePersistence"`
}

type SetConfigurationResponse struct {
}

type GetConfigurationOptions struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurationOptions": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions onvif.PTZConfigurationOptions
}

type SendAuxiliaryCommand struct {
	XMLName       string               `xml:"http://www.onvif.org/ver20/ptz/wsdl SendAuxiliaryCommand": json"-"`
	ProfileToken  onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	AuxiliaryData onvif.AuxiliaryData  `xml:"http://www.onvif.org/ver20/ptz/wsdl AuxiliaryData"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse onvif.AuxiliaryData
}

type GetPresets struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresets": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetPresetsResponse struct {
	Preset []onvif.PTZPreset
}

type SetPreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl SetPreset": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetName   xsd.String           `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetName"`
	PresetToken  onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetToken"`
}

type SetPresetResponse struct {
	PresetToken onvif.ReferenceToken
}

type RemovePreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePreset": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetToken  onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetToken"`
}

type RemovePresetResponse struct {
}

type GotoPreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoPreset": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetToken  onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetToken"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Speed"`
}

type GotoPresetResponse struct {
}

type GotoHomePosition struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoHomePosition": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Speed"`
}

type GotoHomePositionResponse struct {
}

type SetHomePosition struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl SetHomePosition": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type SetHomePositionResponse struct {
}

type ContinuousMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl ContinuousMove": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	Velocity     onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Velocity"`
	Timeout      xsd.Duration         `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl Timeout"`
}

type ContinuousMoveResponse struct {
}

type RelativeMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl RelativeMove": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	Translation  onvif.PTZVector      `xml:"http://www.onvif.org/ver20/ptz/wsdl Translation"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Speed"`
}

type RelativeMoveResponse struct {
}

type GetStatus struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetStatus": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetStatusResponse struct {
	XMLName   string          `xml:"http://www.onvif.org/ver20/ptz/wsdl GetStatusResponse": json"-"`
	PTZStatus onvif.PTZStatus `xml:"http://www.onvif.org/ver20/ptz/wsdl PTZStatus"`
}

type AbsoluteMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl AbsoluteMove": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	Position     onvif.PTZVector      `xml:"http://www.onvif.org/ver20/ptz/wsdl Position"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Speed"`
}

type AbsoluteMoveResponse struct {
}

type GeoMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GeoMove": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	Target       onvif.GeoLocation    `xml:"http://www.onvif.org/ver20/ptz/wsdl Target"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver20/ptz/wsdl Speed"`
	AreaHeight   xsd.Float            `swaggertype:"number" xml:"http://www.onvif.org/ver20/ptz/wsdl AreaHeight"`
	AreaWidth    xsd.Float            `swaggertype:"number" xml:"http://www.onvif.org/ver20/ptz/wsdl AreaWidth"`
}

type GeoMoveResponse struct {
}

type Stop struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl Stop": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PanTilt      xsd.Boolean          `swaggertype:"boolean" xml:"http://www.onvif.org/ver20/ptz/wsdl PanTilt"`
	Zoom         xsd.Boolean          `swaggertype:"boolean" xml:"http://www.onvif.org/ver20/ptz/wsdl Zoom"`
}

type StopResponse struct {
}

type GetPresetTours struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTours": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetPresetToursResponse struct {
	PresetTour []onvif.PresetTour
}

type GetPresetTour struct {
	XMLName         string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTour": json"-"`
	ProfileToken    onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetTourToken"`
}

type GetPresetTourResponse struct {
	PresetTour onvif.PresetTour
}

type GetPresetTourOptions struct {
	XMLName         string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTourOptions": json"-"`
	ProfileToken    onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetTourToken"`
}

type GetPresetTourOptionsResponse struct {
	Options onvif.PTZPresetTourOptions
}

type CreatePresetTour struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl CreatePresetTour": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type CreatePresetTourResponse struct {
	PresetTourToken onvif.ReferenceToken
}

type ModifyPresetTour struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl ModifyPresetTour": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetTour   onvif.PresetTour     `xml:"http://www.onvif.org/ver20/ptz/wsdl PresetTour"`
}

type ModifyPresetTourResponse struct {
}

type OperatePresetTour struct {
	XMLName         string                       `xml:"http://www.onvif.org/ver20/ptz/wsdl OperatePresetTour": json"-"`
	ProfileToken    onvif.ReferenceToken         `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken         `swaggertype:"string" xml:"http://www.onvif.org/ver10/schema PresetTourToken"`
	Operation       onvif.PTZPresetTourOperation `xml:"http://www.onvif.org/ver10/schema Operation"`
}

type OperatePresetTourResponse struct {
}

type RemovePresetTour struct {
	XMLName         string               `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePresetTour": json"-"`
	ProfileToken    onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl PresetTourToken"`
}

type RemovePresetTourResponse struct {
}

type GetCompatibleConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver20/ptz/wsdl GetCompatibleConfigurations": json"-"`
	ProfileToken onvif.ReferenceToken `swaggertype:"string" xml:"http://www.onvif.org/ver20/ptz/wsdl ProfileToken"`
}

type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}
