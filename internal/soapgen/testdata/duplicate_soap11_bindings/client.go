package duplicate_soap11_bindings

import (
	"context"
	"encoding/xml"
	"fmt"
	soap "github.com/way-platform/soap-go"
)

// ClientOption configures a Client.
type ClientOption = soap.ClientOption

// Client is a SOAP client for this service.
type Client struct {
	*soap.Client
}

// NewClient creates a new SOAP client.
func NewClient(opts ...ClientOption) (*Client, error) {
	soapOpts := append([]soap.ClientOption{
		soap.WithEndpoint("http://example.com/dup"),
	}, opts...)
	soapClient, err := soap.NewClient(soapOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &Client{
		Client: soapClient,
	}, nil
}

// Echo executes the Echo SOAP operation.
func (c *Client) Echo(ctx context.Context, req *EchoWrapper, opts ...ClientOption) (*EchoResponseWrapper, error) {
	reqEnvelope, err := soap.NewEnvelope(soap.WithBody(req))
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP envelope: %w", err)
	}
	respEnvelope, err := c.Call(ctx, "http://example.com/dup/Echo", reqEnvelope, opts...)
	if err != nil {
		return nil, fmt.Errorf("SOAP call failed: %w", err)
	}
	var result EchoResponseWrapper
	if err := xml.Unmarshal(respEnvelope.Body.Content, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	return &result, nil
}
