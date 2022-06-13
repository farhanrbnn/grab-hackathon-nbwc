package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"grab-hack-for-good/helper"

	_grabExpressHttpDelivery "grab-hack-for-good/modules/grab_express/delivery/http"
	_grabExpressRepo "grab-hack-for-good/modules/grab_express/repository/mongo"
	_grabExpressUsecase "grab-hack-for-good/modules/grab_express/usecase"

	_userHttpDelivery "grab-hack-for-good/modules/user/delivery/http"
	_userRepo "grab-hack-for-good/modules/user/repository/mongo"
	_userUsecase "grab-hack-for-good/modules/user/usecase"

	_walletOriginHttpDelivery "grab-hack-for-good/modules/wallet_origin/delivery/http"
	_walletOriginRepo "grab-hack-for-good/modules/wallet_origin/repository/mongo"
	_walletOriginUsecase "grab-hack-for-good/modules/wallet_origin/usecase"

	_walletHttpDelivery "grab-hack-for-good/modules/wallet/delivery/http"
	_walletRepo "grab-hack-for-good/modules/wallet/repository/mongo"
	_walletUsecase "grab-hack-for-good/modules/wallet/usecase"

	_userWalletHttpDelivery "grab-hack-for-good/modules/user_wallet/delivery/http"
	_userWalletRepo "grab-hack-for-good/modules/user_wallet/repository/mongo"
	_userWalletUsecase "grab-hack-for-good/modules/user_wallet/usecase"

	_merchantHttpDelivery "grab-hack-for-good/modules/merchant/delivery/http"
	_merchantRepo "grab-hack-for-good/modules/merchant/repository/mongo"
	_merchantUsecase "grab-hack-for-good/modules/merchant/usecase"

	_productHttpDelivery "grab-hack-for-good/modules/product/delivery/http"
	_productRepo "grab-hack-for-good/modules/product/repository/mongo"
	_productUsecase "grab-hack-for-good/modules/product/usecase"

	_dropOffLocationHttpDelivery "grab-hack-for-good/modules/drop_off_location/delivery/http"
	_dropOffLocationRepo "grab-hack-for-good/modules/drop_off_location/repository/mongo"
	_dropOffLocationUsecase "grab-hack-for-good/modules/drop_off_location/usecase"

	_transactionHttpDelivery "grab-hack-for-good/modules/transaction/delivery/http"
	_transactionRepo "grab-hack-for-good/modules/transaction/repository/mongo"
	_transactionUsecase "grab-hack-for-good/modules/transaction/usecase"

	_orderHttpDelivery "grab-hack-for-good/modules/order/delivery/http"
	_orderRepo "grab-hack-for-good/modules/order/repository/mongo"
	_orderUsecase "grab-hack-for-good/modules/order/usecase"
)

func init() {
	viper.SetConfigFile(`../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// API
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// Restricted
	r := e.Group("")
	r.Use(middleware.JWT([]byte(viper.GetString("jwt.secret_key"))))

	e.Static("/face-images", "../public/uploads")
	e.Static("/log-images", "../public/uploads/logs")

	// Connect MongoDB
	// client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&readPreference=primary&ssl=false", viper.GetString(`database.user`), viper.GetString(`database.pass`), viper.GetString(`database.host`), viper.GetString(`database.port`))))
	// client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/%s?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=admin", viper.GetString(`database.host`), viper.GetString(`database.port`), viper.GetString(`database.name`))))
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", viper.GetString(`database.host`), viper.GetString(`database.port`))))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), timeoutContext)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("connected")

	// Database Client
	database := client.Database(viper.GetString(`database.name`))

	// Grab Express
	grabExpressCollection := database.Collection("grab-express")

	grabExpressRepo := _grabExpressRepo.NewMongoGrabExpressRepository(grabExpressCollection)
	grabExpressUsecase := _grabExpressUsecase.NewGrabExpressUsecase(grabExpressRepo, timeoutContext)

	// User
	userCollection := database.Collection("users")

	userRepo := _userRepo.NewMongoUserRepository(userCollection)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	// Merchant
	merchantCollection := database.Collection("merchants")

	merchantRepo := _merchantRepo.NewMongoMerchantRepository(merchantCollection)
	merchantUsecase := _merchantUsecase.NewMerchantUsecase(merchantRepo, timeoutContext)

	// User Wallet
	userWalletCollection := database.Collection("user-wallets")

	userWalletRepo := _userWalletRepo.NewMongoUserWalletRepository(userWalletCollection)
	userWalletUsecase := _userWalletUsecase.NewUserWalletUsecase(userWalletRepo, timeoutContext)

	// Wallet Origin
	walletOriginCollection := database.Collection("wallet-origins")

	walletOriginRepo := _walletOriginRepo.NewMongoWalletOriginRepository(walletOriginCollection)
	walletOriginUsecase := _walletOriginUsecase.NewWalletOriginUsecase(walletOriginRepo, timeoutContext)

	// Wallet
	walletCollection := database.Collection("wallets")

	walletRepo := _walletRepo.NewMongoWalletRepository(walletCollection)
	walletUsecase := _walletUsecase.NewWalletUsecase(walletRepo, timeoutContext)

	// Product
	productCollection := database.Collection("products")

	productRepo := _productRepo.NewMongoProductRepository(productCollection)
	productUsecase := _productUsecase.NewProductUsecase(productRepo, timeoutContext)

	// Drop Off Location
	dropOffLocationCollection := database.Collection("drop-off-locations")

	dropOffLocationRepo := _dropOffLocationRepo.NewMongoDropOffLocationRepository(dropOffLocationCollection)
	dropOffLocationUsecase := _dropOffLocationUsecase.NewDropOffLocationUsecase(dropOffLocationRepo, timeoutContext)

	// Transaction
	transactionCollection := database.Collection("transactions")

	transactionRepo := _transactionRepo.NewMongoTransactionRepository(transactionCollection)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(transactionRepo, timeoutContext)

	// Order
	orderCollection := database.Collection("orders")

	orderRepo := _orderRepo.NewMongoOrderRepository(orderCollection)
	orderUsecase := _orderUsecase.NewOrderUsecase(orderRepo, timeoutContext)

	// Handlers
	_grabExpressHttpDelivery.NewGrabExpressHandler(e, r, grabExpressUsecase)
	_userHttpDelivery.NewUserHandler(e, r, userUsecase)
	_merchantHttpDelivery.NewMerchantHandler(e, r, merchantUsecase)
	_userWalletHttpDelivery.NewUserWalletHandler(e, r, userWalletUsecase)
	_walletOriginHttpDelivery.NewWalletOriginHandler(e, r, walletOriginUsecase)
	_walletHttpDelivery.NewWalletHandler(e, r, walletUsecase)
	_productHttpDelivery.NewProductHandler(e, r, productUsecase)
	_dropOffLocationHttpDelivery.NewDropOffLocationHandler(e, r, dropOffLocationUsecase)
	_transactionHttpDelivery.NewTransactionHandler(e, r, transactionUsecase, grabExpressUsecase, orderUsecase, userWalletUsecase)
	_orderHttpDelivery.NewOrderHandler(e, r, orderUsecase, grabExpressUsecase, transactionUsecase, userWalletUsecase, dropOffLocationUsecase)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
