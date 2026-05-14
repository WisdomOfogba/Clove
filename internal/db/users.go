package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/dto"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
)

func (cust *usersRepo) CreateUser(ctx context.Context, user *model.User) error {
	user.ID = global.SnowFlake.Generate().Int64()
	user.Status = constants.USER_ACTIVE
	user.IsEmailVerified = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := cust.db.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}

	return nil
}

func (cust *usersRepo) GetUserLoginByEmail(ctx context.Context, email string) (*dto.UserLogin, error) {
	user := &dto.UserLogin{}

	if err := cust.db.WithContext(ctx).Select("id", "email", "password").Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (cust *usersRepo) CreateSession(ctx context.Context, session *model.UserSession) error {
	session.CreatedAt = time.Now()
	if err := cust.db.WithContext(ctx).Create(session).Error; err != nil {
		return err
	}
	return nil
}

func (cust *usersRepo) GetSessionByTokenHash(ctx context.Context, tokenHash string) (*model.UserSession, error) {
	userSession := &model.UserSession{}
	if err := cust.db.WithContext(ctx).Where("refresh_token_hash = ?", tokenHash).First(userSession).Error; err != nil {
		return nil, err
	}
	return userSession, nil
}

func (cust *usersRepo) DeleteSession(ctx context.Context, session *model.UserSession) error {
	query := cust.db.WithContext(ctx).Where("refresh_token_hash = ?", session.RefreshTokenHash)
	if session.UserId != 0 {
		query = query.Where("user_id = ?", session.UserId)
	}
	return query.Delete(&model.UserSession{}).Error
}
