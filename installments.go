package iyzipay

import (
	"encoding/json"
)

type InstallmentRequest struct {
	BinNumber      string `json:"binNumber,omitempty"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Price          string `json:"price"`
}
type InstallmentResponse struct {
	Status             string `json:"status"`
	Locale             string `json:"locale"`
	SystemTime         int64  `json:"systemTime"`
	ConversationID     string `json:"conversationId"`
	InstallmentDetails []struct {
		BinNumber         string `json:"binNumber"`
		Price             int    `json:"price"`
		CardType          string `json:"cardType"`
		CardAssociation   string `json:"cardAssociation"`
		CardFamilyName    string `json:"cardFamilyName"`
		Force3Ds          int    `json:"force3ds"`
		BankCode          int    `json:"bankCode"`
		BankName          string `json:"bankName"`
		ForceCvc          int    `json:"forceCvc"`
		InstallmentPrices []struct {
			InstallmentPrice  int `json:"installmentPrice"`
			TotalPrice        int `json:"totalPrice"`
			InstallmentNumber int `json:"installmentNumber"`
		} `json:"installmentPrices"`
	} `json:"installmentDetails"`
}

func (r *InstallmentRequest) prep() {
	r.Price = sanitizePrice(r.Price)
}


// GetInstallmentInformation is used to query possible installment number for specified information.
// https://dev.iyzipay.com/tr/api/taksit-sorgulama
func (c *Client) GetInstallmentInformation(request *InstallmentRequest) (response *InstallmentResponse, err error) {
	if request.prep();!valid(*request) {
		return &InstallmentResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("binNumber", request.BinNumber)
	p.append("price", request.Price)

	r := c.request("POST", "/payment/iyzipos/installment", request, p)

	var ret InstallmentResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}
