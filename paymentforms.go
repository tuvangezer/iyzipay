package iyzipay

import "encoding/json"

type CheckoutFormInitializeRequest struct {
	Locale              string      `json:"locale,omitempty"`
	ConversationID      string      `json:"conversationId,omitempty"`
	Price               string      `json:"price"`
	BasketID            string      `json:"basketId,omitempty"`
	PaymentGroup        string      `json:"paymentGroup,omitempty"`
	PaymentSource       string      `json:"paymentSource,omitempty"`
	PosOrderId          string      `json:"posOrderId,omitempty"`
	Buyer               Buyer       `json:"buyer"`
	ShippingAddress     Address     `json:"shippingAddress"`
	BillingAddress      Address     `json:"billingAddress"`
	BasketItems         BasketItems `json:"basketItems"`
	CallbackURL         string      `json:"callbackUrl"`
	Currency            string      `json:"currency"`
	PaidPrice           string      `json:"paidPrice"`
	EnabledInstallments []int       `json:"enabledInstallments,omitempty"`
	ForceThreeDS        string      `json:"forceThreeDS,omitempty"`
	CardUserKey         string      `json:"cardUserKey,omitempty"`
}
type CheckoutFormInitializeResponse struct {
	Status              string `json:"status"`
	ErrorCode           string `json:"errorCode"`
	ErrorMessage        string `json:"errorMessage"`
	ErrorGroup          string `json:"errorGroup"`
	Locale              string `json:"locale"`
	SystemTime          int64  `json:"systemTime"`
	ConversationID      string `json:"conversationId"`
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkoutFormContent"`
	TokenExpireTime     int    `json:"tokenExpireTime"`
	PaymentPageURL      string `json:"paymentPageUrl"`
}
type CheckoutFormRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	Token          string `json:"token"`
}
type CheckoutFormResponse struct {
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
	CardToken                    string  `json:"cardToken"`
	CardUserKey                  string  `json:"cardUserKey"`
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

func (r *CheckoutFormInitializeRequest) prep() {
	r.Price = sanitizePrice(r.Price)
	r.PaidPrice = sanitizePrice(r.Price)
	for ind := range r.BasketItems {
		r.BasketItems[ind].Price = sanitizePrice(r.BasketItems[ind].Price)
		r.BasketItems[ind].SubMerchantPrice = sanitizePrice(r.BasketItems[ind].SubMerchantPrice)
	}
}

// CheckoutFormInitialize gets the HTML required to render the payment system
// https://dev.iyzipay.com/tr/odeme-formu/odeme-formu-baslatma
func (c *Client) CheckoutFormInitialize(request *CheckoutFormInitializeRequest) (response *CheckoutFormInitializeResponse, err error) {
	if request.prep(); !valid(*request) {
		return &CheckoutFormInitializeResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationID)
	p.append("price", request.Price)
	p.append("basketId", request.BasketID)
	p.append("paymentGroup", request.PaymentGroup)
	p.append("buyer", request.Buyer.getPKI())
	p.append("shippingAddress", request.ShippingAddress.getPKI())
	p.append("billingAddress", request.BillingAddress.getPKI())
	p.appendArray("basketItems", request.BasketItems.getPKIArray())
	p.append("callbackUrl", request.CallbackURL)
	p.append("paymentSource", request.PaymentSource)
	p.append("currency", request.Currency)
	p.append("posOrderId", request.PosOrderId)
	p.append("paidPrice", request.PaidPrice)
	p.append("forceThreeDS", request.ForceThreeDS)
	p.append("cardUserKey", request.CardUserKey)
	p.appendIntArray("enabledInstallments", request.EnabledInstallments)

	r := c.request("POST", "/payment/iyzipos/checkoutform/initialize/ecom", request, p)

	var ret CheckoutFormInitializeResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// CheckoutForm gets the details and status of the form with specified token
// https://dev.iyzipay.com/tr/odeme-formu/odeme-formu-sonucu
func (c *Client) CheckoutForm(request *CheckoutFormRequest) (response *CheckoutFormResponse, err error) {
	if !valid(*request) {
		return &CheckoutFormResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationID)
	p.append("token", request.Token)

	r := c.request("POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", request, p)

	var ret CheckoutFormResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}
