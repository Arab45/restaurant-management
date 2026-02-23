package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

// CreateInvoice godoc
// @Summary Create a new invoice
// @Description Create a new invoice for an order with payment status and method
// @Tags Invoice
// @Accept json
// @Produce json
// @Param invoice body model.InvoiceModel true "Invoice data (order_id required)"
// @Success 200 {object} map[string]interface{} "Invoice created successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoice [post]
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var orderCollection = database.Collection("orders")
		var invoiceCollection = database.Collection("invoices")
		var invoice model.InvoiceModel

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var order model.OrderModel

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		defer cancel()

		if err != nil {
			msg := "message: Order was not found"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_id = invoice.ID.Hex()

		validate := validator.New()
		validationErr := validate.Struct(invoice)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := invoiceCollection.InsertOne(ctx, invoice)

		if err != nil {
			msg := "invoice item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// GetInvoice godoc
// @Summary Get a specific invoice
// @Description Retrieve invoice details with order items by invoice ID
// @Tags Invoice
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} InvoiceViewFormat "Invoice details with order information"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoice/{id} [get]
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")

		var invoiceCollection = database.Collection("invoices")
		var invoice model.InvoiceModel

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing invoice item"})
		}

		var invoiceView InvoiceViewFormat

		allOrderItem, err := ItemByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = invoice.Payment_status
		invoiceView.Payment_due = allOrderItem[0]["payment_due"]
		invoiceView.Table_number = allOrderItem[0]["table_number"]
		invoiceView.Order_details = allOrderItem[0]["order_items"]

		c.JSON(http.StatusOK, invoiceView)
	}
}

// GetInvoices godoc
// @Summary Get all invoices
// @Description Retrieve a list of all invoices
// @Tags Invoice
// @Produce json
// @Success 200 {object} []model.InvoiceModel "List of invoices"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoices [get]
func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoiceCollection = database.Collection("invoices")
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occur"})
		}

		var allInvoice []bson.M
		if err := result.All(ctx, &allInvoice); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoice)
	}
}

// UpdateInvoice godoc
// @Summary Update an invoice
// @Description Update invoice payment status and payment method by invoice ID
// @Tags Invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param invoice body model.InvoiceModel true "Updated invoice data (payment_method and payment_status)"
// @Success 200 {object} map[string]interface{} "Invoice updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoice/{id} [put]
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var invoiceCollection = database.Collection("invoices")
		var invoice model.InvoiceModel
		invoiceId := c.Param("invoice_id")

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"invoice_id": invoiceId}

		var updatedObj primitive.D

		if invoice.Payment_method != nil {
			updatedObj = append(updatedObj, bson.E{Key: "payment_method", Value: invoice.Payment_method})
		}

		if invoice.Payment_status != nil {
			updatedObj = append(updatedObj, bson.E{Key: "payment_status", Value: invoice.Payment_status})

		}

		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: invoice.Updated_at})

		opt := options.Update().SetUpsert(true)

		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		result, err := invoiceCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{
				Key:   "$set",
				Value: updatedObj,
			}},
			opt,
		)

		if err != nil {
			msg := "invoice item update filed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// DeleteInvoice godoc
// @Summary Delete an invoice
// @Description Delete an invoice by invoice ID
// @Tags Invoice
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} map[string]string "Invoice deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoice/{id} [delete]
func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Delete Invoice",
		})
	}
}
