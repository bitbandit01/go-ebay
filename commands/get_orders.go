package commands

import (
	"encoding/xml"

	"github.com/bitbandit01/go-ebay/config"
)

// GetOrdersByCreationDate returns a pointer to GetOrdersRequest.  CreateTimeFrom and CreateTimeTo
//are ISO-8601 formatted date strings within which to return orders. Pagination is available by setting the
//Page variable (min:1).
func GetOrdersByCreationDate(cfg *config.Config, CreateTimeFrom string,
	CreateTimeTo string,
	OrderRole string,
	SortingOrder string,
	Page string) *GetOrdersRequest {

	return &GetOrdersRequest{
		CreateTimeFrom: CreateTimeFrom,
		CreateTimeTo:   CreateTimeTo,
		OrderRole:      OrderRole,
		SortingOrder:   SortingOrder,
		Pagination: PaginationRequest{
			EntriesPerPage: "100",
			PageNumber:     Page,
		},
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		Xmlns:   "urn:ebay:apis:eBLBaseComponents",
		Version: "1055",
	}
}

// GetRecentOrders returns a pointer to GetOrdersRequest.  Takes NumDays to look
//backwards in time for created or modified orders (min: 1, max: 30)
func GetRecentOrders(cfg *config.Config, NumDays int,
	OrderRole string,
	SortingOrder string,
	Page string) *GetOrdersRequest {

	return &GetOrdersRequest{
		NumberOfDays: NumDays,
		OrderRole:    OrderRole,
		SortingOrder: SortingOrder,
		Pagination: PaginationRequest{
			EntriesPerPage: "100",
			PageNumber:     Page,
		},
		RequesterCredentials: RequesterCredentials{
			EBayAuthToken: cfg.AuthToken,
		},
		DetailLevel: "ReturnAll",
		Xmlns:       "urn:ebay:apis:eBLBaseComponents",
		Version:     "1055",
	}
}

