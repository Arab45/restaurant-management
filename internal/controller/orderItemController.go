package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItemPack struct {
	Table_id *string
	Order_item  []models.OrderItem
}

var OrderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")


func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderItemPack orderItemPack

		var order models.Order

		if err := c.BindJSON(&orderItemPack); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		orderItemToBeInserted := []interface{}
		order.Table_id = orderItemPack.Table_id
		order_id := OrderItemOrderCreator(order)

		for _, orderItem := range orderItemPack.Order_item{
			orderItem.Order_id = order_id

			validationErr := validate.Struct(orderItem)

			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_item_id = orderItem.ID.Hex()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemsToBeInserted = append(orderItemToBeInserted, orderItem)
		}

		insertOrderItems, err := OrderItemCollection.InsertOne(ctx, orderItemsToBeInserted)
		if err != err{
			log.Fatal()
		}

		c.JSON(http.StatusOK, insertOrderItems)
	}

}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

	 allOrderItem, err := ItemByOrder(orderId)

	 if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to order itemby orderId"})
		return
	 }
	 c.JSON(http.StatusOK, allOrderItem)
}
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := OrderItemCollection.Find(context.TODO(), bson.M{})

		if err != err {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered items"})
			return
		}
		var allOrderItem []bson.M 
		if err = result.All(ctx, &allOrderItem); err != nil{
			log.Fatal(err)
			return
		}

		c.JSON(http.StatusOK, allOrderItem)
	}
}

//performing lookup with order or food id
func ItemByOrder( id string) (OrderItem []primitive.M, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	matchStage := bson.D{
		{
			"$match", 
			bson.D{{"order_id", id}}
		}
	}
	//lookup to db with a particular id collection
	lookupStage := bson.D{
		{
			"$lookup",
			bson.D{
				{"from", "food"}, 
				{"localField", "food_id"}, 
				{"foreignField", "food_id"}, 
				{"as", "food"}
			}
		}
	}

	//having access to the array of object the lookup stage provide to unwind it
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}} 

	//order lookup
	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

	projectStage := bson.D{
		{
			"$project",
			bson.D{
				{"id", 0},
				{"amount", "$food.price"},
				{"total_count", 1},
				{"food_name", "$food.name"},
				{"food_image", "$food_image"},
				{"table_number", "$table.table_number"},
				{"table_id", "$table.table_id"},
				{"order_id", "$order.order_id"},
				{"price", "$food_price"},
				{"quantity", 1}
			}
		}
	}

	groupStage := bson.D{
		{"$group", bson.D{{
			"_id",
			bson.D{
				{"order_id", "$order_id"},
				{"table_id", "$table_id"},
				{"table_number", "$table_number"}
			}},
			{"payment_due", bson.D{{"$sum", "$amount"}}},
			{"total_items", bson.D{}}
		}}
	}

	projectStage2 := bson.D{
		{"$project", bson.D{
			{"id", o},
			{"payment_due", 1},
			{"total_count", 1},
			{"total_number", "$_id.table_number"},
			{"order_items", 1}
		}}
	}

	result err := orderItemCollection.Aggregrate(
		ctx, mongo.Pipeline{
			matchStage,
			lookupStage,
			unwindStage,
			lookupOrderStage,
			unwindOrderStage,
			lookupTableStage,
			unwindTableStage,
			projectStage,
			groupStage,
			projectStage2
		}
	)

	if err != nil {
		panic(err)
	}

	if err = result.All(ctx, &OrderItems);  err != nil {
		panic(err)
	}

	defer cancel()

	return OrderItems, err
}

func UpdateOrderItem() gin.HandlerFunc {
	return func (c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var orderItem models.OrderItem 

		orderItemId := c.Param("order_item_id")

		filter := bson.M{"order_item_id": orderItemId} 

		var uodatedObj primitive.D 

		if orderItem.Unit_price != nil{
			updatedObj = append(updatedObj, bson.E{"unit_price": *&orderItem.Unit_price})
		}

		if orderItem.Quantity != nil {
			updatedObj = append(updatedObj, bson.E{"quantity": *orderItem.Quantity})
		}

		if orderItem.Food_id != nil {
			updatedObj = append(updateObj, bson.E{"food_id": *orderItem.Food_id})
		}

		orderItem.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{"updated_at": OrderItem.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert
		}

		result, err := OrderItemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set": updatedObj}
			},
			&opt
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)
	}
	
}

func GetOrderItem() gin.HandlerFunc {
	return func (c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time,Second)
		defer cancel()

		orderItemId := c.Param("order_item_id")
		var orderItem models.OrderItem 

		err := OrderItemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered item"})
			return
		}

		C.JSON(http.StatusOK, orderItem)
}
	
}