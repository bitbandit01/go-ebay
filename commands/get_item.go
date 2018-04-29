package commands

import (
	"encoding/xml"

	"github.com/bitbandit01/go-ebay/config"
)

func NewGetItemRequest(cfg *config.Config, itemID string) *GetItemRequest {
	return &GetItemRequest{
		ItemID: itemID,
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		Xmlns:       "urn:ebay:apis:eBLBaseComponents",
		DetailLevel: "ItemReturnDescription",
	}
}

type GetItemRequest struct {
	ItemID               string
	Xmlns                string               `xml:"xmlns,attr"`
	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`
	DetailLevel          string
}

func (c GetItemRequest) CallName() string {
	return "GetItem"
}

// GetRequestBody returns bytes
func (c GetItemRequest) GetRequestBody() []byte {
	body, _ := xml.Marshal(c)
	return body
}

func (c GetItemRequest) SetToken(token string) {
	c.RequesterCredentials.EBayAuthToken = token
}

type GetItemResponse struct {
	ebayResponse
	Item ItemResponse `xml:"Item"`
}

type ItemResponse struct {
	Title                       string                `xml:"Title"`
	Description                 string                `xml:"Description"`
	ItemID                      string                `xml:"ItemID"`
	SKU                         string                `xml:"SKU"`
	ViewItemURLForNaturalSearch string                `xml:"ViewItemURLForNaturalSearch"`
	PictureDetails              PictureDetails        `xml:"PictureDetails"`
	Quantity                    int                   `xml:"Quantity"`
	SellingStatus               SellingStatus         `xml:"SellingStatus"`
	ProductListingDetails       ProductListingDetails `xml:"ProductListingDetails"`
}

type PictureDetails struct {
}
