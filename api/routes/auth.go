package routes

import (
	"crypto/subtle"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/chibx/vendor-pulse/internal/types/response"
	"github.com/chibx/vendor-pulse/internal/constants"
	server "github.com/chibx/vendor-pulse/internal/server_errors"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/api"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

func CustomerRegister() fiber.Handler {
	err500 := fiber.NewError(fiber.StatusInternalServerError, "Error occurred while creating your account, please try again")

	return func(ctx fiber.Ctx) error {
		var err error
		// var errorBag = []serverErrors.ErrorDetail{}
		userForRegister := &request.RegisterCustomerRequest{}

		err = ctx.Bind().Body(userForRegister)
		if err != nil {
			errorBag := server.ValErrToBag(err)

			logger.Error("Error occured while parsing login values", zap.Error(err))
			return response.FromFiberError(ctx, err500)
		}

		regTokenJWT := ctx.Cookies("reg_token")
		now := time.Now()
		if len(strings.TrimSpace(regTokenJWT)) == 0 {
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "Invalid Request!")
		}

		regToken, err := auth.ValidateRegToken(api, regTokenJWT, api.Config.SecretKey)
		if err != nil {
			serverErrors := new(serverErrors.ServerErr)
			if errors.As(err, &serverErrors) {
				return response.WriteResponse(ctx, fiber.StatusBadRequest, serverErrors.Message)
			}
			return response.FromFiberError(ctx, err500)
		}

		if now.After(regToken.ExpiresAt.Time) {
			logger.Warn("Reg Token: Used after jwt expiration")
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "Expired Registration Token!!")
		}

		err = utils.Validator().Struct(userForRegister)

		isFatal, errorBag := serverErrors.HandleValidationError(err)
		if isFatal {
			logger.Error("InvalidValidationError while registering a user", zap.Error(err))
			return response.WriteResponse(ctx, fiber.ErrBadRequest.Code, err500.Message)
		}
		if len(errorBag) > 0 {
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "One or more fields are invalid", errorBag)
		}

		tokenStruc, err := db.BackendUsers().GetRegToken(ctx.Context(), regTokenJWT)
		if err != nil {
			if errors.Is(err, serverErrors.ErrDBRecordNotFound) {
				return response.WriteResponse(ctx, fiber.StatusBadRequest, "Registration Token not found.")
			}

			return response.FromFiberError(ctx, err500)
		}

		if now.After(tokenStruc.ExpiryAt) {
			logger.Warn("Reg Token: Used after db expiration")
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "Expired Registration Token!!")
		}

		if subtle.ConstantTimeCompare([]byte(tokenStruc.Code), []byte(userForRegister.Code)) == 0 {
			logger.Warn("Invalid Code used")
			// TODO: Maybe add a counter that would delete the token and alert the app owner of a potential cyber attack
			return response.WriteResponse(ctx, fiber.StatusUnauthorized, "You cannot proceed")
		}

		if len(userForRegister.UserName) > constants.MaxUsernameLimit {
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "Please enter a username less than "+strconv.Itoa(constants.MaxUsernameLimit)+" characters")
		}

		if len(userForRegister.Password) > constants.MaxPasswordLimit {
			return response.WriteResponse(ctx, fiber.StatusBadRequest, "Please enter a password less than "+strconv.Itoa(constants.MaxPasswordLimit)+" characters")
		}

		user, err := userForRegister.ToDBBackendUser(ctx.Context(), api, ctx)
		if err != nil {
			serverErr := new(server.ServerErr)
			if errors.As(err, &serverErr) {
				logger.Error(serverErr.Message)
				return response.WriteResponse(ctx, serverErr.Code, serverErr.Message)
			}

			logger.Error("Error converting user to db backend user", zap.Error(err))
			return response.FromFiberError(ctx, err500)
		}

		err = db.BackendUsers().CreateUser(ctx.Context(), user)
		if err != nil {
			logger.Error("DB Error creating backend user", zap.Error(err))
			return response.FromFiberError(ctx, err500)
		}

		// TODO: Add an ip addition feature

		return response.WriteResponse(ctx, fiber.StatusOK, "User created successfully.")
	}
}

func CustomerLogin() fiber.Handler {
	err500 := fiber.NewError(fiber.StatusInternalServerError, "Error occurred while signing you in, please try again")

	return func(ctx fiber.Ctx) error {
		return nil
	}
}

func CustomerLogout() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return nil
	}
}
