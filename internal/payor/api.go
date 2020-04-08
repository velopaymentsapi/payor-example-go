package payor

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/antihax/optional"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	velo "github.com/velopaymentsapi/velo-go/v2"
	"golang.org/x/crypto/bcrypt"
)

type authClaims struct {
	Token string `json:"token"`
	jwt_lib.StandardClaims
}

func apiIndex(c *gin.Context) {
	c.JSON(200, "payor api index")
}

func apiAuth(c *gin.Context) {
	var jsonMap map[string]interface{}
	c.BindJSON(&jsonMap)

	repo := &repo{DB: c.MustGet("DB").(*gorm.DB)}
	user, err := repo.findUserByUsername(context.TODO(), jsonMap["username"].(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		c.Abort()
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonMap["password"].(string)))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		c.Abort()
		return
	}

	claims := authClaims{
		user.APIKey,
		jwt_lib.StandardClaims{
			Audience:  "api",
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "payor-example",
			Subject:   "authentication",
		},
	}

	// Create the token
	token := jwt_lib.NewWithClaims(jwt_lib.GetSigningMethod("HS256"), claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("FAILURE: apiAuth", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func apiVeloInfo(c *gin.Context) {
	cfg := velo.NewConfiguration()
	client := velo.NewAPIClient(cfg)
	auth := context.WithValue(context.TODO(), velo.ContextAccessToken, os.Getenv("VELO_API_ACCESSTOKEN"))
	info, _, err := client.PayorsApi.GetPayorById(auth, os.Getenv("VELO_API_PAYORID"))
	if err != nil {
		fmt.Println("FAILURE: apiVeloInfo", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}

	c.JSON(200, info)
}

func apiVeloAccounts(c *gin.Context) {
	cfg := velo.NewConfiguration()
	client := velo.NewAPIClient(cfg)
	auth := context.WithValue(context.TODO(), velo.ContextAccessToken, os.Getenv("VELO_API_ACCESSTOKEN"))
	saop := velo.GetSourceAccountsV2Opts{PayorId: optional.NewInterface(os.Getenv("VELO_API_PAYORID"))}
	accounts, _, err := client.FundingManagerApi.GetSourceAccountsV2(auth, &saop)
	if err != nil {
		fmt.Println("FAILURE: apiVeloAccounts", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}

	c.JSON(200, accounts)
}

func apiVeloFundAccount(c *gin.Context) {
	var jsonMap map[string]interface{}
	c.BindJSON(&jsonMap)

	inputamount := jsonMap["amount"].(float64)
	amount := int64(inputamount)

	cfg := velo.NewConfiguration()
	client := velo.NewAPIClient(cfg)
	auth := context.WithValue(context.TODO(), velo.ContextAccessToken, os.Getenv("VELO_API_ACCESSTOKEN"))
	args := velo.FundingRequestV1{Amount: amount}
	funding, err := client.FundingManagerApi.CreateAchFundingRequest(auth, jsonMap["source_account"].(string), args)
	if err != nil {
		fmt.Println("FAILURE: apiVeloFundAccount", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}
	if funding.StatusCode == 202 {
		var jsonRes map[string]interface{}
		c.JSON(200, jsonRes)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}
}

func apiVeloSupportedCountries(c *gin.Context) {
	cfg := velo.NewConfiguration()
	client := velo.NewAPIClient(cfg)
	auth := context.WithValue(context.TODO(), velo.ContextAccessToken, os.Getenv("VELO_API_ACCESSTOKEN"))
	countries, _, err := client.CountriesApi.ListSupportedCountriesV1(auth)
	if err != nil {
		fmt.Println("FAILURE: apiVeloSupportedCountries", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}

	c.JSON(200, countries)
}

func apiVeloSupportedCurrencies(c *gin.Context) {
	cfg := velo.NewConfiguration()
	client := velo.NewAPIClient(cfg)
	auth := context.WithValue(context.TODO(), velo.ContextAccessToken, os.Getenv("VELO_API_ACCESSTOKEN"))
	currencies, _, err := client.CurrenciesApi.ListSupportedCurrencies(auth)
	if err != nil {
		fmt.Println("FAILURE: apiVeloSupportedCurrencies", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Internal Server Error"})
		c.Abort()
		return
	}

	c.JSON(200, currencies)
}

func apiPayees(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiPayees",
	})
}

func apiCreatePayee(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiCreatePayee",
	})
}

func apiPayeeDetails(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiPayeeDetails",
	})
}

func apiVeloPayeeInvite(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiVeloPayeeInvite",
	})
}

func apiCreatePayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiCreatePayment",
	})
}

func apiInstructPayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiInstructPayment",
	})
}

func apiCancelPayment(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiCancelPayment",
	})
}

func apiPaymentDetails(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": "apiPaymentDetails",
	})
}
