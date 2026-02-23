package controller

import (
	"RESTAURANT-MANAGEMENT/internal/database"
	"RESTAURANT-MANAGEMENT/internal/model"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderItemPack struct {
	Table_id   *string
	Order_item []model.OrderItemModel
}

// CreateOrderItem godoc
// @Summary Create order items for an order
// @Description Create one or more order items associated with an order
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param orderItemPack body OrderItemPack true "Order items data with table_id and items list"
// @Success 200 {object} map[string]interface{} "Order items created successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orderItem [post]
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var OrderItemCollection = database.Collection("order_items")
		var orderItemPack OrderItemPack

		var order model.OrderModel

		if err := c.BindJSON(&orderItemPack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		now := time.Now()
		order.Order_Date = &now
		orderItemToBeInserted := []interface{}{}
		order.Table_id = orderItemPack.Table_id
		order_id := orderItemOrderCreator(order)

		for _, orderItem := range orderItemPack.Order_item {
			orderItem.Order_id = order_id

			validationErr := validate.Struct(orderItem)

			if validationErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
				return
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_item_id = orderItem.ID.Hex()
			var num = toFixed(*orderItem.Unit_price, 2)
			orderItem.Unit_price = &num
			orderItemToBeInserted = append(orderItemToBeInserted, orderItem)
		}

		insertOrderItems, err := OrderItemCollection.InsertOne(ctx, orderItemToBeInserted)
		if err != err {
			log.Fatal()
		}

		c.JSON(http.StatusOK, insertOrderItems)
	}

}

// GetOrderItemByOrder godoc
// @Summary Get order items by order ID
// @Description Retrieve all items for a specific order with details (food, table, etc)
// @Tags OrderItem
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} []map[string]interface{} "List of order items with details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orderItem/{id} [get]
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

// GetOrderItems godoc
// @Summary Get all order items
// @Description Retrieve a list of all order items across all orders
// @Tags OrderItem
// @Produce json
// @Success 200 {object} []model.OrderItemModel "List of order items"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orderItems [get]
func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var OrderItemCollection = database.Collection("order_items")

		result, err := OrderItemCollection.Find(context.TODO(), bson.M{})

		if err != err {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered items"})
			return
		}
		var allOrderItem []bson.M
		if err = result.All(ctx, &allOrderItem); err != nil {
			log.Fatal(err)
			return
		}

		c.JSON(http.StatusOK, allOrderItem)
	}
}

//performing lookup with order or food id
// func ItemByOrder( id string) (OrderItems []primitive.M, err error) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	matchStage := bson.D{
// 		{
// 			"$match",
// 			bson.D{{"order_id", id}},
// 		},
// 	}
// 	//lookup to db with a particular id collection
// 	lookupStage := bson.D{
// 		{
// 			"$lookup",
// 			bson.D{
// 				{"from", "food"},
// 				{"localField", "food_id"},
// 				{"foreignField", "food_id"},
// 				{"as", "food"},
// 			},
// 		},
// 	}

// 	//having access to the array of object the lookup stage provide to unwind it
// 	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}}

// 	//order lookup
// 	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
// 	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

// 	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
// 	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

// 	projectStage := bson.D{
// 		{
// 			"$project",
// 			bson.D{
// 				{"id", 0},
// 				{"amount", "$food.price"},
// 				{"total_count", 1},
// 				{"food_name", "$food.name"},
// 				{"food_image", "$food_image"},
// 				{"table_number", "$table.table_number"},
// 				{"table_id", "$table.table_id"},
// 				{"order_id", "$order.order_id"},
// 				{"price", "$food_price"},
// 				{"quantity", 1},
// 			},
// 		},
// 	}

// 	groupStage := bson.D{
// 		{"$group", bson.D{{
// 			"_id",
// 			bson.D{
// 				{"order_id", "$order_id"},
// 				{"table_id", "$table_id"},
// 				{"table_number", "$table_number"},
// 			}},
// 			{"payment_due", bson.D{{"$sum", "$amount"}}},
// 			{"total_items", bson.D{}},
// 		}},
// 	}

