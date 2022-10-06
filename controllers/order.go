package controllers

import (
	"fmt"
	"net/http"
	"time"
	"assignment-2-golang-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateOrder(ctx *gin.Context){
	var (
		newOrder 	models.Order
	)

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error creating order data: %v", err),
		})
		return
	}
	
	newOrder.OrderedAt = time.Now()
	
	if err := c.DB.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error creating order data: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"result": newOrder,
	})
}

func (c *Controller) GetOrders(ctx *gin.Context) {
	var (
		orders 		[]models.Order
	)

	if err := c.DB.Model(&models.Order{}).Preload("Items").Order("order_id asc").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error getting order datas: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}

func (c *Controller) UpdateOrder(ctx *gin.Context) {
	var (
		orderId 		string
		updatedOrder	models.Order
		countItem		int
	)

	orderId = ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v not found", orderId),
		})
		return
	}

	if err := c.DB.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v not found", orderId),
		})
		return
	}

	if err := c.DB.Model(&models.Order{}).Where("order_id = ?", orderId).Updates(updatedOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error updating order data: %v", err.Error()),
		})
		return
	}

	countItem = 0
	for _, item := range updatedOrder.Items {
		if err := c.DB.First(&models.Item{}, "order_id = ? and item_id = ?", orderId, item.ItemId).Error; err == nil {
			c.DB.Model(&models.Item{}).Where("item_id = ?", item.ItemId).Updates(item)
			countItem += 1
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v has been successfully updated (%v items updated)", orderId, countItem),
	})
}

func (c *Controller) DeleteOrder(ctx *gin.Context) {
	var (
		orderId 	string
	)

	orderId = ctx.Param("orderId")

	if err := c.DB.First(&models.Order{}, "order_id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v not found", orderId),
		})
		return
	}

	if err := c.DB.Where("order_id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error deleting order: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v has been successfully deleted", orderId),
	})
}