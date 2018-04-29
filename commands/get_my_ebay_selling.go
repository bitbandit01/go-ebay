package commands

import (
	"encoding/xml"

	"github.com/bitbandit01/go-ebay/config"
)

func NewGetMyEbaySelling(cfg *config.Config, version string, pageNumber string, perPage string) *GetMyeBaySellingRequest {
	return &GetMyeBaySellingRequest{
		ActiveList: ActiveList{
			Include: true,
			Pagination: PaginationRequest{
				PageNumber:     pageNumber,
				EntriesPerPage: perPage,
			},
			Sort: "CurrentPriceDescending",
		},
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		Xmlns: "urn:ebay:apis:eBLBaseComponents",
	}
}

type GetMyeBaySellingRequest struct {
	ActiveList           ActiveList           `xml:"ActiveList"`
	Xmlns                string               `xml:"xmlns,attr"`
	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`
}
type ActiveList struct {
	Include    bool
	Pagination PaginationRequest
	Sort       string
}

func (c GetMyeBaySellingRequest) CallName() string {
	return "GetMyeBaySelling"
}

// GetRequestBody returns bytes
func (c GetMyeBaySellingRequest) GetRequestBody() []byte {
	body, _ := xml.Marshal(c)
	return body
}

func (c GetMyeBaySellingRequest) SetToken(token string) {
	c.RequesterCredentials.EBayAuthToken = token
}

type MyeBaySellingResponse struct {
	Ack                             string
	Build                           string
	CorrelationID                   string
	HardExpirationWarning           string
	Timestamp                       string
	Version                         string
	Summary                         SummaryStruct
	SellingSummary                  SummaryStruct
	ActiveList                      ItemArrayStruct         `xml:"ActiveList>ItemArray"`
	DeletedFromSoldList             []OrderTransactionArray `xml:"DeletedFromSoldList>OrderTransactionArray"`
	DeletedFromUnsoldList           []ItemArrayStruct       `xml:"DeletedFromUnsoldList>ItemArray"`
	ScheduledList                   []ItemArrayStruct       `xml:"ScheduledList>ItemArray"`
	SoldList                        []OrderTransactionArray `xml:"SoldList>OrderTransactionArray"`
	UnsoldList                      []ItemArrayStruct       `xml:"UnsoldList>ItemArray"`
	ActiveListPagination            PaginationStruct        `xml:"ActiveList>PaginationResult"`
	DeletedFromSoldListPagination   PaginationStruct        `xml:"DeletedFromSoldList>PaginationResult"`
	DeletedFromUnsoldListPagination PaginationStruct        `xml:"DeletedFromUnsoldList>PaginationResult"`
	ScheduledListPagination         PaginationStruct        `xml:"ScheduledList>PaginationResult"`
	SoldListPagination              PaginationStruct        `xml:"SoldList>PaginationResult"`
	UnsoldListPagination            PaginationStruct        `xml:"UnsoldList>PaginationResult"`
	Errors                          []ErrorMessage
}
type PaginationStruct struct {
	TotalNumberOfEntries string
	TotalNumberOfPages   string
}
type SummaryStruct struct {
	ActiveAuctionCount       string
	AmountLimitRemaining     string
	AuctionBidCount          string
	AuctionSellingCount      string
	ClassifiedAdCount        string
	ClassifiedAdOfferCount   string
	QuantityLimitRemaining   string
	SoldDurationInDays       string
	TotalAuctionSellingValue string
	TotalLeadCount           string
	TotalListingsWithLeads   string
	TotalSoldCount           string
	TotalSoldValue           string
}
type BuyerStruct struct {
	Email           string
	StaticAlias     string
	UserID          string
	ShippingAddress struct {
		PostalCode string
	}
}
type TransactionStruct struct {
	ConvertedTransactionPrice string
	CreatedDate               string
	IsMultiLegShipping        string
	LineItemID                string
	PaidTime                  string
	PaisaPayID                string
	Platform                  string
	QuantityPurchased         string
	SellerPaidStatus          string
	ShippedTime               string
	ID                        string `xml:"TransactionID"`
	Price                     string
	Buyer                     BuyerStruct `xml:"Buyer>BuyerInfo"`
	OrderLineItemID           string

	FeedbackLeft string `xml:"FeedbackLeft>CommentType"`

	FeedbackReceived string `xml:"FeedbackReceived>CommentType"`

	Item ItemResultStruct `xml:"Item"`

	Status struct {
		PaymentHold string
	}
}
type ItemResultStruct struct {
	BestOfferDetails struct {
		BestOfferCount string
	}
	BuyItNowPrice             string
	ClassifiedAdPayPerLeadFee string
	eBayNotes                 string
	HideFromSearch            string
	ItemID                    string
	ListingDuration           string
	ListingType               string
	PrivateNotes              string
	Quantity                  string
	QuantityAvailable         string
	QuestionCount             string
	ReasonHideFromSearch      string
	ReservePrice              string
	SKU                       string
	StartPrice                string
	TimeLeft                  string
	Title                     string
	WatchCount                string
}

type OrderStruct struct {
	ID           string `xml:"OrderID"`
	Subtotal     string
	Transactions []TransactionStruct `xml:"TransactionArray>Transaction"`
}
type OrderTransactionArray struct {
	Order        OrderStruct         `xml:"OrderTransaction>Order"`
	Transactions []TransactionStruct `xml:"Transaction"`
	Pagination   PaginationStruct    `xml:"PaginationResult"`
}

type ItemArrayStruct struct {
	Items []ItemResultStruct `xml:"Item"`
}
