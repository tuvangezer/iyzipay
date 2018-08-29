package iyzipay

import (
	"encoding/json"
)

type PaymentCard struct {
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireYear     string `json:"expireYear"`
	ExpireMonth    string `json:"expireMonth"`
	Cvc            string `json:"cvc"`
	RegisterCard   string `json:"registerCard,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
	CardToken      string `json:"cardToken,omitempty"`
}
type Buyer struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	IdentityNumber      string `json:"identityNumber"`
	Email               string `json:"email"`
	GsmNumber           string `json:"gsmNumber,omitempty"`
	RegistrationDate    string `json:"registrationDate,omitempty"`
	LastLoginDate       string `json:"lastLoginDate,omitempty"`
	RegistrationAddress string `json:"registrationAddress"`
	City                string `json:"city"`
	Country             string `json:"country"`
	ZipCode             string `json:"zipCode,omitempty"`
	IP                  string `json:"ip"`
}
type Address struct {
	Address     string `json:"address"`
	ZipCode     string `json:"zipCode,omitempty"`
	ContactName string `json:"contactName"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
type BasketItem struct {
	ID               string `json:"id"`
	Price            string `json:"price"`
	Name             string `json:"name"`
	Category1        string `json:"category1"`
	Category2        string `json:"category2"`
	ItemType         string `json:"itemType"`
	SubMerchantKey   string `json:"subMerchantKey,omitempty"`
	SubMerchantPrice string `json:"subMerchantPrice,omitempty"`
}
type BasketItems []BasketItem

func (c *PaymentCard) getPKI() string {
	p := pkiBuilder{}
	p.append("cardHolderName", c.CardHolderName)
	p.append("cardNumber", c.CardNumber)
	p.append("expireYear", c.ExpireYear)
	p.append("expireMonth", c.ExpireMonth)
	p.append("cvc", c.Cvc)
	p.append("registerCard", c.RegisterCard)
	if c.RegisterCard == "1" {
		p.append("cardAlias", c.CardAlias)
		p.append("cardToken", c.CardToken)
		p.append("cardUserKey", c.CardUserKey)
	}
	return p.GetReqString()
}
func (a *Address) getPKI() string {
	p := pkiBuilder{}
	p.append("address", a.Address)
	p.append("zipCode", a.ZipCode)
	p.append("contactName", a.ContactName)
	p.append("city", a.City)
	p.append("country", a.Country)

	return p.GetReqString()
}
func (b *Buyer) getPKI() string {
	p := pkiBuilder{}
	p.append("id", b.ID)
	p.append("name", b.Name)
	p.append("surname", b.Surname)
	p.append("identityNumber", b.IdentityNumber)
	p.append("email", b.Email)
	p.append("gsmNumber", b.GsmNumber)
	p.append("registrationDate", b.RegistrationDate)
	p.append("lastLoginDate", b.LastLoginDate)
	p.append("registrationAddress", b.RegistrationAddress)
	p.append("city", b.City)
	p.append("country", b.Country)
	p.append("zipCode", b.ZipCode)
	p.append("ip", b.IP)
	return p.GetReqString()
}
func (b *BasketItems) getPKIArray() []string {
	var arr []string
	for _, item := range *b {
		p := pkiBuilder{}
		p.append("id", item.ID)
		p.append("price", sanitizePrice(item.Price))
		p.append("name", item.Name)
		p.append("category1", item.Category1)
		p.append("category2", item.Category2)
		p.append("itemType", item.ItemType)
		p.append("subMerchantKey", item.SubMerchantKey)
		p.append("subMerchantPrice", sanitizePrice(item.SubMerchantPrice))
		arr = append(arr, p.GetReqString())
	}
	return arr
}

type CreatePaymentRequest struct {
	Locale          string      `json:"locale,omitempty"`
	ConversationId  string      `json:"conversationId,omitempty"`
	Price           string      `json:"price"`
	PaidPrice       string      `json:"paidPrice"`
	Installment     string      `json:"installment"`
	PaymentChannel  string      `json:"paymentChannel,omitempty"`
	BasketID        string      `json:"basketId,omitempty"`
	PaymentGroup    string      `json:"paymentGroup,omitempty"`
	PaymentCard     PaymentCard `json:"paymentCard"`
	Buyer           Buyer       `json:"buyer"`
	ShippingAddress Address     `json:"shippingAddress"`
	BillingAddress  Address     `json:"billingAddress"`
	BasketItems     BasketItems `json:"basketItems"`
	Currency        string      `json:"currency"`
	PosOrderId      string      `json:"posOrderId,omitempty"`
	ConnectorName   string      `json:"connectorName,omitempty"`
	CallbackURL     string      `json:"callbackUrl,omitempty"`
	PaymentSource   string      `json:"paymentSource,omitempty"`
}
type CreatePaymentResponse struct {
	Status                       string  `json:"status"`
	ErrorCode                    string  `json:"errorCode"`
	ErrorMessage                 string  `json:"errorMessage"`
	ErrorGroup                   string  `json:"errorGroup"`
	Locale                       string  `json:"locale"`
	SystemTime                   int64   `json:"systemTime"`
	ConversationId               string  `json:"conversationId"`
	Price                        int     `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        int     `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  int     `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              int     `json:"subMerchantPrice"`
		SubMerchantPayoutRate         int     `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            int     `json:"iyziConversionRate"`
			IyziConversionRateAmount      int     `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
}
type CancelPaymentRequest struct {
	IP             string `json:"ip"`
	PaymentId      string `json:"paymentId"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Description    string `json:"description,omitempty"`
}
type CancelPaymentResponse struct {
	Status         string `json:"status"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationId string `json:"conversationId"`
	Price          string `json:"price"`
	Currency       string `json:"currency"`
	PaymentId      string `json:"paymentId"`
}
type GetPaymentRequest struct {
	IP                    string `json:"ip"`
	PaymentId             string `json:"paymentId"`
	Locale                string `json:"locale,omitempty"`
	ConversationId        string `json:"conversationId,omitempty"`
	PaymentConversationId string `json:"paymentConversationId,omitempty"`
}
type GetPaymentResponse struct {
	Status                       string  `json:"status"`
	ErrorCode                    string  `json:"errorCode"`
	ErrorMessage                 string  `json:"errorMessage"`
	ErrorGroup                   string  `json:"errorGroup"`
	Locale                       string  `json:"locale"`
	SystemTime                   int64   `json:"systemTime"`
	ConversationID               string  `json:"conversationId"`
	Price                        int     `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        int     `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  int     `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              int     `json:"subMerchantPrice"`
		SubMerchantPayoutRate         int     `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            int     `json:"iyziConversionRate"`
			IyziConversionRateAmount      int     `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
	PaymentStatus string `json:"paymentStatus"`
}
type RefundPaymentRequest struct {
	Locale               string `json:"locale,omitempty"`
	IP                   string `json:"ip"`
	Price                string `json:"price"`
	PaymentTransactionId string `json:"paymentTransactionId"`
	ConversationId       string `json:"conversationId,omitempty"`
	Currency             string `json:"currency,omitempty"`
	Reason               string `json:"reason,omitempty"`
	Description          string `json:"description,omitempty"`
}
type RefundPaymentResponse struct {
	Status               string `json:"status"`
	ErrorCode            string `json:"errorCode"`
	ErrorMessage         string `json:"errorMessage"`
	ErrorGroup           string `json:"errorGroup"`
	Locale               string `json:"locale"`
	SystemTime           int64  `json:"systemTime"`
	ConversationID       string `json:"conversationId"`
	PaymentID            string `json:"paymentId"`
	PaymentTransactionID string `json:"paymentTransactionId"`
	Price                int    `json:"price"`
	Currency             string `json:"currency"`
}

func (r *RefundPaymentRequest) prep() {
	r.Price = sanitizePrice(r.Price)
}
func (r *CreatePaymentRequest) prep() {

	r.Price = sanitizePrice(r.Price)
	r.PaidPrice = sanitizePrice(r.PaidPrice)
	for _, v := range r.BasketItems {
		v.Price = sanitizePrice(v.Price)

	}
}

// CreatePayment is used for simple payment processes.
// https://dev.iyzipay.com/tr/api/odeme#request
func (c *Client) CreatePayment(request *CreatePaymentRequest) (response *CreatePaymentResponse, err error) {
	if request.prep(); !valid(*request) {
		return &CreatePaymentResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("price", request.Price)
	p.append("paidPrice", request.PaidPrice)
	p.append("installment", request.Installment)
	p.append("paymentChannel", request.PaymentChannel)
	p.append("basketId", request.BasketID)
	p.append("paymentGroup", request.PaymentGroup)
	p.append("paymentCard", request.PaymentCard.getPKI())
	p.append("buyer", request.Buyer.getPKI())
	p.append("shippingAddress", request.ShippingAddress.getPKI())
	p.append("billingAddress", request.BillingAddress.getPKI())
	p.appendArray("basketItems", request.BasketItems.getPKIArray())
	p.append("currency", request.Currency)
	p.append("posOrderId", request.PosOrderId)
	p.append("connectorName", request.ConnectorName)
	p.append("callbackUrl", request.CallbackURL)

	r := c.request("POST", "/payment/auth", request, p)

	var ret CreatePaymentResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// CancelPayment cancels the specified payment.
// https://dev.iyzipay.com/tr/api/iptal#request
func (c *Client) CancelPayment(request *CancelPaymentRequest) (response *CancelPaymentResponse, err error) {
	if !valid(*request) {
		return &CancelPaymentResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("paymentId", request.PaymentId)
	p.append("ip", request.IP)
	p.append("reason", request.Reason)
	p.append("description", request.Description)

	r := c.request("POST", "/payment/cancel", request, p)

	var ret CancelPaymentResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// GetPayment gets information of the specified payment.
// https://dev.iyzipay.com/tr/api/odeme-sorgulama
func (c *Client) GetPayment(request *GetPaymentRequest) (response *GetPaymentResponse, err error) {
	if !valid(*request) {
		return &GetPaymentResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("paymentId", request.PaymentId)
	p.append("paymentConversationId", request.PaymentConversationId)

	r := c.request("POST", "/payment/detail", request, p)

	var ret GetPaymentResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// RefundPayment is used to refund the specified payment.
// https://dev.iyzipay.com/tr/api/iade
func (c *Client) RefundPayment(request *RefundPaymentRequest) (response *RefundPaymentResponse, err error) {
	if request.prep(); !valid(*request) {
		return &RefundPaymentResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("paymentTransactionId", request.PaymentTransactionId)
	p.append("price", request.Price)
	p.append("ip", request.IP)
	p.append("currency", request.Currency)
	p.append("reason", request.Reason)
	p.append("description", request.Description)

	r := c.request("POST", "/payment/refund", request, p)

	var ret RefundPaymentResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

/*
=========== 3D Secure ===========
*/

type ThreeDSInitRequest struct {
	Locale          string      `json:"locale,omitempty"`
	ConversationId  string      `json:"conversationId,omitempty"`
	Price           string      `json:"price"`
	PaidPrice       string      `json:"paidPrice"`
	Installment     string      `json:"installment"`
	PaymentChannel  string      `json:"paymentChannel,omitempty"`
	BasketID        string      `json:"basketId,omitempty"`
	PaymentGroup    string      `json:"paymentGroup,omitempty"`
	PaymentCard     PaymentCard `json:"paymentCard"`
	Buyer           Buyer       `json:"buyer"`
	ShippingAddress Address     `json:"shippingAddress"`
	BillingAddress  Address     `json:"billingAddress"`
	BasketItems     BasketItems `json:"basketItems"`
	Currency        string      `json:"currency"`
	CallbackURL     string      `json:"callbackUrl"`
	PaymentSource   string      `json:"paymentSource,omitempty"`
}
type ThreeDSInitResponse struct {
	Status         string `json:"status"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationId string `json:"conversationId"`
	HtmlContent    string `json:"threeDSHtmlContent"`
}
type ThreeDSPayRequest struct {
	Locale           string `json:"locale,omitempty"`
	ConversationId   string `json:"conversationId,omitempty"`
	PaymentId        string `json:"paymentId"`
	ConversationData string `json:"conversationData,omitempty"`
}
type ThreeDSPayResponse struct {
	Status                       string  `json:"status"`
	ErrorCode                    string  `json:"errorCode"`
	ErrorMessage                 string  `json:"errorMessage"`
	ErrorGroup                   string  `json:"errorGroup"`
	Locale                       string  `json:"locale"`
	SystemTime                   int64   `json:"systemTime"`
	ConversationID               string  `json:"conversationId"`
	Price                        int     `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        int     `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  int     `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              int     `json:"subMerchantPrice"`
		SubMerchantPayoutRate         int     `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            int     `json:"iyziConversionRate"`
			IyziConversionRateAmount      int     `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
}

func (r *ThreeDSInitRequest) prep() {
	r.Price = sanitizePrice(r.Price)
	r.PaidPrice = sanitizePrice(r.PaidPrice)
	for _, v := range r.BasketItems {
		v.Price = sanitizePrice(v.Price)

	}
}

// ThreeDSInit is used to initialize a 3DS payment process
// https://dev.iyzipay.com/tr/api/3d-ile-odeme#request_init
func (c *Client) ThreeDSInit(request *ThreeDSInitRequest) (response *ThreeDSInitResponse, err error) {
	if request.prep(); !valid(*request) {
		return &ThreeDSInitResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("price", request.Price)
	p.append("paidPrice", request.PaidPrice)
	p.append("installment", request.Installment)
	p.append("paymentChannel", request.PaymentChannel)
	p.append("basketId", request.BasketID)
	p.append("paymentGroup", request.PaymentGroup)
	p.append("paymentCard", request.PaymentCard.getPKI())
	p.append("buyer", request.Buyer.getPKI())
	p.append("shippingAddress", request.ShippingAddress.getPKI())
	p.append("billingAddress", request.BillingAddress.getPKI())
	p.appendArray("basketItems", request.BasketItems.getPKIArray())
	p.append("paymentSource", request.PaymentSource)
	p.append("currency", request.Currency)
	p.append("callbackUrl", request.CallbackURL)

	r := c.request("POST", "/payment/3dsecure/initialize", request, p)

	var ret ThreeDSInitResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// ThreeDSPay is used to finalize 3DS payments after getting the callback from bank.
// https://dev.iyzipay.com/tr/api/3d-ile-odeme#request_auth
// WARNING: Not covered by tests as it requires a paymentId
func (c *Client) ThreeDSPay(request *ThreeDSPayRequest) (response *ThreeDSPayResponse, err error) {
	if !valid(*request) {
		return &ThreeDSPayResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("paymentId", request.PaymentId)
	p.append("conversationData", request.ConversationData)

	r := c.request("POST", "/payment/3dsecure/auth", request, p)

	var ret ThreeDSPayResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

/*
========== BKM Express ==========
*/
type BkmInitRequest struct {
	Locale              string      `json:"locale,omitempty"`
	ConversationId      string      `json:"conversationId,omitempty"`
	Price               string      `json:"price"`
	Installment         string      `json:"installment,omitempty"`
	PaymentChannel      string      `json:"paymentChannel,omitempty"`
	BasketID            string      `json:"basketId,omitempty"`
	PaymentGroup        string      `json:"paymentGroup,omitempty"`
	Buyer               Buyer       `json:"buyer"`
	ShippingAddress     Address     `json:"shippingAddress"`
	BillingAddress      Address     `json:"billingAddress"`
	BasketItems         BasketItems `json:"basketItems"`
	Currency            string      `json:"currency,omitempty"`
	CallbackURL         string      `json:"callbackUrl"`
	PaymentSource       string      `json:"paymentSource,omitempty"`
	EnabledInstallments string      `json:"enabledInstallments,omitempty"`
}
type BKMInitResponse struct {
	Status          string `json:"status"`
	ErrorCode       string `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
	ErrorGroup      string `json:"errorGroup"`
	Locale          string `json:"locale"`
	SystemTime      int64  `json:"systemTime"`
	ConversationID  string `json:"conversationId"`
	HTMLContent     string `json:"htmlContent"`
	RedirectURL     string `json:"redirectUrl"`
	Token           string `json:"token"`
	TokenExpireTime string `json:"tokenExpireTime"`
}
type BKMGetRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Token          string `json:"token"`
}
type BKMGetResponse struct {
	Status                       string  `json:"status"`
	ErrorCode                    string  `json:"errorCode"`
	ErrorMessage                 string  `json:"errorMessage"`
	ErrorGroup                   string  `json:"errorGroup"`
	Locale                       string  `json:"locale"`
	ConversationId               string  `json:"conversationId"`
	SystemTime                   int64   `json:"systemTime"`
	Price                        int     `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       int     `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        int     `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  int     `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              int     `json:"subMerchantPrice"`
		SubMerchantPayoutRate         int     `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant int     `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       int     `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            int     `json:"iyziConversionRate"`
			IyziConversionRateAmount      int     `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
	Token         string `json:"token"`
	CallbackURL   string `json:"callbackUrl"`
	PaymentStatus string `json:"paymentStatus"`
}

func (r *BkmInitRequest) prep() {
	r.Price = sanitizePrice(r.Price)
	for _, v := range r.BasketItems {
		v.Price = sanitizePrice(v.Price)
	}
}

// BKMInit is used to initialize a BKM payment process.
// https://dev.iyzipay.com/tr/api/bkm-express-ile-odeme#request_init
func (c *Client) BKMInit(request *BkmInitRequest) (response *BKMInitResponse, err error) {
	if !valid(*request) {
		return &BKMInitResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("price", request.Price)
	p.append("basketId", request.BasketID)
	p.append("paymentGroup", request.PaymentGroup)
	p.append("buyer", request.Buyer.getPKI())
	p.append("shippingAddress", request.ShippingAddress.getPKI())
	p.append("billingAddress", request.BillingAddress.getPKI())
	p.appendArray("basketItems", request.BasketItems.getPKIArray())
	p.append("callbackUrl", request.CallbackURL)
	p.append("paymentSource", request.PaymentSource)
	p.append("enabledInstallments", request.EnabledInstallments)

	r := c.request("POST", "/payment/bkm/initialize", request, p)

	var ret BKMInitResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// BKMGet Gets the status of BKM payment with the token specified.
// https://dev.iyzipay.com/tr/api/bkm-express-ile-odeme#request_auth
func (c *Client) BKMGet(request *BKMGetRequest) (response *BKMGetResponse, err error) {
	if !valid(*request) {
		return &BKMGetResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("token", request.Token)

	r := c.request("POST", "/payment/bkm/auth/detail", request, p)

	var ret BKMGetResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}
