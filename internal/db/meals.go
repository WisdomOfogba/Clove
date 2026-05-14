package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
	server "github.com/chibx/vendor-pulse/internal/server_errors"
	"github.com/chibx/vendor-pulse/internal/types"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"gorm.io/gorm"
)

func (m *mealsRepo) CreateMeal(ctx context.Context, meal *model.Meal, pictures []*model.MealPicture) (int64, error) {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	meal.ID = global.SnowFlake.Generate().Int64()
	meal.CreatedAt = time.Now()
	meal.UpdatedAt = time.Now()
	if meal.Status == "" {
		meal.Status = "active"
	}

	if err := tx.Create(meal).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("insert meal: %w", err)
	}

	for _, pic := range pictures {
		pic.ID = global.SnowFlake.Generate().Int64()
		pic.MealID = meal.ID
		pic.CreatedAt = time.Now()

		if err := tx.Create(pic).Error; err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("insert meal picture: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return meal.ID, nil
}

func (m *mealsRepo) GetMealPicturesForVendor(ctx context.Context, vendorID, mealID int64, pictureIDs []int64) ([]*model.MealPicture, error) {
	if len(pictureIDs) == 0 {
		return nil, fmt.Errorf("no media ids provided")
	}

	var pictures []*model.MealPicture
	if err := m.db.WithContext(ctx).
		Joins("JOIN meals ON meals.id = meal_pictures.meal_id").
		Where("meals.vendor_id = ? AND meal_pictures.meal_id = ? AND meal_pictures.id = ANY(?)", vendorID, mealID, pictureIDs).
		Find(&pictures).Error; err != nil {
		return nil, fmt.Errorf("query meal pictures: %w", err)
	}

	return pictures, nil
}

func (m *mealsRepo) DeleteMealPictures(ctx context.Context, vendorID, mealID int64, pictureIDs []int64) error {
	if len(pictureIDs) == 0 {
		return fmt.Errorf("no media ids provided")
	}

	result := m.db.WithContext(ctx).
		Joins("JOIN meals ON meals.id = meal_pictures.meal_id").
		Where("meals.vendor_id = ? AND meal_pictures.meal_id = ? AND meal_pictures.id = ANY(?)", vendorID, mealID, pictureIDs).
		Delete(&model.MealPicture{})

	if result.Error != nil {
		event := global.Logger.Err(result.Error)
		arr := event.CreateArray()
		for _, v := range pictureIDs {
			arr.Int64(v)
		}
		event.Array("picture_ids", arr).Int64("meal_id", mealID).Msg("Couldn't delete meal images")
		return fmt.Errorf("delete meal pictures: %w", result.Error)
	}

	return nil
}

func (m *mealsRepo) UpdateMeal(ctx context.Context, vendorID, mealID int64, update *request.UpdateMealRequest) error {
	result := m.db.WithContext(ctx).Model(&model.Meal{}).
		Where("id = ? AND vendor_id = ?", mealID, vendorID).
		Updates(map[string]interface{}{
			"name":        update.Name,
			"enabled":     update.Enabled,
			"description": update.Description,
			"price":       update.Price.IntPart(),
			"category":    update.Category,
			"status":      update.Status,
			"updated_at":  time.Now(),
		})

	if result.Error != nil {
		return fmt.Errorf("update meal: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *mealsRepo) DeleteMeal(ctx context.Context, vendorID, mealID int64) error {
	if err := m.db.WithContext(ctx).Where("id = ? AND vendor_id = ?", mealID, vendorID).Delete(&model.Meal{}).Error; err != nil {
		return fmt.Errorf("delete meal: %w", err)
	}

	return nil
}

func (m *mealsRepo) GetMealPicturesByMealID(ctx context.Context, vendorID, mealID int64) ([]*model.MealPicture, error) {
	var pictures []*model.MealPicture
	if err := m.db.WithContext(ctx).
		Joins("JOIN meals ON meals.id = meal_pictures.meal_id").
		Where("meals.vendor_id = ? AND meal_pictures.meal_id = ?", vendorID, mealID).
		Find(&pictures).Error; err != nil {
		return nil, fmt.Errorf("query meal pictures: %w", err)
	}

	return pictures, nil
}

func (m *mealsRepo) AddReview(ctx context.Context, mealId int64, review *model.Review) error {
	// Validate that the meal exists and vendor_id matches
	var count int64
	if err := m.db.WithContext(ctx).Model(&model.Meal{}).
		Where("id = ? AND vendor_id = ?", mealId, review.VendorID).
		Count(&count).Error; err != nil {
		return fmt.Errorf("validate meal and vendor: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("meal not found or vendor mismatch")
	}

	now := time.Now()
	review.ID = global.SnowFlake.Generate().Int64()
	review.MealID = mealId
	review.CreatedAt = now
	review.UpdatedAt = now

	if err := m.db.WithContext(ctx).Create(review).Error; err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}

func (m *mealsRepo) EditReview(ctx context.Context, reviewId int64, customerId int64, req *request.EditReviewRequest) error {
	var review model.Review
	if err := m.db.WithContext(ctx).Where("id = ? AND customer_id = ?", reviewId, customerId).First(&review).Error; err != nil {
		return err
	}

	if review.Edits >= 5 {
		return server.NewUserErr(server.MaxReviewReached, "Maximum of 5 review edits")
	}

	result := m.db.WithContext(ctx).Model(&model.Review{}).
		Where("id = ? AND customer_id = ?", reviewId, customerId).
		Updates(map[string]interface{}{
			"rating":     int64(req.Rating),
			"comment":    req.Comment,
			"edits":      gorm.Expr("edits + 1"),
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return fmt.Errorf("update review: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *mealsRepo) ListReviews(ctx context.Context, mealId int64, pagination types.Pagination) ([]*model.Review, error) {
	pagination.Normalize()
	var reviews []*model.Review

	if err := m.db.WithContext(ctx).
		Where("meal_id = ?", mealId).
		Limit(int(pagination.PageSize)).
		Offset(int(pagination.Page*pagination.PageSize)).
		Select("id", "customer_id", "rating", "comment", "edits", "created_at", "updated_at").
		Find(&reviews).Error; err != nil {
		global.Logger.Err(err).Msg("Couldn't load reviews from db")
		return nil, err
	}

	return reviews, nil
}
