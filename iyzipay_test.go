package iyzipay

import (
	"testing"
	"log"
	"strings"
)


var (
	sandboxKey    = "--"
	sandboxSecret = "--"
	sandboxURL    = "https://sandbox-api.iyzipay.com"

	client = Create(sandboxKey, sandboxSecret, sandboxURL)

	dummyAddress = Address{
		Address:     "Nidakule Goztepe, Merdivenkoy Mah. Bora Sok. No:1",
		ZipCode:     "34732",
		ContactName: "Jane Doe",
		City:        "Istanbul",
		Country:     "Turkey",
	}
	dummyBuyer = Buyer{
		ID:                  "BY789",
		Name:                "John",
		Surname:             "Doe",
		GsmNumber:           "+905350000000",
		Email:               "email@email.com",
		IdentityNumber:      "74300864791",
		LastLoginDate:       "2015-10-05 12:43:35",
		RegistrationDate:    "2013-04-21 15:12:09",
		RegistrationAddress: "Nidakule Goztepe, Merdivenkoy Mah. Bora Sok. No:1",
		IP:                  "85.34.78.112",
		City:                "Istanbul",
		Country:             "Turkey",
		ZipCode:             "34732",
	}
)

func TestOne(t *testing.T) {
	testCredentialChecker(t)
}

/*
=============== CARD TESTS ===============
*/
func TestCreateCard(t *testing.T) {
	testCredentialChecker(t)

	testList := []CreateCardRequest{
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			Email:          "email@email.com",
			ExternalId:     "external id",
			ConversationId: "123456",
			Locale:         "tr",
		},
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			Email:      "email@email.com",
			ExternalId: "external id",
			Locale:     "tr",
		},
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			Email:          "email@email.com",
			ExternalId:     "external id",
			ConversationId: "123456",
		},
	}
	for index, testCase := range testList {
		resp, err := client.CreateCard(&testCase)
		if err != nil {
			log.Println(err)
			t.Fail()
		}
		if resp.Status != "success" {
			log.Println(resp)
			log.Println("Error at index: ", index)
			t.Fail()
		}
	}
}
func TestAddCard(t *testing.T) {
	testCredentialChecker(t)
	initialCreation, err := client.CreateCard(&CreateCardRequest{
		Card: CardInformation{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			CardAlias:      "card alias",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
		},
		Email:          "email@email.com",
		ExternalId:     "external id",
		ConversationId: "123456",
	})
	if err != nil {
		t.Fatal(err)
	}
	cuk := initialCreation.CardUserKey
	testList := []AddCardRequest{
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			CardUserKey:    cuk,
			ConversationId: "123456",
			Locale:         "tr",
		},
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			CardUserKey: cuk,
			Locale:      "tr",
		},
		{
			Card: CardInformation{
				CardHolderName: "John Doe",
				CardNumber:     "5528790000000008",
				CardAlias:      "card alias",
				ExpireMonth:    "12",
				ExpireYear:     "2030",
			},
			CardUserKey:    cuk,
			ConversationId: "123456",
		},
	}
	for index, testCase := range testList {
		resp, err := client.AddCard(&testCase)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Status != "success" {
			log.Println(resp)
			log.Println("Error at index: ", index)
			t.Fail()
		}
	}
}
func TestDeleteCard(t *testing.T) {
	testCredentialChecker(t)
	// Create a card to delete
	resp, err := client.CreateCard(&CreateCardRequest{
		Card: CardInformation{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			CardAlias:      "card alias",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
		},
		Email:          "email@email.com",
		ExternalId:     "external id",
		ConversationId: "123456",
		Locale:         "tr",
	})
	if err != nil || resp.Status != "success" {
		log.Fatal("Couldn't create card to remove.")
		t.Fail()
	}
	deleteResp, err := client.DeleteCard(&DeleteCardRequest{
		CardToken:   resp.CardToken,
		CardUserKey: resp.CardUserKey,
	})
	if err != nil {
		t.Fatal(err)
	}
	if deleteResp.Status != "success" {
		log.Fatal("Couldn't remove card.")
		t.Fail()
	}
}
func TestGetCard(t *testing.T) {
	testCredentialChecker(t)
	resp, err := client.CreateCard(&CreateCardRequest{
		Card: CardInformation{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			CardAlias:      "card alias",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
		},
		Email:          "email@email.com",
		ExternalId:     "external id",
		ConversationId: "123456",
		Locale:         "tr",
	})
	if err != nil || resp.Status != "success" {
		t.Fatal("Couldn't create card to get.")
	}
	getResp, err := client.GetCard(&GetCardRequest{
		CardUserKey: resp.CardUserKey,
	})
	if err != nil {
		t.Fatal(err)
	}
	if getResp.Status != "success" {
		t.Fatal("Error while getting card list.")
	}
}

