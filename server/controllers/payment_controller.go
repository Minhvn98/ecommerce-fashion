package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Minhvn98/ecommerce-fashion/config"
	repo "github.com/Minhvn98/ecommerce-fashion/repository"
)

type Payload struct {
	PartnerCode string `json:"partnerCode"`
	AccessKey   string `json:"accessKey"`
	RequestID   string `json:"requestId"`
	Amount      string `json:"amount"`
	OrderID     string `json:"orderId"`
	OrderInfo   string `json:"orderInfo"`
	ReturnURL   string `json:"returnUrl"`
	NotifyURL   string `json:"notifyUrl"`
	ExtraData   string `json:"extraData"`
	RequestType string `json:"requestType"`
	Signature   string `json:"signature"`
}

func PaymentWithMomo(userId int, orderId int) {
	order := repo.GetOrderById(orderId)
	productName := ""
	for _, product := range order.Products {
		productName += product.Product.Name + ". "
	}

	var requestId = strconv.Itoa(userId)
	var endpoint = config.Config.Payment.Momo.ApiEndpoint
	var partnerCode = config.Config.Payment.Momo.PartnerCode
	var accessKey = config.Config.Payment.Momo.AccsessKey
	var serectkey = config.Config.Payment.Momo.SecretKey
	var orderInfo = "Đơn hàng của " + order.FirstName + order.LastName + ". " + productName
	var returnUrl = "http://localhost:3000/api/v1/payment"
	var notifyurl = "http://localhost:3000/api/v1/payment"
	var amount = strconv.Itoa(order.TotalPrice)
	var requestType = "captureMoMoWallet"
	var extraData = "merchantName=fashionshop;merchantId=1102"

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(amount)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(strconv.Itoa(orderId))
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&returnUrl=")
	rawSignature.WriteString(returnUrl)
	rawSignature.WriteString("&notifyUrl=")
	rawSignature.WriteString(notifyurl)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(serectkey))

	// Write Data to it
	hmac.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))

	var payload = Payload{
		PartnerCode: partnerCode,
		AccessKey:   accessKey,
		RequestID:   requestId,
		Amount:      amount,
		OrderID:     strconv.Itoa(orderId),
		OrderInfo:   orderInfo,
		ReturnURL:   returnUrl,
		NotifyURL:   notifyurl,
		ExtraData:   extraData,
		RequestType: requestType,
		Signature:   signature,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Payload: " + string(jsonPayload))
	fmt.Println("Signature: " + signature)

	//send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response from Momo: ", result)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	//PayUrl
	fmt.Println("PayUrl is: %s\n", result["payUrl"])
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	payload := r.Form
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": payload})
}

func PostPayment(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Ok"})
}
