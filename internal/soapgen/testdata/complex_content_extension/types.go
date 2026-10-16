package complex_content_extension

import (
	"encoding/xml"
)

// Complex types

// BaseRequest represents the BaseRequest complex type
type BaseRequest struct {
	SessionId *string `xml:"SessionId,omitempty"`
}

// DerivedRequest represents the DerivedRequest complex type
type DerivedRequest struct {
	SessionId *string `xml:"SessionId,omitempty"`
	Plate     *string `xml:"Plate,omitempty"`
}

// DerivedResult represents the DerivedResult complex type
type DerivedResult struct {
	Success *bool   `xml:"Success,omitempty"`
	Value   *string `xml:"Value,omitempty"`
}

// GrandDerivedRequest represents the GrandDerivedRequest complex type
type GrandDerivedRequest struct {
	SessionId *string `xml:"SessionId,omitempty"`
	Plate     *string `xml:"Plate,omitempty"`
	Extra     *string `xml:"Extra,omitempty"`
}

// ResultBase represents the ResultBase complex type
type ResultBase struct {
	Success *bool `xml:"Success,omitempty"`
}

// DoWorkWrapper represents the DoWork element
type DoWorkWrapper struct {
	XMLName xml.Name             `xml:"http://example.com/inherit DoWork"`
	Context *GrandDerivedRequest `xml:"Context,omitempty"`
}

// DoWorkResponseWrapper represents the DoWorkResponse element
type DoWorkResponseWrapper struct {
	XMLName      xml.Name       `xml:"http://example.com/inherit DoWorkResponse"`
	DoWorkResult *DerivedResult `xml:"DoWorkResult,omitempty"`
}
