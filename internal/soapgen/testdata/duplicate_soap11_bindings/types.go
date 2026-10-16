package duplicate_soap11_bindings

import (
	"encoding/xml"
)

// EchoWrapper represents the Echo element
type EchoWrapper struct {
	XMLName      xml.Name `xml:"http://example.com/dup Echo"`
	InputMessage *string  `xml:"inputMessage,omitempty"`
}

// EchoResponseWrapper represents the EchoResponse element
type EchoResponseWrapper struct {
	XMLName    xml.Name `xml:"http://example.com/dup EchoResponse"`
	EchoResult *string  `xml:"EchoResult,omitempty"`
}