// 	projectStage2 := bson.D{
// 		{"$project", bson.D{
// 			{"id", 0},
// 			{"payment_due", 1},
// 			{"total_count", 1},
// 			{"total_number", "$_id.table_number"},
// 			{"order_items", 1},
// 		}},
// 	}

// 	result, err := OrderItemCollection.Aggregate(
// 		ctx, mongo.Pipeline{
// 			matchStage,
// 			lookupStage,
// 			unwindStage,
// 			lookupOrderStage,
// 			unwindOrderStage,
// 			lookupTableStage,
// 			unwindTableStage,
// 			projectStage,
// 			groupStage,
// 			projectStage2,
// 		},
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = result.All(ctx, &OrderItems);  err != nil {
// 		panic(err)
// 	}

// 	defer cancel()

// 	return OrderItems, err
// }

func ItemByOrder(id string) (orderItems []bson.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var OrderItemCollection = database.Collection("order_items")

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{{Key: "order_id", Value: id}}},
	}

	lookupFoodStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "foods"},
			{Key: "localField", Value: "food_id"},
			{Key: "foreignField", Value: "food_id"},
			{Key: "as", Value: "food"},
		}},
	}

	unwindFoodStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$food"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}

	lookupOrderStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "orders"},
			{Key: "localField", Value: "order_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "order"},
		}},
	}

	unwindOrderStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$order"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}

	lookupTableStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "tables"},
			{Key: "localField", Value: "order.table_id"},
			{Key: "foreignField", Value: "table_id"},
			{Key: "as", Value: "table"},
		}},
	}

	unwindTableStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$table"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "amount", Value: "$food.price"},
			{Key: "food_name", Value: "$food.name"},
			{Key: "food_image", Value: "$food.food_image"},
			{Key: "table_number", Value: "$table.table_number"},
			{Key: "table_id", Value: "$table.table_id"},
			{Key: "order_id", Value: "$order.order_id"},
			{Key: "price", Value: "$food.price"},
			{Key: "quantity", Value: 1},
		}},
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "order_id", Value: "$order_id"},
				{Key: "table_id", Value: "$table_id"},
				{Key: "table_number", Value: "$table_number"},
			}},
			{Key: "payment_due", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
			{Key: "order_items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		}},
	}

	projectStage2 := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "payment_due", Value: 1},
			{Key: "table_number", Value: "$_id.table_number"},
			{Key: "order_items", Value: 1},
		}},
	}

	cursor, err := OrderItemCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		lookupFoodStage,
		unwindFoodStage,
		lookupOrderStage,
		unwindOrderStage,
		lookupTableStage,
		unwindTableStage,
		projectStage,
		groupStage,
		projectStage2,
	})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &orderItems); err != nil {
		return nil, err
	}

	return orderItems, nil
}

// UpdateOrderItem godoc
// @Summary Update an order item
// @Description Update order item details by order item ID
// @Tags OrderItem
// @Accept json
// @Produce json
// @Param id path string true "Order Item ID"
// @Param orderItem body model.OrderItemModel true "Updated order item data"
// @Success 200 {object} map[string]interface{} "Order item updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orderItem/{id} [put]
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var OrderItemCollection = database.Collection("order_items")

		var orderItem model.OrderItemModel

		orderItemId := c.Param("order_item_id")

		filter := bson.M{"order_item_id": orderItemId}

		var updatedObj primitive.D

		if orderItem.Unit_price != nil {
			updatedObj = append(updatedObj, bson.E{Key: "unit_price", Value: *orderItem.Unit_price})
		}

		if orderItem.Quantity != nil {
			updatedObj = append(updatedObj, bson.E{Key: "quantity", Value: *orderItem.Quantity})
		}

		if orderItem.Food_id != nil {
			updatedObj = append(updatedObj, bson.E{Key: "food_id", Value: *orderItem.Food_id})
		}

		orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: orderItem.Updated_at})

		opts := options.Update().SetUpsert(true)

		result, err := OrderItemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updatedObj},
			},
			opts,
		)

		if err != nil {
			msg := "order item update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)
	}

}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var OrderItemCollection = database.Collection("order_items")

		orderItemId := c.Param("order_item_id")
		var orderItem model.OrderItemModel

		err := OrderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered item"})
			return
		}

		c.JSON(http.StatusOK, orderItem)
	}

}
