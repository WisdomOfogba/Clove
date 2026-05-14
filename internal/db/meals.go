package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"github.com/jackc/pgx/v5"
)

func (m *mealsRepo) CreateMeal(ctx context.Context, meal *model.Meal, pictures []*model.MealPicture) (int64, error) {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	meal.ID = global.SnowFlake.Generate().Int64()
	meal.CreatedAt = time.Now()
	meal.UpdatedAt = time.Now()
	if meal.Status == "" {
		meal.Status = "active"
	}

	query := `INSERT INTO meals (id, vendor_id, name, description, price, category, status, score, created_at, updated_at)
	VALUES (@id, @vendorId, @name, @description, @price, @category, @status, @score, @createdAt, @updatedAt)`
	args := pgx.NamedArgs{
		"id":          meal.ID,
		"vendorId":    meal.VendorID,
		"name":        meal.Name,
		"description": meal.Description,
		"price":       meal.Price,
		"category":    meal.Category,
		"status":      meal.Status,
		"score":       meal.Score,
		"createdAt":   meal.CreatedAt,
		"updatedAt":   meal.UpdatedAt,
	}

	_, err = tx.Exec(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("insert meal: %w", err)
	}

	for _, pic := range pictures {
		pic.ID = global.SnowFlake.Generate().Int64()
		pic.MealID = meal.ID
		pic.CreatedAt = time.Now()

		pictureQuery := `INSERT INTO meal_pictures (id, meal_id, image_url, public_id, is_primary, created_at)
		VALUES (@id, @mealId, @imageUrl, @publicId, @isPrimary, @createdAt)`
		pictureArgs := pgx.NamedArgs{
			"id":        pic.ID,
			"mealId":    pic.MealID,
			"imageUrl":  pic.ImageURL,
			"publicId":  pic.PublicID,
			"isPrimary": pic.IsPrimary,
			"createdAt": pic.CreatedAt,
		}

		_, err = tx.Exec(ctx, pictureQuery, pictureArgs)
		if err != nil {
			return 0, fmt.Errorf("insert meal picture: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return meal.ID, nil
}

func (m *mealsRepo) GetMealPicturesForVendor(ctx context.Context, vendorID, mealID int64, pictureIDs []int64) ([]*model.MealPicture, error) {
	if len(pictureIDs) == 0 {
		return nil, fmt.Errorf("no media ids provided")
	}

	query := `SELECT mp.id, mp.meal_id, mp.image_url, mp.public_id, mp.is_primary, mp.created_at
	FROM meal_pictures mp
	JOIN meals m ON m.id = mp.meal_id
	WHERE m.vendor_id = $1 AND mp.meal_id = $2 AND mp.id = ANY($3)`

	rows, err := m.db.Query(ctx, query, vendorID, mealID, pictureIDs)
	if err != nil {
		return nil, fmt.Errorf("query meal pictures: %w", err)
	}
	defer rows.Close()

	pictures := make([]*model.MealPicture, 0)
	for rows.Next() {
		pic := &model.MealPicture{}
		if err := rows.Scan(&pic.ID, &pic.MealID, &pic.ImageURL, &pic.PublicID, &pic.IsPrimary, &pic.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan meal picture: %w", err)
		}
		pictures = append(pictures, pic)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return pictures, nil
}

func (m *mealsRepo) DeleteMealPictures(ctx context.Context, vendorID, mealID int64, pictureIDs []int64) error {
	if len(pictureIDs) == 0 {
		return fmt.Errorf("no media ids provided")
	}

	// tx, err := m.db.Begin(ctx)
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	if err != nil {
	// 		_ = tx.Rollback(ctx)
	// 	}
	// }()

	query := `DELETE FROM meal_pictures
	USING meals
	WHERE meal_pictures.meal_id = meals.id
	AND meals.vendor_id = @vendorId
	AND meal_pictures.meal_id = @mealId
	AND meal_pictures.id = ANY(@mediaIds)`

	args := pgx.NamedArgs{
		"vendorId": vendorID,
		"mealId":   mealID,
		"mediaIds": pictureIDs,
	}

	// result, err := tx.Exec(ctx, query, args)
	_, err := m.db.Exec(ctx, query, args)
	if err != nil {
		// arr := zerolog.Arr()
		event := global.Logger.Err(err)
		arr := event.CreateArray()
		for _, v := range pictureIDs {
			arr.Int64(v)
		}
		event.Array("picture_ids", arr).Int64("meal_id", mealID)
		return fmt.Errorf("delete meal pictures: %w", err)
	}

	// err = tx.Commit(ctx)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (m *mealsRepo) UpdateMeal(ctx context.Context, vendorID, mealID int64, update *request.UpdateMealRequest) error {
	query := `UPDATE meals SET name = @name, description = @description, price = @price, category = @category, status = @status, updated_at = @updatedAt
	WHERE id = @mealId AND vendor_id = @vendorId`

	args := pgx.NamedArgs{
		"name":        update.Name,
		"description": update.Description,
		"price":       update.Price.IntPart(),
		"category":    update.Category,
		"status":      update.Status,
		"updatedAt":   time.Now(),
		"mealId":      mealID,
		"vendorId":    vendorID,
	}

	result, err := m.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("update meal: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (m *mealsRepo) DeleteMeal(ctx context.Context, vendorID, mealID int64) error {
	query := `DELETE FROM meals WHERE id = @mealId AND vendor_id = @vendorId`

	args := pgx.NamedArgs{
		"mealId":   mealID,
		"vendorId": vendorID,
	}

	_, err := m.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("delete meal: %w", err)
	}

	return nil
}

func (m *mealsRepo) GetMealPicturesByMealID(ctx context.Context, vendorID, mealID int64) ([]*model.MealPicture, error) {
	query := `SELECT mp.id, mp.meal_id, mp.image_url, mp.public_id, mp.is_primary, mp.created_at
	FROM meal_pictures mp
	JOIN meals m ON m.id = mp.meal_id
	WHERE m.vendor_id = $1 AND mp.meal_id = $2`

	rows, err := m.db.Query(ctx, query, vendorID, mealID)
	if err != nil {
		return nil, fmt.Errorf("query meal pictures: %w", err)
	}
	defer rows.Close()

	pictures := make([]*model.MealPicture, 0)
	for rows.Next() {
		pic := &model.MealPicture{}
		if err := rows.Scan(&pic.ID, &pic.MealID, &pic.ImageURL, &pic.PublicID, &pic.IsPrimary, &pic.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan meal picture: %w", err)
		}
		pictures = append(pictures, pic)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return pictures, nil
}
