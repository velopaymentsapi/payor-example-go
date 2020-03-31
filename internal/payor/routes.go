package payor

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitRoutes will define all api endpoints
func InitRoutes(db *gorm.DB) *gin.Engine {
	gin.SetMode("debug")

	r := gin.Default()

	// Authorization: Bearer `token`
	r.Use(BearerTokenAuthMiddleware(db))
	{
		r.GET("/", apiIndex)
		r.POST("/auth/login", apiAuth)

		r.GET("/settings", apiVeloInfo)
		r.GET("/settings/accounts", apiVeloAccounts)
		r.POST("/settings/fundings", apiVeloFundAccount)
		r.GET("/settings/countries", apiVeloSupportedCountries)
		r.GET("/settings/currencies", apiVeloSupportedCurrencies)

		r.GET("/payees", apiPayees)
		r.POST("/payees", apiCreatePayee)
		r.GET("/payees/:payeeid", apiPayeeDetails)
		r.POST("/payees/:payeeid/invite", apiVeloPayeeInvite)

		r.POST("/payments", apiCreatePayment)
		r.PUT("/payments/:paymentid", apiInstructPayment)
		r.DELETE("/payments/:paymentid", apiCancelPayment)
		r.GET("/payments/:paymentid", apiPaymentDetails)

	}
	return r
}
