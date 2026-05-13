package db

import (
	"context"
	"fmt"
	"time"

	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/dto"
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/chibx/vendor-pulse/internal/model"
	"github.com/jackc/pgx/v5"
)

func (cust *usersRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, full_name, email, phone_number, password, status, is_email_verified, created_at, updated_at) 
	VALUES (@id, @fullName, @email, @phoneNumber, @password, @status, @emailVerified, @createdAt, @updatedAt);`

	args := pgx.NamedArgs{
		"id":            global.SnowFlake.Generate().Int64(),
		"fullName":      user.FullName,
		"email":         user.Email,
		"phoneNumber":   user.PhoneNumber,
		"password":      user.Password,
		"status":        constants.USER_ACTIVE,
		"emailVerified": true,
		"createdAt":     time.Now(),
		"updatedAt":     time.Now(),
	}
	_, err := cust.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (cust *usersRepo) GetUserLoginByEmail(ctx context.Context, email string) (*dto.UserLogin, error) {
	user := &dto.UserLogin{}

	query := `SELECT id, email, password FROM users WHERE email = $1 LIMIT 1;`
	err := cust.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (cust *usersRepo) CreateSession(ctx context.Context, session *model.UserSession) error {
	query := `INSERT INTO user_sessions (user_id, refresh_token_hash, last_ip, device_id, user_agent, created_at, expires_at) 
	VALUES (@userId, @tokenHash, @lastIP, @deviceId, @userAgent, createdAt, expiresAt)`
	args := pgx.NamedArgs{
		"userId":    session.UserId,
		"tokenHash": session.RefreshTokenHash,
		"lastIP":    session.LastIP,
		"deviceId":  session.DeviceId,
		"userAgent": session.UserAgent,
		"createdAt": time.Now(),
		"expiresAt": session.ExpiresAt,
	}
	_, err := cust.db.Exec(ctx, query, args)
	return err
}

func (cust *usersRepo) GetSessionByTokenHash(ctx context.Context, tokenHash string) (*model.UserSession, error) {
	userSession := &model.UserSession{}
	query := `SELECT user_id, refresh_token_hash, last_ip, device_id, user_agent, created_at, expires_at FROM user_sessions 
	WHERE refresh_token_hash = $1 LIMIT 1`
	err := cust.db.QueryRow(ctx, query, tokenHash).Scan(
		&userSession.UserId,
		&userSession.RefreshTokenHash,
		&userSession.LastIP,
		&userSession.DeviceId,
		&userSession.UserAgent,
		&userSession.CreatedAt,
		&userSession.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}

func (cust *usersRepo) DeleteSession(ctx context.Context, session *model.UserSession) error {
	query := `DELETE FROM user_sessions WHERE refresh_token_hash = $1`
	_, err := cust.db.Exec(ctx, query, session.RefreshTokenHash)
	return err
}
