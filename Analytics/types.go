package Analytics

import (
	"gitlab.markany.wm/external/goonvif/xsd"
	"gitlab.markany.wm/external/goonvif/xsd/onvif"
)

type GetSupportedRules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetSupportedRules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type CreateRules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl CreateRules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	Rule               onvif.Config         `xml:"http://www.onvif.org/ver20/analytics/wsdl Rule"`
}

type DeleteRules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl DeleteRules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	RuleName           xsd.String           `xml:"http://www.onvif.org/ver20/analytics/wsdl RuleName"`
}

type GetRules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetRules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type GetRuleOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetRuleOptions" json"-"`
	RuleType           xsd.QName            `xml:"http://www.onvif.org/ver20/analytics/wsdl RuleType"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type ModifyRules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl ModifyRules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	Rule               onvif.Config         `xml:"http://www.onvif.org/ver20/analytics/wsdl Rule"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver20/analytics/wsdl GetServiceCapabilities" json"-"`
}

type GetSupportedAnalyticsModules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetSupportedAnalyticsModules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type GetAnalyticsModuleOptions struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetAnalyticsModuleOptions" json"-"`
	Type               xsd.QName            `xml:"http://www.onvif.org/ver20/analytics/wsdl Type"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type CreateAnalyticsModules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver10/events/wsdl CreateAnalyticsModules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	AnalyticsModule    onvif.Config         `xml:"http://www.onvif.org/ver20/analytics/wsdl AnalyticsModule"`
}

type DeleteAnalyticsModules struct {
	XMLName             string               `xml:"http://www.onvif.org/ver20/analytics/wsdl DeleteAnalyticsModules" json"-"`
	ConfigurationToken  onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	AnalyticsModuleName xsd.String           `xml:"http://www.onvif.org/ver20/analytics/wsdl AnalyticsModuleName"`
}

type GetAnalyticsModules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl GetAnalyticsModules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
}

type ModifyAnalyticsModules struct {
	XMLName            string               `xml:"http://www.onvif.org/ver20/analytics/wsdl ModifyAnalyticsModules" json"-"`
	ConfigurationToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver20/analytics/wsdl ConfigurationToken"`
	AnalyticsModule    onvif.Config         `xml:"http://www.onvif.org/ver20/analytics/wsdl AnalyticsModule"`
}
