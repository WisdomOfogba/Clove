package db

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
	server "github.com/chibx/vendor-pulse/internal/server_errors"
	"github.com/chibx/vendor-pulse/internal/types"
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
		event.Array("picture_ids", arr).Int64("meal_id", mealID).Msg("Couldn't delete meal images")
		return fmt.Errorf("delete meal pictures: %w", err)
	}

	// err = tx.Commit(ctx)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (m *mealsRepo) UpdateMeal(ctx context.Context, vendorID, mealID int64, update *request.UpdateMealRequest) error {
	query := `UPDATE meals SET name = @name, enabled = @enabled, description = @description, price = @price, category = @category, status = @status, updated_at = @updatedAt
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
		"enabled":     update.Enabled,
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

func (m *mealsRepo) AddReview(ctx context.Context, mealId int64, review *model.Review) error {
	// Validate that the meal exists and vendor_id matches
	var exists bool
	err := m.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM meals WHERE id = $1 AND vendor_id = $2)", mealId, review.VendorID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("validate meal and vendor: %w", err)
	}
	if !exists {
		return fmt.Errorf("meal not found or vendor mismatch")
	}

	now := time.Now()
	review.ID = global.SnowFlake.Generate().Int64()
	review.CreatedAt = now
	review.UpdatedAt = now

	query := `INSERT INTO reviews (id, customer_id, vendor_id, meal_id, rating, comment, edits, sentiment, sentiment_score, created_at, updated_at)
	VALUES (@id, @customerId, @vendorId, @mealId, @rating, @comment, @edits, @sentiment, @sentimentScore, @createdAt, @updatedAt)`

	args := pgx.NamedArgs{
		"id":             review.ID,
		"customerId":     review.CustomerID,
		"vendorId":       review.VendorID,
		"mealId":         review.MealID,
		"rating":         review.Rating,
		"comment":        review.Comment,
		"edits":          review.Edits,
		"sentiment":      review.Sentiment,
		"sentimentScore": review.SentimentScore,
		"createdAt":      review.CreatedAt,
		"updatedAt":      review.UpdatedAt,
	}

	_, err = m.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}

func (m *mealsRepo) EditReview(ctx context.Context, reviewId int64, customerId int64, req *request.EditReviewRequest) error {
	edits := 0

	err := m.db.QueryRow(
		ctx,
		`SELECT edits FROM reviews WHERE id = $1 AND customer_id = $2 LIMIT 1`,
		reviewId, customerId,
	).Scan(&edits)
	if err != nil {
		return err
	}

	if edits >= 5 {
		return server.NewUserErr(server.MaxReviewReached, "Maximum of 5 review edits")
	}

	query := `UPDATE reviews SET rating = @rating, comment = @comment, edits = edits + 1, updated_at = @updatedAt
	WHERE id = @reviewId AND customer_id = @customerId`

	args := pgx.NamedArgs{
		"rating":     int64(req.Rating),
		"comment":    req.Comment,
		"updatedAt":  time.Now(),
		"reviewId":   reviewId,
		"customerId": customerId,
	}

	result, err := m.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("update review: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (m *mealsRepo) ListReviews(ctx context.Context, mealId int64, pagination types.Pagination) ([]*model.Review, error) {
	pagination.Normalize()
	reviews := make([]*model.Review, 0, 10)

	query := `SELECT id, customer_id, rating, comment, edits, created_at, updated_at 
	WHERE meal_id = @mealID
	FROM reviews LIMIT @limit OFFSET @offset`

	args := pgx.NamedArgs{
		"mealId": mealId,
		"limit":  pagination.PageSize,
		"offset": pagination.Page * pagination.Page,
	}

	rows, err := m.db.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		review := &model.Review{}
		err = rows.Scan(
			&review.ID, &review.CustomerID, &review.Rating,
			&review.Comment, &review.Edits, &review.CreatedAt, &review.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		if cap(reviews) <= len(reviews)+1 {
			reviews = slices.Grow(reviews, int(float64(cap(reviews))*1.8))
		}
		reviews = append(reviews, review)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return reviews, nil
}
