package inline_enum_dedup_stability

import (
	"encoding/xml"
)

// RawXML captures raw XML content for untyped elements.
type RawXML []byte

// Inline enumeration types

// DoSomethingRequest_Mode represents an inline enumeration type
type DoSomethingRequest_Mode string

// DoSomethingRequest_Mode enumeration values
const (
	DoSomethingRequest_ModeAlpha DoSomethingRequest_Mode = "alpha"
	DoSomethingRequest_ModeBeta  DoSomethingRequest_Mode = "beta"
)

// String returns the string representation of DoSomethingRequest_Mode
func (e DoSomethingRequest_Mode) String() string {
	return string(e)
}

// IsValid returns true if the DoSomethingRequest_Mode value is valid
func (e DoSomethingRequest_Mode) IsValid() bool {
	switch e {
	case DoSomethingRequest_ModeAlpha, DoSomethingRequest_ModeBeta:
		return true
	default:
		return false
	}
}

// DoSomethingRequestWrapper represents the DoSomethingRequest element
type DoSomethingRequestWrapper struct {
	XMLName xml.Name                `xml:"http://example.com/dedupstability DoSomethingRequest"`
	Mode    DoSomethingRequest_Mode `xml:"Mode"`
}

// DoSomethingResponseWrapper represents the DoSomethingResponse element
type DoSomethingResponseWrapper struct {
	XMLName xml.Name                `xml:"http://example.com/dedupstability DoSomethingResponse"`
	Mode    DoSomethingRequest_Mode `xml:"Mode"`
}
