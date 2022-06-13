package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"grab-hack-for-good/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/spf13/viper"
)

type grabExpressUsecase struct {
	grabExpressRepo domain.GrabExpressRepository
	contextTimeout  time.Duration
}

var GrabExpressToken domain.GrabExpressToken

var grabExpressURL = struct {
	GetToken              string
	GetDeliveryQuotes     string
	CreateDeliveryRequest string
	GetDeliveryDetails    string
	CancelDelivery        string
}{
	"https://partner-api.stg-myteksi.com/grabid/v1/oauth2/token",
	"https://partner-api.stg-myteksi.com/grab-express-sandbox/v1/deliveries/quotes",
	"https://partner-api.stg-myteksi.com/grab-express-sandbox/v1/deliveries",
	"https://partner-api.stg-myteksi.com/grab-express-sandbox/v1/deliveries/%s",
	"https://partner-api.stg-myteksi.com/grab-express-sandbox/v1/merchant/deliveries/%s",
}

// NewGrabExpressUsecase will create new an grabExpressUsecase object representation of domain.GrabExpressUsecase interface
func NewGrabExpressUsecase(cr domain.GrabExpressRepository, timeout time.Duration) domain.GrabExpressUsecase {
	return &grabExpressUsecase{
		grabExpressRepo: cr,
		contextTimeout:  timeout,
	}
}

func (usecase *grabExpressUsecase) GetToken(c context.Context) (err error) {
	var grabExpressCredential = &domain.GrabExpressCredential{
		ClientId:     viper.GetString("grab.client_id"),
		ClientSecret: viper.GetString("grab.client_secret"),
		GrantType:    viper.GetString("grab.grant_type"),
		Scope:        viper.GetString("grab.scope"),
	}

	// ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	// defer cancel()
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(grabExpressCredential)
	if err != nil {
		return
	}

	res, err := http.Post(grabExpressURL.GetToken, "application/json", &buf)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &GrabExpressToken)
	if err != nil {
		return err
	}

	return
}

func (usecase *grabExpressUsecase) GetDeliveryQuotes(ctx context.Context, grabExpressDeliveryQuotesRequest *domain.GrabExpressDeliveryQuotesRequest) (grabExpressDeliveryQuotesResponse domain.GrabExpressDeliveryQuotesResponse, err error) {
	if GrabExpressToken.AccessToken == "" || GrabExpressToken.ExpiresIn < time.Now().Unix() {
		err = usecase.GetToken(ctx)
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(grabExpressDeliveryQuotesRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", grabExpressURL.GetDeliveryQuotes, &buf)
	if err != nil {
		return
	}

	bearer := "Bearer " + GrabExpressToken.AccessToken
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &grabExpressDeliveryQuotesResponse)
	if err != nil {
		return
	}

	return
}

func (usecase *grabExpressUsecase) CreateDeliveryRequest(ctx context.Context, grabExpressDeliveryRequest *domain.GrabExpressDeliveryRequest) (grabExpressDeliveryResponse domain.GrabExpressDeliveryResponse, err error) {
	if GrabExpressToken.AccessToken == "" || GrabExpressToken.ExpiresIn < time.Now().Unix() {
		err = usecase.GetToken(ctx)
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(grabExpressDeliveryRequest)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", grabExpressURL.CreateDeliveryRequest, &buf)
	if err != nil {
		return
	}

	bearer := "Bearer " + GrabExpressToken.AccessToken
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &grabExpressDeliveryResponse)
	if err != nil {
		return
	}

	return
}

func (usecase *grabExpressUsecase) QueueCreateDeliveryRequest(ctx context.Context, order *domain.Order, transaction *domain.Transaction) (err error) {
	if GrabExpressToken.AccessToken == "" || GrabExpressToken.ExpiresIn < time.Now().Unix() {
		err = usecase.GetToken(ctx)
	}

	grabExpressDeliveryRequest := domain.GrabExpressDeliveryRequest{
		MerchantOrderId: order.Id.Hex(),
		ServiceType:     "INSTANT",
		PaymentMethod:   "CASHLESS",
	}

	var _package domain.Package
	// var _dimensions domain.Dimensions
	for i := 0; i < len(order.Manifest.Products); i++ {
		_package.Name = order.Manifest.Products[i].Product.Name
		_package.Description = order.Manifest.Products[i].Product.Description
		_package.Price = order.Manifest.Products[i].Product.Price
		_package.Quantity = order.Manifest.Products[i].Quantity
		_package.Dimensions = order.Manifest.Products[i].Product.Dimensions
		// _package.Dimensions = domain.Dimensions{
		// 	Height: 0,
		// 	Width:  0,
		// 	Depth:  0,
		// 	Weight: 0,
		// }
		_package.Dimensions = order.Manifest.Products[i].Product.Dimensions

		grabExpressDeliveryRequest.Packages = append(grabExpressDeliveryRequest.Packages, _package)
	}

	grabExpressDeliveryRequest.Origin = domain.Origin{
		Address:     order.Manifest.Merchant.Address,
		Coordinates: order.Manifest.Merchant.Coordinate,
	}

	grabExpressDeliveryRequest.Destination = domain.Destination{
		Address:     order.DropOffLocation.Address,
		Coordinates: order.DropOffLocation.Coordinate,
	}
	grabExpressDeliveryRequest.Sender = domain.Sender{
		FirstName:   order.Manifest.Merchant.Name,
		CompanyName: order.Manifest.Merchant.Name,
		Email:       order.Manifest.Merchant.Email,
		Phone:       order.Manifest.Merchant.Phone,
		SmsEnabled:  true,
	}

	grabExpressDeliveryRequest.Recipient = domain.Recipient{
		FirstName:   order.DropOffLocation.PIC,
		CompanyName: order.DropOffLocation.Name,
		Phone:       order.DropOffLocation.Phone,
		Email:       order.DropOffLocation.Email,
		SmsEnabled:  true,
	}

	now := time.Now()
	yyyy, mm, dd := now.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 11, 0, 0, 0, now.Location())
	grabExpressDeliveryRequest.Schedule = domain.Schedule{
		PickupTimeFrom: string(tomorrow.Format(time.RFC3339)),
		PickupTimeTo:   string(tomorrow.Add(time.Hour).Format(time.RFC3339)),
	}

	sqsRequest := domain.SQSGrabExpressDeliveryRequest{
		Payload:            grabExpressDeliveryRequest,
		GrabToken:          GrabExpressToken.AccessToken,
		GrabEndpointURL:    viper.GetString("aws.grab_endpoint_url"),
		BackendEndpointURL: viper.GetString("aws.backend_endpoint_url"),
	}

	messageBody, err := json.Marshal(sqsRequest)
	if err != nil {
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)

	queueUrl := viper.GetString("aws.grab_create_delivery_queue_url")
	_, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    &queueUrl,
	})

	return
}

func (usecase *grabExpressUsecase) CancelDelivery(ctx context.Context, order domain.Order) (err error) {
	if GrabExpressToken.AccessToken == "" || GrabExpressToken.ExpiresIn < time.Now().Unix() {
		err = usecase.GetToken(ctx)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf(grabExpressURL.CancelDelivery, *order.DeliveryId), nil)
	if err != nil {
		return
	}

	bearer := "Bearer " + GrabExpressToken.AccessToken
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()

	return
}
