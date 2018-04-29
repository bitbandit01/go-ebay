package commands

import (
	"encoding/xml"

	"github.com/bitbandit01/go-ebay/config"
)

func NewReviseItemRequest(cfg *config.Config, item ItemResponse) *ReviseItemRequest {
	return &ReviseItemRequest{
		Item: item,
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		Xmlns: "urn:ebay:apis:eBLBaseComponents",
	}
}

type ReviseItemRequest struct {
	Item                 ItemResponse
	Xmlns                string               `xml:"xmlns,attr"`
	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`
}

func (c ReviseItemRequest) CallName() string {
	return "ReviseItem"
}

// GetRequestBody returns bytes
func (c ReviseItemRequest) GetRequestBody() []byte {
	body, _ := xml.Marshal(c)
	return body
}

func (c ReviseItemRequest) SetToken(token string) {
	c.RequesterCredentials.EBayAuthToken = token
}
