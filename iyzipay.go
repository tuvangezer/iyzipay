package iyzipay

import (
	"github.com/pkg/errors"
	"net/http"
	"math/rand"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"bytes"
	"encoding/json"
	"strconv"
	"math"
	"reflect"
	"strings"
	)

const (
	invalidFields    = "Missing mandatory values!"
	randomStringSize = 8
	randomAlphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	versionText      = "iyzipay-go-0.0.0"
)

var (
	invalidFieldsErr = errors.New(invalidFields)
)

type Client struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
}

func Create(apiKey, apiSecret, baseURL string) *Client {
	c := &http.Client{}
	iyziClient := Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		baseURL:    baseURL,
		httpClient: c,
	}
	return &iyziClient
}

func (c *Client) request(method string, path string, request interface{}, pki *pkiBuilder) *http.Response {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(request)

	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(method, c.baseURL+path, buf)

	if err != nil {
		log.Fatal(err)
	}

	setHeader(req, c, pki)

	//fmt.Println(req)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func setHeader(r *http.Request, c *Client, p *pkiBuilder) {
	randString := getRandomString()
	pkiStr := p.GetReqString()

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-type", "application/json")
	r.Header.Add("Authorization", createAuthString(c, randString, pkiStr))
	r.Header.Add("x-iyzi-rnd", randString)
	r.Header.Add("x-iyzi-client-version", versionText)
}

// Creates a random string of length and characters specified in the constants.
func getRandomString() string {
	b := make([]byte, randomStringSize)
	for i := range b {
		b[i] = randomAlphabet[rand.Intn(len(randomAlphabet))]
	}
	return string(b)
}

func createAuthString(c *Client, randString string, pki string) string {
	hashStr := c.apiKey + randString + c.apiSecret + pki

	hasher := sha1.New()
	hasher.Write([]byte(hashStr))

	bs := hasher.Sum(nil)

	encoded := base64.StdEncoding.EncodeToString(bs)
	return "IYZWS " + c.apiKey + ":" + encoded
}

/*
====== UTILITY FUNCTIONS ======
 */
func sanitizePrice(priceStr string) string {
	var s string
	if priceStr != "" {
		f, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatal("Couldn't convert float.", err)
		}
		if math.Ceil(f) == f {
			s = strconv.FormatFloat(f, 'f', 0, 64)
			s += ".0"
		} else if math.Ceil(f*10) == f*10 {
			rounded := math.Round(f*10) / 10
			s = strconv.FormatFloat(rounded, 'f', 1, 64)
		} else {
			rounded := math.Round(f*100) / 100
			s = strconv.FormatFloat(rounded, 'f', 2, 64)
		}
	}
	return s
}
func sanitizeFPrice(price float64) string {
	return sanitizePrice(strconv.FormatFloat(price, 'f', 10, 64))
}
// valid Checks if every string field without the tag omitempty are entered
// Simple validation.
func valid(s interface{}) bool {
	ref := reflect.TypeOf(s)
	for i := 0 ; i<ref.NumField();i++{
		if ref.Field(i).Type.String() == "string" {
			tag := ref.Field(i).Tag.Get("json")
			if tag == "" || tag == "-" {
				continue
			}
			if !strings.Contains(tag, "omitempty") {
				val := reflect.ValueOf(s).Field(i).String()
				if val == "" {
					return false
				}
			}
		}
	}
	return true
}