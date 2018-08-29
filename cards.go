package iyzipay

import (
	"encoding/json"
)

type CreateCardRequest struct {
	Card           CardInformation `json:"card"`
	Email          string          `json:"email"`
	ExternalId     string          `json:"externalId,omitempty"`
	Locale         string          `json:"locale,omitempty"`
	ConversationId string          `json:"conversationId,omitempty"`
}
type CardInformation struct {
	CardAlias      string `json:"cardAlias"`
	CardHolderName string `json:"cardHolderName"`
	CardNumber     string `json:"cardNumber"`
	ExpireMonth    string `json:"expireMonth"`
	ExpireYear     string `json:"expireYear"`
}

type CreateCardResponse struct {
	Status          string `json:"status"`
	ErrorCode       string `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
	ErrorGroup      string `json:"errorGroup"`
	Locale          string `json:"locale"`
	SystemTime      int64  `json:"systemTime"`
	ConversationId  string `json:"conversationId"`
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	CardBankName    string `json:"cardBankName"`
	CardBankCode    int    `json:"cardBankCode"`
	Email           string `json:"email"`
	CardUserKey     string `json:"cardUserKey"`
	CardToken       string `json:"cardToken"`
	CardAlias       string `json:"cardAlias"`
}
type GetCardRequest struct {
	CardUserKey    string `json:"cardUserKey"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

type GetCardInformation struct {
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	CardBankName    string `json:"cardBankName"`
	CardBankCode    int    `json:"cardBankCode"`
	CardToken       string `json:"cardToken"`
	CardAlias       string `json:"cardAlias"`
}
type GetCardResponse struct {
	Status         string               `json:"status"`
	ErrorCode      string               `json:"errorCode"`
	ErrorMessage   string               `json:"errorMessage"`
	ErrorGroup     string               `json:"errorGroup"`
	Locale         string               `json:"locale"`
	SystemTime     int64                `json:"systemTime"`
	ConversationId string               `json:"conversationId"`
	CardUserKey    string               `json:"cardUserKey"`
	CardDetails    []GetCardInformation `json:"cardDetails"`
}
type AddCardRequest struct {
	CardUserKey    string          `json:"cardUserKey"`
	Card           CardInformation `json:"card"`
	Locale         string          `json:"locale,omitempty"`
	ConversationId string          `json:"conversationId,omitempty"`
}

type AddCardResponse CreateCardResponse

type DeleteCardRequest struct {
	CardToken      string `json:"cardToken"`
	CardUserKey    string `json:"cardUserKey"`
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
}

type DeleteCardResponse struct {
	Status         string `json:"status"`
	ErrorCode      string `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
	ErrorGroup     string `json:"errorGroup"`
	Locale         string `json:"locale"`
	SystemTime     int64  `json:"systemTime"`
	ConversationId string `json:"conversationId"`
}

func (c *CardInformation) getPKI() string {
	p := pkiBuilder{}
	p.append("cardAlias", c.CardAlias)
	p.append("cardNumber", c.CardNumber)
	p.append("expireYear", c.ExpireYear)
	p.append("expireMonth", c.ExpireMonth)
	p.append("cardHolderName", c.CardHolderName)
	return p.GetReqString()
}

// CreateCard creates a card in the iyzico card storage system.
// https://dev.iyzipay.com/tr/api/kart-saklama#request_create
func (c *Client) CreateCard(request *CreateCardRequest) (response *CreateCardResponse, err error) {
	if !valid(*request) {
		return &CreateCardResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("externalId", request.ExternalId)
	p.append("email", request.Email)
	p.append("card", request.Card.getPKI())

	r := c.request("POST", "/cardstorage/card", request, p)

	var ret CreateCardResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// GetCard gets the information regarding the card with specified attributes.
// https://dev.iyzipay.com/tr/api/kart-saklama#request_retrieve
func (c *Client) GetCard(request *GetCardRequest) (response *GetCardResponse, err error) {
	if !valid(*request) {
		return &GetCardResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("cardUserKey", request.CardUserKey)

	r := c.request("POST", "/cardstorage/cards", request, p)

	var ret GetCardResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// AddCard adds a card to a existing user.
// https://dev.iyzipay.com/tr/api/kart-saklama#request_add
func (c *Client) AddCard(request *AddCardRequest) (response *AddCardResponse, err error) {
	if !valid(*request) {
		return &AddCardResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("cardUserKey", request.CardUserKey)
	p.append("card", request.Card.getPKI())

	r := c.request("POST", "/cardstorage/card", request, p)

	var ret AddCardResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}

// DeleteCard deletes the specified card from the system.
// https://dev.iyzipay.com/tr/api/kart-saklama#request_delete
func (c *Client) DeleteCard(request *DeleteCardRequest) (response *DeleteCardResponse, err error) {
	if !valid(*request) {
		return &DeleteCardResponse{}, invalidFieldsErr
	}

	p := newResourcePKI(request.Locale, request.ConversationId)
	p.append("cardUserKey", request.CardUserKey)
	p.append("cardToken", request.CardToken)

	r := c.request("DELETE", "/cardstorage/card", request, p)

	var ret DeleteCardResponse

	json.NewDecoder(r.Body).Decode(&ret)

	return &ret, nil
}