type GetOrdersRequest struct {
	CreateTimeFrom       string `xml:",omitempty"`
	CreateTimeTo         string `xml:",omitempty"`
	NumberOfDays         int
	OrderRole            string
	SortingOrder         string
	Pagination           PaginationRequest
	DetailLevel          string
	Xmlns                string `xml:"xmlns,attr"`
	Version              string
	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`
}

func (c GetOrdersRequest) CallName() string {
	return "GetOrders"
}

// GetRequestBody returns bytes
func (c GetOrdersRequest) GetRequestBody() []byte {
	body, _ := xml.Marshal(c)
	return body
}

func (c GetOrdersRequest) SetToken(token string) {
	c.RequesterCredentials.EBayAuthToken = token
}

type ErrorMessage struct {
	XmlName xml.Name `xml:"errorMessage"`
	Error   Error    `xml:"error"`
}

type Errors struct {
	XmlName      xml.Name `xml:"Errors"`
	ShortMessage string   `xml:"ShortMessage"`
	LongMessage  string   `xml:"LongMessage"`
	ErrorCode    string   `xml:"ErrorCode"`
}

type Error struct {
	ErrorId      string `xml:"errorId"`
	Domain       string `xml:"domain"`
	SeverityCode string `xml:"SeverityCode"`
	Severity     string `xml:"Severity"`
	Line         string `xml:"Line"`
	Column       string `xml:"Column"`
	Category     string `xml:"Category"`
	Message      string `xml:"ShortMessage"`
	SubDomain    string `xml:"subdomain"`
}

type EBay struct {
	ApplicationId string
}

type Item struct {
	Title  string
	ItemID string
	ListingDetails
	SellingStatus
}

type ListingDetails struct {
	EndTime   string
	StartTime string
}

/*<SellingStatus>
      <BidCount>0</BidCount>
      <BidIncrement currencyID="USD">0.25</BidIncrement>
      <ConvertedCurrentPrice currencyID="USD">1.0</ConvertedCurrentPrice>
      <CurrentPrice currencyID="USD">1.0</CurrentPrice>
      <LeadCount>0</LeadCount>
      <MinimumToBid currencyID="USD">1.0</MinimumToBid>
      <QuantitySold>0</QuantitySold>
      <ReserveMet>true</ReserveMet>
      <SecondChanceEligible>false</SecondChanceEligible>
      <ListingStatus>Active</ListingStatus>
	</SellingStatus>*/

type SellingStatus struct {
	CurrentPrice float64
	BidCount     string
	QuantitySold int
}

type GetOrdersRequestResponse struct {
	Xmlns         string `xml:"xmlns,attr"`
	Timestamp     string `xml:"Timestamp"`
	Ack           string `xml:"Ack"`
	Build         string `xml:"Build"`
	Errors        Errors `xml:"Errors"`
	HasMoreOrders bool   `xml:"HasMoreOrders"`

	CorrelationID         string `xml:"CorrelationID"`
	HardExpirationWarning string `xml:"HardExpirationWarning"`

	Paginations PaginationRequest `xml:"PaginationResult"`
	OrderArray  OrderArray

	OrdersPerPage            int `xml:"OrdersPerPage"`
	PageNumber               int `xml:"PageNumber"`
	ReturnedOrderCountActual int `xml:"ReturnedOrderCountActual"`
}

type Amount struct {
}
type BuyerPackageEnclosures struct {
}
type PaymentInstructionCode struct {
}
type TaxIdentifierAttributeCode struct {
}
type ValueTypeCode struct {
}
type OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type OrderArray struct {
	XMLName xml.Name `xml:"OrderArray"`
	Orders  []Order  `xml:"Order"`
}

type Order struct {
	XMLName                             xml.Name `xml:"Order"`
	OrderID                             string
	OrderStatus                         string
	BuyerCheckoutMessage                string
	BuyerUserID                         string
	CheckoutStatus                      CheckoutStatus
	CreatedTime                         string
	ModifiedTime                        string
	CreatingUserRole                    string
	EIASToken                           string
	ExtendedOrderID                     string
	ExternalTransactions                ExternalTransaction `xml:"ExternalTransaction"`
	IntegratedMerchantCreditCardEnabled string
	IsMultiLegShipping                  string
	LogisticsPlanType                   string
	MonetaryDetails                     MonetaryDetails         `xml:"MonetaryDetails"`
	MultiLegShippingDetails             MultiLegShippingDetails `xml:"MultiLegShippingDetails"`
	PaidTime                            string
	PaymentHoldDetails                  PaymentHoldDetails `xml:"PaymentHoldDetails"`
	PaymentHoldStatus                   string
	PaymentMethods                      string
	PickupDetailss                      PickupDetails `xml:"PickupDetails"`
	PickupMethodSelected
	RefundArrays      RefundArray `xml:"RefundArray"`
	SellerEIASToken   string
	SellerEmail       string
	SellerUserID      string
	ShippedTime       string
	ShippingDetails   ShippingDetails `xml:"ShippingDetails"`
	ShippingAddress   ShippingAddress `xml:"ShippingAddress"`
	Subtotal          string
	Total             string
	TransactionsArray TransactionArray
}
type Transaction struct {
	XMLName xml.Name `xml:"Transaction"`
	//Buyer             Buyer
	ShippingDetails   ShippingDetails `xml:"ShippingDetails"`
	CreatedDate       string
	Item              Item
	QuantityPurchased string
	//	Status Status
	TransactionID    string
	TransactionPrice string
	//ShippingServiceSelected ShippingServiceSelected
	FinalValueFee     string
	TransactionSiteID string
	Platform          string
	//Taxes Taxes
	ActualShippingCost  float32
	ActualHandlingCost  float32
	OrderLineItemID     string
	ExtendedOrderID     string
	eBayPlusTransaction bool
	Buyer               Buyer
}
type TransactionArray struct {
	XMLName      xml.Name      `xml:"TransactionArray"`
	Transactions []Transaction `xml:"Transaction"`
}

type PickupMethodSelected struct {
	XMLName               xml.Name `xml:"PickupMethodSelected"`
	MerchantPickupCode    string
	PickupFulfillmentTime string
	PickupLocationUUID    string
	PickupMethod          string
	PickupStatus          string
	PickupStoreID         string
}
type ShippingDetails struct {
	XMLName xml.Name `xml:"ShippingDetails"`
	CalculatedShippingRate
	CODCost         string
	InsuranceFee    string
	InsuranceOption string
	InsuranceWanted string

	InternationalShippingServiceOption
	//SellingManagerSalesRecordNumber
	ShipmentTrackingDetails
	ShippingServiceOptions ShippingServiceOptions `xml:"ShippingServiceOptions"`
	TaxTable
}
type ShippingAddress struct {
	XMLName           xml.Name `xml:"ShippingAddress"`
	AddressAttribute  string
	AddressID         string
	AddressOwner      string
	CityName          string
	Country           string
	CountryName       string
	ExternalAddressID string
	Name              string
	Phone             string
	PostalCode        string
	//ReferenceID       string
	StateOrProvince string
	Street1         string
	Street2         string
}
type TaxJurisdiction struct {
	JurisdictionID        string
	SalesTaxPercent       string
	ShippingIncludedInTax string
}
type TaxTable struct {
	TaxJurisdiction
}
type ShippingPackageInfo struct {
	ActualDeliveryTime       string
	ScheduledDeliveryTimeMax string
	ScheduledDeliveryTimeMin string
	ShippingTrackingEvent    string
	StoreID                  string
}
type ShippingServiceOptions struct {
	ExpeditedService              string
	ImportCharge                  string
	LogisticPlanType              string
	ShippingInsuranceCost         string
	ShippingService               string `xml:"ShippingService"`
	ShippingServiceAdditionalCost string
	ShippingServiceCost           string
	ShippingServicePriority       int
}
type ShipmentTrackingDetails struct {
	ShipmentTrackingNumber string
	ShippingCarrierUsed    string
}
type InternationalShippingServiceOption struct {
	//	ImportCharge  string
	//     ShippingInsuranceCost string
	//     ShippingService string
	//      ShippingServiceAdditionalCost  string
	//     ShippingServiceCost string
	//     ShippingServicePriority int
	//    ShipToLocation string
}

type CalculatedShippingRate struct {
	InternationalPackagingHandlingCosts string
	OriginatingPostalCode               string
	PackageDepth                        string
	PackageLength                       string
	PackageWidth                        string
	PackagingHandlingCosts              string
	ShippingIrregular                   string
	ShippingPackage                     string
	WeightMajor                         string
	WeightMinor                         string
}

type PickupOptions struct {
	PickupMethod   string
	PickupPriority string
}

type RefundArray struct{}
type PickupDetails struct {
	PickupOptions
}

type RequiredSellerActionArray struct {
	RequiredSellerAction string
}
type PaymentHoldDetails struct {
	ExpectedReleaseDate        string
	NumOfReqSellerActions      string
	PaymentHoldReason          string
	RequiredSellerActionArrays RequiredSellerActionArray
}
type ShipToAddress struct {
	XMLName           xml.Name `xml:"ShipToAddress"`
	AddressAttribute  string
	AddressID         string
	AddressOwner      string
	CityName          string
	Country           string
	CountryName       string
	ExternalAddressID string
	Name              string
	Phone             string
	PostalCode        string
	//ReferenceID       string
	StateOrProvince string
	Street1         string
	Street2         string
}
type SellerShipmentToLogisticsProvider struct {
	ShippingServiceDetails
	ShipToAddress
	ShippingTimeMax int
	ShippingTimeMin int
}
type ShippingServiceDetails struct {
	ShippingService   string
	TotalShippingCost string
}
type MultiLegShippingDetails struct {
	SellerShipmentToLogisticsProvider
}
type MonetaryDetails struct {
	Payments []Payment
	Refunds  []Refund
}

type Refund struct {
	Refund RefundDetails
}

type RefundDetails struct {
	FeeOrCreditAmount string
	//ReferenceID       string
	RefundAmount string
	RefundStatus string
	RefundTime   string
	RefundTo     string
	RefundType   string

	//RefundAmount string
	RefundFromSeller string
	//	RefundID string
	//	RefundStatus string
	//	RefundTime string
	TotalRefundToBuyer string
}

type SalesTax struct {
	SalesTaxAmount        string
	SalesTaxPercent       string
	SalesTaxState         string
	ShippingIncludedInTax string
}

type Payment struct {
	Payment PaymentDetails
}

type PaymentDetails struct {
	FeeOrCreditAmount string
	Payee             string
	Payer             string
	PaymentAmount     string
	ReferenceID       string
	PaymentStatus     string
	PaymentTime       string
}

type ExternalTransaction struct {
	ExternalTransactionID     string
	ExternalTransactionStatus string
	ExternalTransactionTime   string
	FeeOrCreditAmount         string
	PaymentOrRefundAmount     string
}
type BuyerPackageEnclosure struct {
	BuyerPackageEnclosureType string
}

type BuyerTaxIdentifier struct {
	Attribute string
	ID        string
	Type      string
}

type CancelDetail struct {
	CancelCompleteDate  string
	CancelIntiationDate string
	CancelIntiator      string
	CancelReason        string
	CancelReasonDetails string
}

type CheckoutStatus struct {
	eBayPaymentStatus                   string
	IntegratedMerchantCreditCardEnabled string
	LastModifiedTime                    string
	PaymentInstrument                   string
	PaymentMethod                       string
	Status                              string
}

type Buyer struct {
	Email         string
	UserFirstName string
	UserLastName  string
}
