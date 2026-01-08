package controller

import (
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var oderCollection *mongo.Collection = database,OpenCollection(database.Client, "order")

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		var table models.Menu 
		var order models.Order 

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("message: Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil{
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order model.OrderModel

		err := orderCollection.FindOne(ctx, bson.M{"order_id", orderId}).Decode(&order)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occure while fetching the order item"})
		}

		c.JSON(http.StatusOK, order)
}
}

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context){
	var err, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listening order items"})
		}

		var allOrders []bson.M

		if err = result.All(ctx, &allOrders); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context,WithTimeout(context.Background(), 100*time.second)
		var table models.Table 
		var order models.Order 

		var updatedObj primitive.D 

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id", food.Table_id}).Decode(&table)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("message: Menu was not found")
				c.JSON(http.StatusInternalServerError, g.H{"error": msg})
			}
			updatedObj = append(updatedObj, bson.E{"menu", order.Table_id})
		}

		order.Update_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
		updateObj = append(updateObj, bson.E("update_at", food.Updated_at));

		upsert := true

		filter := bson.M{"order_id", orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert
		}

		result, err := orderCollection.UpdateOne(
			ctx, filter, bson.D{{"$set", updatedObj}},
			&opt
		)

		if err != nil{
			mesg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Order",
	})
}
}

func orderItemOrderCreator(order model.Order)string{
			order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID
		order.Order_id = order.ID.Hex()

		orderCollection.InsertOne(ctx, order)
		defer cancel()

		return order.Order_id
}