/*
=========== INSTALLMENT INFO ============
*/
func TestInstallmentInformation(t *testing.T) {
	testCredentialChecker(t)
	reqList := []InstallmentRequest{
		{
			Price: "100",
		},
		{
			Price: "100.6",
		},
		{
			Price: "100.11",
		},
		{
			Price: "100.166",
		},
	}
	for ind, r := range reqList {
		resp, err := client.GetInstallmentInformation(&r)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Status != "success" {
			t.Fatal("Unsuccessful at index: ", ind)
		}
	}
}

/*
============= PAYMENT TESTS =============
*/
func TestPayment(t *testing.T) {
	testCredentialChecker(t)
	resp, err := client.CreatePayment(&CreatePaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1",
		PaidPrice:      "1.2",
		Installment:    "1",
		BasketID:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		Currency:       "TRY",
		PaymentCard: PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			Cvc:            "123",
			RegisterCard:   "0",
		},
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "success" {
		t.Fatal("Didn't get a success response.")
	}
}
func TestThreeDSInit(t *testing.T) {
	testCredentialChecker(t)
	r, err := client.ThreeDSInit(&ThreeDSInitRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1",
		PaidPrice:      "1.2",
		Installment:    "1",
		BasketID:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		Currency:       "TRY",
		PaymentCard: PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5170410000000004",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			Cvc:            "123",
			RegisterCard:   "0",
		},
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
		CallbackURL: "https://www.website.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get a success response while initializing 3DS.")
	}
}
func TestGetPayment(t *testing.T) {
	testCredentialChecker(t)

	// Create a payment to retrieve later.
	resp, err := client.CreatePayment(&CreatePaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1",
		PaidPrice:      "1.2",
		Installment:    "1",
		BasketID:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		Currency:       "TRY",
		PaymentCard: PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			Cvc:            "123",
			RegisterCard:   "0",
		},
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal("Encountered error while creating test payment. - ", err)
	}
	if resp.Status != "success" {
		t.Fatal("Didn't get success response while creating test payment.")
	}
	r, err := client.GetPayment(&GetPaymentRequest{
		PaymentId: resp.PaymentID,
		IP:        "192.168.10.10",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response.")
	}
}
func TestCancelPayment(t *testing.T) {
	testCredentialChecker(t)

	// Create a payment to cancel later.
	resp, err := client.CreatePayment(&CreatePaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1",
		PaidPrice:      "1.2",
		Installment:    "1",
		BasketID:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		Currency:       "TRY",
		PaymentCard: PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			Cvc:            "123",
			RegisterCard:   "0",
		},
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal("Encountered error while creating test payment. - ", err)
	}
	if resp.Status != "success" {
		t.Fatal("Didn't get success response while creating test payment.")
	}
	r, err := client.CancelPayment(&CancelPaymentRequest{
		IP:        "85.34.78.112",
		PaymentId: resp.PaymentID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get a success response while attempting to cancel payment.")
	}
}
func TestBKMInit(t *testing.T) {
	testCredentialChecker(t)
	r, err := client.BKMInit(&BkmInitRequest{
		Locale:          "tr",
		ConversationId:  "123456789",
		Price:           "1.0",
		CallbackURL:     "https://www.website.com/callback",
		BasketID:        "B67832",
		PaymentGroup:    "PRODUCT",
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response.")
	}
}
func TestBKMGet(t *testing.T) {
	testCredentialChecker(t)
	// Initialize payment
	r, err := client.BKMInit(&BkmInitRequest{
		Locale:          "tr",
		ConversationId:  "123456789",
		Price:           "1.0",
		CallbackURL:     "https://www.website.com/callback",
		BasketID:        "B67832",
		PaymentGroup:    "PRODUCT",
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response for test payment.")
	}

	resp, err := client.BKMGet(&BKMGetRequest{
		Token: r.Token,
	})
	if resp.ErrorCode != "20013" {
		t.Fatal("Didn't get desired error code while checking.")
	}
}
func TestRefundPayment(t *testing.T) {
	testCredentialChecker(t)
	resp, err := client.CreatePayment(&CreatePaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1",
		PaidPrice:      "1.2",
		Installment:    "1",
		BasketID:       "B67832",
		PaymentChannel: "WEB",
		PaymentGroup:   "PRODUCT",
		Currency:       "TRY",
		PaymentCard: PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireMonth:    "12",
			ExpireYear:     "2030",
			Cvc:            "123",
			RegisterCard:   "0",
		},
		Buyer:           dummyBuyer,
		BillingAddress:  dummyAddress,
		ShippingAddress: dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != "success" {
		t.Fatal("Didn't get a success response while creating test payment.")
	}

	r, err := client.RefundPayment(&RefundPaymentRequest{
		PaymentTransactionId: resp.ItemTransactions[0].PaymentTransactionID,
		IP:                   "85.34.78.112",
		Price:                sanitizeFPrice(resp.ItemTransactions[0].PaidPrice),
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response.")
	}
}

/*
================= FORMS =================
*/
func TestCheckoutFormInitialize(t *testing.T) {
	r, err := client.CheckoutFormInitialize(&CheckoutFormInitializeRequest{
		Price:           "1.0",
		PaidPrice:       "1.1",
		Buyer:           dummyBuyer,
		ShippingAddress: dummyAddress,
		BillingAddress:  dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
		Currency:"TRY",
		CallbackURL:"https://www.website.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response.")
	}
}
func TestCheckoutForm(t *testing.T) {
	r, err := client.CheckoutFormInitialize(&CheckoutFormInitializeRequest{
		Price:           "1.0",
		PaidPrice:       "1.1",
		Buyer:           dummyBuyer,
		ShippingAddress: dummyAddress,
		BillingAddress:  dummyAddress,
		BasketItems: BasketItems{
			{
				ID:        "BI101",
				Name:      "Binocular",
				Category1: "Collectibles",
				Category2: "Accessories",
				ItemType:  "PHYSICAL",
				Price:     "0.3",
			},
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
			{
				ID:        "BI103",
				Name:      "Usb",
				Category1: "Electronics",
				Category2: "Usb / Cable",
				ItemType:  "PHYSICAL",
				Price:     "0.2",
			},
		},
		Currency:"TRY",
		CallbackURL:"https://www.website.com",
	})
	if err != nil {
		t.Fatalf("Error while creating the form to get: %v", err)
	}
	if r.Status != "success" {
		t.Fatal("Didn't get success response while creating form to get.")
	}
	resp, err := client.CheckoutForm(&CheckoutFormRequest{
		Token:r.Token,
	})
	if resp.ErrorCode != "5122" {
		t.Fatalf("Didn't get the expected error code. Was expecting 5122, got: %v", resp.ErrorCode)
	}
}
/*
================== MISC =================
*/
func testCredentialChecker(t *testing.T) {
	if !(strings.Contains(sandboxKey, "sandbox") &&
		strings.Contains(sandboxSecret, "sandbox") &&
		strings.Contains(sandboxURL, "sandbox")) {
		t.Fatal("Please enter sandbox credentials to the variables defined in this file.")
	}
}
