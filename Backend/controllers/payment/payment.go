package payment

import (
	"github.com/gin-gonic/gin"
	"limbic/models"
	"net/http"
)

type RequestForCreatePayment struct {
	Amount int `json:"amount"`
}

func (h handler) CreatePayment(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	var request RequestForCreatePayment

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	payment := models.Payment{Amount: request.Amount, UserId: int(user.ID)}

	h.DB.Create(&payment)
	c.JSON(http.StatusOK, payment)
}

type Request struct {
	InvoiceId     int
	TransactionId int
	Amount        int
}

func (h handler) ConfirmPayment(c *gin.Context) {
	requestData := Request{}
	err := c.Bind(&requestData)
	if err != nil {
		return
	}

	var payment models.Payment
	h.DB.First(&payment, requestData.InvoiceId)
	payment.SystemPaymentId = requestData.TransactionId
	h.DB.Save(&payment)

	c.JSON(http.StatusOK, "success")
}
