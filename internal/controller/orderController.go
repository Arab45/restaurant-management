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

var validate = validator.New()

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order for a table
// @Tags Order
// @Accept json
// @Produce json
// @Param order body model.OrderModel true "Order data (table_id required)"
// @Success 200 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]string "Bad request - validation error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /order [post]
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var orderCollection = database.Collection("orders")
		var tableCollection = database.Collection("tables")
		var table model.MenuModel
		var order model.OrderModel

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := "message: Table was not found"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := "order item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// GetOrder godoc
// @Summary Get a specific order
// @Description Retrieve order details by order ID
// @Tags Order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} model.OrderModel "Order details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /order/{id} [get]
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderCollection = database.Collection("orders")
		orderId := c.Param("order_id")
		var order model.OrderModel

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occure while fetching the order item"})
		}

		c.JSON(http.StatusOK, order)
	}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Retrieve a list of all orders
// @Tags Order
// @Produce json
// @Success 200 {object} []model.OrderModel "List of orders"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orders [get]
func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderCollection = database.Collection("orders")
		result, err := orderCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listening order items"})
		}

		var allOrders []bson.M

		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update order details by order ID
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body model.OrderModel true "Updated order data"
// @Success 200 {object} map[string]interface{} "Order updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /order/{id} [put]
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderCollection = database.Collection("orders")
		var menuCollection = database.Collection("menus")
		var table model.TableModel
		var order model.OrderModel

		var updatedObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := "message: Menu was not found"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			}
			updatedObj = append(updatedObj, bson.E{Key: "menu", Value: order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		// upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.Update().SetUpsert(true)

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{Key: "$set", Value: updatedObj}},
			opt,
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

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by order ID
// @Tags Order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]string "Order deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /order/{id} [delete]
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Delete Order",
		})
	}
}

func orderItemOrderCreator(order model.OrderModel) string {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var orderCollection = database.Collection("orders")
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}
