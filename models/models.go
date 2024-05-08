package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
)

type MPesaConfig struct {
	Username                  string `json:"username,omitempty"`
	Password                  string `json:"password,omitempty"`
	BusinessShortCode         string `json:"business_short_code,omitempty"`
	Passkey                   string `json:"passkey,omitempty"`
	OnlineCheckoutCallbackURL string `json:"online_checkout_callback_url,omitempty"`
	OnlineCheckoutInitiateURL string `json:"online_checkout_initiate_url,omitempty"`
	TransactionType           string `json:"transaction_type,omitempty"`
}

func (m MPesaConfig) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Username, validation.Required.Error("Username is required")),
		validation.Field(&m.Password, validation.Required.Error("Password is required")),
		validation.Field(&m.BusinessShortCode, validation.Required.Error("Business short code is required")),
		validation.Field(&m.Passkey, validation.Required.Error("Pass key is required")),
		validation.Field(&m.OnlineCheckoutCallbackURL, validation.Required.Error("online checkout callback url is required")),
		validation.Field(&m.OnlineCheckoutInitiateURL, validation.Required.Error("online checkout initiate url is required")),
		validation.Field(&m.TransactionType, validation.Required.Error("transaction type is required")),
	)
}

type MpesaPaymentInitiate struct {
	CommandID     string `json:"CommandID"`     //nolint:tagliatelle //required by the payment provider
	Amount        string `json:"Amount"`        //nolint:tagliatelle
	Msisdn        string `json:"Msisdn"`        //nolint:tagliatelle
	BillRefNumber string `json:"BillRefNumber"` //nolint:tagliatelle
	ShortCode     string `json:"ShortCode"`     //nolint:tagliatelle
}

type MpesaPaymentInitiateResponse struct {
	ConversationID           string `json:"ConversationID"`           //nolint:tagliatelle
	OriginatorConversationID string `json:"OriginatorConversationID"` //nolint:tagliatelle
	ResponseDescription      string `json:"ResponseDescription"`      //nolint:tagliatelle
}

type ValidationTransaction struct {
	RequestType       string `json:"RequestType"`       //nolint:tagliatelle
	TransactionType   string `json:"TransactionType"`   //nolint:tagliatelle
	TransID           string `json:"TransID"`           //nolint:tagliatelle
	TransTime         string `json:"TransTime"`         //nolint:tagliatelle
	TransAmount       string `json:"TransAmount"`       //nolint:tagliatelle
	BusinessShortCode string `json:"BusinessShortCode"` //nolint:tagliatelle
	BillRefNumber     string `json:"BillRefNumber"`     //nolint:tagliatelle
	InvoiceNumber     string `json:"InvoiceNumber"`     //nolint:tagliatelle
	OrgAccountBalance string `json:"OrgAccountBalance"` //nolint:tagliatelle
	ThirdPartyTransID string `json:"ThirdPartyTransID"` //nolint:tagliatelle
	MSISDN            string `json:"MSISDN"`            //nolint:tagliatelle
	FirstName         string `json:"FirstName"`         //nolint:tagliatelle
	MiddleName        string `json:"MiddleName"`        //nolint:tagliatelle
	LastName          string `json:"LastName"`          //nolint:tagliatelle
}

type ValidationCallback struct {
	ResultCode string `json:"ResultCode"` //nolint:tagliatelle
	ResultDesc string `json:"ResultDesc"` //nolint:tagliatelle
}

func (v *ValidationCallback) Error() string {
	return v.ResultCode
}

type MpesaOcInitiate struct {
	BusinessShortCode string          `json:"BusinessShortCode"` //nolint:tagliatelle
	Password          string          `json:"Password"`          //nolint:tagliatelle
	Timestamp         string          `json:"Timestamp"`         //nolint:tagliatelle
	TransactionType   string          `json:"TransactionType"`   //nolint:tagliatelle
	Amount            decimal.Decimal `json:"Amount"`            //nolint:tagliatelle
	PartyA            string          `json:"PartyA"`            //nolint:tagliatelle
	PartyB            string          `json:"PartyB"`            //nolint:tagliatelle
	PhoneNumber       string          `json:"PhoneNumber"`       //nolint:tagliatelle
	CallBackURL       string          `json:"CallBackURL"`       //nolint:tagliatelle
	AccountReference  string          `json:"AccountReference"`  //nolint:tagliatelle
	TransactionDesc   string          `json:"TransactionDesc"`   //nolint:tagliatelle
}

func (c MpesaOcInitiate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Amount, validation.Required.Error("amount is required")),
		validation.Field(&c.PhoneNumber, validation.Required.Error("phone-number is required")),
		validation.Field(&c.PartyA, validation.Required.Error("Pparty-A is required")),
		validation.Field(&c.AccountReference, validation.Required.Error("account-reference is required")),
	)
}

type MpesaOcInitiateResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`   //nolint:tagliatelle
	CheckoutRequestID   string `json:"CheckoutRequestID"`   //nolint:tagliatelle
	ResponseCode        string `json:"ResponseCode"`        //nolint:tagliatelle
	ResponseDescription string `json:"ResponseDescription"` //nolint:tagliatelle
	CustomerMessage     string `json:"CustomerMessage"`     //nolint:tagliatelle
}
type MpesaOcInitiateErrorResponse struct {
	RequestID    string `json:"requestId"`    //nolint:tagliatelle
	ErrorCode    string `json:"errorCode"`    //nolint:tagliatelle
	ErrorMessage string `json:"errorMessage"` //nolint:tagliatelle
}

type CallbackMetadataItem struct {
	Name  string      `json:"Name"`  //nolint:tagliatelle
	Value interface{} `json:"Value"` //nolint:tagliatelle
}

type CallbackMetadata struct {
	Item []CallbackMetadataItem `json:"Item"` //nolint:tagliatelle
}

type StkCallback struct {
	MerchantRequestID string           `json:"MerchantRequestID"` //nolint:tagliatelle
	CheckoutRequestID string           `json:"CheckoutRequestID"` //nolint:tagliatelle
	ResultCode        int              `json:"ResultCode"`        //nolint:tagliatelle
	ResultDesc        string           `json:"ResultDesc"`        //nolint:tagliatelle
	CallbackMetadata  CallbackMetadata `json:"CallbackMetadata"`  //nolint:tagliatelle
}

type Body struct {
	StkCallback StkCallback `json:"stkCallback"` //nolint:tagliatelle
}
type EnvelopeBody struct {
	Body Body `json:"Body"` //nolint:tagliatelle
}

type Envelope struct {
	Envelope EnvelopeBody `json:"Envelope"` //nolint:tagliatelle
}
