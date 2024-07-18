package openrtb2

// List: Plcmt Subtypes - Video
//
// The following table lists the various types of video placements in accordance with updated IAB Digital Video Guidelines.
// To be sent using plcmt attribute in Object:Video.
// Please refer to the implementation guide for examples and information on how to use the updated signals.
type VideoPlcmtType int8

const (
	VideoPlcmtTypeTypeInstream				VideoPlcmtType = 1 // Pre-roll, mid-roll, and post-roll ads that are played before, during or after the streaming video content that the consumer has requested. Instream video must be set to “sound on” by default at player start, or have explicitly clear user intent to watch the video content. While there may be other content surrounding the player, the video content must be the focus of the user’s visit. It should remain the primary content on the page and the only video player in-view capable of audio when playing. If the player converts to floating/sticky subsequent ad calls should accurately convey the updated player size.
	VideoPlcmtTypeTypeAccompanyingContent	VideoPlcmtType = 2 // Pre-roll, mid-roll, and post-roll ads that are played before, during, or after streaming video content. The video player loads and plays before, between, or after paragraphs of text or graphical content, and starts playing only when it enters the viewport. Accompanying content should only start playback upon entering the viewport. It may convert to a floating/sticky player as it scrolls off the page.
	VideoPlcmtTypeTypeInterstitial			VideoPlcmtType = 3 // Video ads that are played without video content. During playback, it must be the primary focus of the page and take up the majority of the viewport and cannot be scrolled out of view. This can be in placements like in-app video or slideshows.
	VideoPlcmtTypeTypeNoContentStandalone	VideoPlcmtType = 4 // Video ads that are played without streaming video content. This can be in placements like slideshows, native feeds, in-content or sticky/floating.
)
