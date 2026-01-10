package controller

import "github.com/gin-gonic/gin"

type InvoiceViewFormat struct {
	Invoice_id        string
	Payment_method    string
	Order_id          string
	Payment_status    *string
	Payment_due       interface{}
	Table_number      interface{}
	Payment_due_date  time.time
	Order_details     interface{}

}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)

		result, err  := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Invoice",
	})
}
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Invoices",
	})
}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Update Invoice",
	})
}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Invoice",
	})
}
}