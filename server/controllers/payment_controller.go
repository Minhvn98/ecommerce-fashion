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

func PaymentWithMomo(userId int, orderId int) interface{} {
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
	// var returnUrl = config.Config.Payment.Momo.ReturnUrl
	// var notifyurl = config.Config.Payment.Momo.NotifyUrl + "/api/v1/payment"
	var returnUrl = "http://localhost:8080/bill-detail/" + strconv.Itoa(orderId)
	var notifyurl = "https://7fa52d2187ca.ngrok.io/api/v1/payment"

	var amount = strconv.Itoa(order.TotalPrice)
	var requestType = "captureMoMoWallet"
	var extraData = "merchantName=Fashion-Shop;merchantId=1102"

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

	//send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	//result
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result["payUrl"]
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	payload := r.Form
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": payload})
}

func PostPayment(w http.ResponseWriter, r *http.Request) {
	partnerCode := r.FormValue("partnerCode")
	accessKey := r.FormValue("accessKey")
	requestId := r.FormValue("requestId")
	amount := r.FormValue("amount")
	orderId := r.FormValue("orderId")
	orderInfo := r.FormValue("orderInfo")
	orderType := r.FormValue("orderType")
	transId := r.FormValue("transId")
	message := r.FormValue("message")
	localMessage := r.FormValue("localMessage")
	responseTime := r.FormValue("responseTime")
	errorCode := r.FormValue("errorCode")
	payType := r.FormValue("payType")
	extraData := r.FormValue("extraData")
	signature := r.FormValue("signature")

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
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&orderType=")
	rawSignature.WriteString(orderType)
	rawSignature.WriteString("&transId=")
	rawSignature.WriteString(transId)
	rawSignature.WriteString("&message=")
	rawSignature.WriteString(message)
	rawSignature.WriteString("&localMessage=")
	rawSignature.WriteString(localMessage)
	rawSignature.WriteString("&responseTime=")
	rawSignature.WriteString(responseTime)
	rawSignature.WriteString("&errorCode=")
	rawSignature.WriteString(errorCode)
	rawSignature.WriteString("&payType=")
	rawSignature.WriteString(payType)
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	serectkey := config.Config.Payment.Momo.SecretKey
	hmac := hmac.New(sha256.New, []byte(serectkey))

	// Write Data to it

	hmac.Write(rawSignature.Bytes())
	// fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signatureMomo := hex.EncodeToString(hmac.Sum(nil))
	fmt.Println("------ ", signature)
	fmt.Println("------ ", signatureMomo)
	if signature == signatureMomo {
		// update status order
		id, _ := strconv.Atoi(orderId)
		if errorCode == "0" {
			repo.UpdateStatusOrder("Đã thanh toán", id)
		} else {
			repo.UpdateStatusOrder("Thanh toán thất bại", id)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Ok"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Signature not match"})
	}
}
