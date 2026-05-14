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
	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func (o *ordersRepo) CreateOrder(ctx context.Context, customerID, vendorID int64, req *request.PlaceOrderRequest) (string, error) {
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	orderID := "ord_" + strconv.FormatInt(global.SnowFlake.Generate().Int64(), 10)
	totalAmount := decimal.NewFromInt(0)

	for _, item := range req.Items {
		// TODO: Optimize these queries
		var price decimal.Decimal
		// Can cache such values for quick responses
		err = tx.QueryRow(ctx, `SELECT price FROM meals WHERE id = $1 AND vendor_id = $2`, item.MealID, vendorID).Scan(&price)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				// We could email the user in case one of their items no longer exists :)
				continue
			}
			return "", fmt.Errorf("fetch meal price: %w", err)
		}

		totalAmount = totalAmount.Add(price.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	orderQuery := `INSERT INTO orders (id, customer_id, vendor_id, status, total_amount, delivery_fee, delivery_address, created_at, updated_at)
	VALUES (@id, @customerId, @vendorId, @status, @totalAmount, @deliveryFee, @deliveryAddress, @createdAt, @updatedAt)`

	args := pgx.NamedArgs{
		"id":              orderID,
		"customerId":      customerID,
		"vendorId":        vendorID,
		"status":          "pending",
		"totalAmount":     totalAmount,
		"deliveryFee":     decimal.NewFromInt(constants.FAKE_DELIVERY), // For now
		"deliveryAddress": req.DeliveryAddress,
		"createdAt":       time.Now(),
		"updatedAt":       time.Now(),
	}

	_, err = tx.Exec(ctx, orderQuery, args)
	if err != nil {
		return "", fmt.Errorf("insert order: %w", err)
	}

	now := time.Now()

	mealIds := make([]int64, 0, len(req.Items))
	for _, v := range req.Items {
		mealIds = append(mealIds, v.MealID)
	}
	prices := make(map[int64]decimal.Decimal)
	rows, err := tx.Query(ctx, `SELECT id, price FROM meals WHERE id = ANY($1)`, mealIds)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var mealId int64 = 0
		price := decimal.Decimal{}
		err = rows.Scan(&mealId, &price)
		if err != nil {
			return "", fmt.Errorf("scan error: %v", err)
		}
		prices[mealId] = price
	}

	err = rows.Err()
	if err != nil {
		return "", fmt.Errorf("rows error: %v", err)
	}

	entries := [][]any{}
	tableName := "order_items"
	columns := []string{"id", "order_id", "meal_id", "quantity", "price", "created_at"}

	for _, item := range req.Items {
		price, ok := prices[item.MealID]
		if !ok {
			continue
			// collect missing IDs, email customer, return error, whatever your plan is
		}
		orderItemID := global.SnowFlake.Generate().Int64()
		entries = append(entries, []any{orderItemID, orderID, item.MealID, item.Quantity, price, now})
	}

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		columns,
		pgx.CopyFromRows(entries),
	)
	if err != nil {
		return "", fmt.Errorf("error copying into %s table: %w", tableName, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (o *ordersRepo) CancelOrder(ctx context.Context, customerID int64, orderID string) error {
	query := `UPDATE orders SET status = @status, updated_at = @updatedAt
	WHERE id = @orderId AND customer_id = @customerId AND status IN ('pending', 'confirmed')`

	args := pgx.NamedArgs{
		"status":     "cancelled",
		"updatedAt":  time.Now(),
		"orderId":    orderID,
		"customerId": customerID,
	}

	_, err := o.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("cancel order: %w", err)
	}

	return nil
}

func (o *ordersRepo) GetOrder(ctx context.Context, customerID int64, orderID string) (*model.Order, error) {
	order := &model.Order{}
	query := `SELECT id, customer_id, vendor_id, status, total_amount, delivery_fee, delivery_address, created_at, updated_at
	FROM orders WHERE id = $1 AND customer_id = $2 LIMIT 1`

	err := o.db.QueryRow(ctx, query, orderID, customerID).Scan(
		&order.ID,
		&order.CustomerID,
		&order.VendorID,
		&order.Status,
		&order.TotalAmount,
		&order.DeliveryFee,
		&order.DeliveryAddress,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("fetch order: %w", err)
	}

	return order, nil
}
