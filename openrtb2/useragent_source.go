package openrtb2

// List: User-Agent Source in AdCOM 1.0
//
// The following table lists the possible sources for User-Agent metadata.
type UserAgentSource int8

const (
	UserAgentSourceUnspecified                     UserAgentSource = 0 // Unspecified/unknown
	UserAgentSourceUserAgentClientHintsLowEntropy  UserAgentSource = 1 // User-Agent Client Hints (only low-entropy headers were available)
	UserAgentSourceUserAgentClientHintsHighEntropy UserAgentSource = 2 // User-Agent Client Hints (with high-entropy headers available)
	UserAgentSourceParsedFromUserAgentHeader       UserAgentSource = 3 // Parsed from User-Agent header (the same string carried by the ua field)
)
