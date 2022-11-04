package openrtb2

import "encoding/json"

// 3.2.29 Object: UserAgent
//
//Structured user agent information, which can be used when a client supports User-Agent Client Hints. If
//both device.ua and device.sua are present in the bid request, device.sua should be considered the more
//accurate representation of the device attributes. This is because the device.ua may contain a frozen or
//reduced user agent string due to deprecation of user agent strings by browsers.

type UserAgent struct {

	// Attribute:
	//   browsers
	// Type:
	//   array of BrandVersion objects; recommended
	// Description:
	//   Each BrandVersion object (see Section 3.2.30) identifies a browser or similar software component.
	//	 Implementers should send brands and versions derived from the Sec-CH-UA-Full-Version-List header*.
	Browsers []BrandVersion `json:"browsers,omitempty"`

	// Attribute:
	//   platform
	// Type:
	//   BrandVersion objects; recommended
	// Description:
	//	A BrandVersion object (see Section 3.2.30) that identifies the user agent’s
	//	execution platform / OS. Implementers should send a brand derived from the
	//	Sec-CH-UA-Platform header, and version derived from the Sec-CH-UAPlatform-Version header *
	Platform *BrandVersion `json:"platform,omitempty"`

	// Attribute:
	//   mobile
	// Type:
	//   integer
	// Description:
	//   1 if the agent prefers a “mobile” version of the content, if available, i.e.
	//	optimized for small screens or touch input. 0 if the agent prefers the “desktop”
	//	or “full” content. Implementers should derive this value from the Sec-CH-UAMobile header *.
	Mobile int8 `json:"mobile,omitempty"`

	// Attribute:
	//   architecture
	// Type:
	//   string
	// Description:
	//	Device’s major binary architecture, e.g. “x86” or “arm”. Implementers should retrieve this value from the Sec-CH-UA-Arch header*.
	Architecture string `json:"architecture,omitempty"`

	// Attribute:
	//   bitness
	// Type:
	//   string
	// Description:
	//   Device’s bitness, e.g. “64” for 64-bit architecture. Implementers should retrieve this value from the Sec-CH-UA-Bitness header*.
	Bitness string `json:"bitness,omitempty"`

	// Attribute:
	//   model
	// Type:
	//   string
	// Description:
	//   Device model. Implementers should retrieve this value from the Sec-CH-UAModel header*.
	Model string `json:"model,omitempty"`

	// Attribute:
	//   source
	// Type
	//   Integer; default 0
	// Description:
	//   The source of data used to create this object, List: User-Agent Source in AdCOM 1.0
	Source UserAgentSource `json:"source,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
