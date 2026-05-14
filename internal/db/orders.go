package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (o *ordersRepo) CreateOrder(ctx context.Context, customerID, vendorID int64, req *request.PlaceOrderRequest) (string, error) {
	tx := o.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return "", tx.Error
	}

	orderID := "ord_" + strconv.FormatInt(global.SnowFlake.Generate().Int64(), 10)
	totalAmount := decimal.NewFromInt(0)

	for _, item := range req.Items {
		// TODO: Optimize these queries
		var price decimal.Decimal
		// Can cache such values for quick responses
		if err := tx.Model(&model.Meal{}).
			Where("id = ? AND vendor_id = ?", item.MealID, vendorID).
			Select("price").
			Scan(&price).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// We could email the user in case one of their items no longer exists :)
				continue
			}
			tx.Rollback()
			return "", fmt.Errorf("fetch meal price: %w", err)
		}

		totalAmount = totalAmount.Add(price.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	order := &model.Order{
		ID:              orderID,
		CustomerID:      customerID,
		VendorID:        vendorID,
		Status:          "pending",
		TotalAmount:     totalAmount,
		DeliveryFee:     decimal.NewFromInt(constants.FAKE_DELIVERY),
		DeliveryAddress: req.DeliveryAddress,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("insert order: %w", err)
	}

	now := time.Now()

	// Fetch meal prices
	mealIds := make([]int64, 0, len(req.Items))
	for _, v := range req.Items {
		mealIds = append(mealIds, v.MealID)
	}

	var meals []*model.Meal
	if err := tx.Select("id", "price").Where("id = ANY(?)", mealIds).Find(&meals).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	prices := make(map[int64]decimal.Decimal)
	for _, meal := range meals {
		prices[meal.ID] = decimal.NewFromInt(meal.Price)
	}

	// Create order items
	var orderItems []*model.OrderItem
	for _, item := range req.Items {
		price, ok := prices[item.MealID]
		if !ok {
			continue
		}
		orderItemID := global.SnowFlake.Generate().Int64()
		orderItems = append(orderItems, &model.OrderItem{
			ID:        orderItemID,
			OrderID:   orderID,
			MealID:    item.MealID,
			Quantity:  item.Quantity,
			Price:     price,
			CreatedAt: now,
		})
	}

	if len(orderItems) > 0 {
		if err := tx.CreateInBatches(orderItems, 100).Error; err != nil {
			tx.Rollback()
			return "", fmt.Errorf("error creating order items: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return orderID, nil
}

func (o *ordersRepo) CancelOrder(ctx context.Context, customerID int64, orderID string) error {
	result := o.db.WithContext(ctx).Model(&model.Order{}).
		Where("id = ? AND customer_id = ? AND status = 'pending'", orderID, customerID).
		Update("status", "cancelled")

	if result.Error != nil {
		return fmt.Errorf("cancel order: %w", result.Error)
	}

	return nil
}

func (o *ordersRepo) GetOrder(ctx context.Context, customerID int64, orderID string) (*model.Order, error) {
	order := &model.Order{}
	if err := o.db.WithContext(ctx).Where("id = ? AND customer_id = ?", orderID, customerID).First(order).Error; err != nil {
		return nil, fmt.Errorf("fetch order: %w", err)
	}

	return order, nil
}
