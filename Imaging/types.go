package Imaging

import (
	"gitlab.markany.wm/external/goonvif/xsd"
	"gitlab.markany.wm/external/goonvif/xsd/onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver20/imaging/wsdl GetServiceCapabilities"  json"-"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetImagingSettings"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"http://www.onvif.org/ver20/imaging/wsdl SetImagingSettings"  json"-"`
	VideoSourceToken onvif.ReferenceToken    `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"http://www.onvif.org/ver20/imaging/wsdl ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"http://www.onvif.org/ver20/imaging/wsdl ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetOptions"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl Move"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"http://www.onvif.org/ver20/imaging/wsdl Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetMoveOptions"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl Stop"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetStatus"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetPresets"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl GetCurrentPreset"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"http://www.onvif.org/ver20/imaging/wsdl SetCurrentPreset"  json"-"`
	VideoSourceToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/imaging/wsdl PresetToken"`
}
