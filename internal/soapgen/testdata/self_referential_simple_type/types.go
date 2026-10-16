package self_referential_simple_type

import (
	"encoding/xml"
)

// Simple types

// Char represents the char simpleType (restricting xs:int)
type Char int32

// Duration represents the duration simpleType (restricting xs:duration)
type Duration string

// MeasureWrapper represents the Measure element
type MeasureWrapper struct {
	XMLName xml.Name `xml:"http://example.com/selfref Measure"`
	Elapsed *string  `xml:"Elapsed,omitempty"`
	Code    *int32   `xml:"Code,omitempty"`
}

// MeasureResponseWrapper represents the MeasureResponse element
type MeasureResponseWrapper struct {
	XMLName xml.Name `xml:"http://example.com/selfref MeasureResponse"`
	Ok      *bool    `xml:"Ok,omitempty"`
}
