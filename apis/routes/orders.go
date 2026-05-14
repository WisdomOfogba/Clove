package routes

import (
	"github.com/chibx/vendor-pulse/internal/constants"
	"github.com/chibx/vendor-pulse/internal/db"
	server "github.com/chibx/vendor-pulse/internal/server_errors"
	"github.com/chibx/vendor-pulse/internal/types/request"
	"github.com/chibx/vendor-pulse/internal/types/response"
	"github.com/gofiber/fiber/v3"
)

func PlaceOrder() fiber.Handler {
	err500 := fiber.NewError(fiber.StatusInternalServerError, "Couldn't place order, please try again")
	return func(ctx fiber.Ctx) error {
		reqBody := &request.PlaceOrderRequest{}

		err := ctx.Bind().Body(reqBody)
		if err != nil {
			errorBags := server.ValErrToBag(err)
			return response.FromFiberError(ctx, fiber.ErrBadRequest, errorBags)
		}

		customerID := int64(0)
		if userCtx, ok := ctx.Locals(constants.UserCtxKey).(request.UserCtx); ok {
			customerID = userCtx.ID
		}
		if customerID == 0 {
			return response.FromFiberError(ctx, fiber.ErrUnauthorized)
		}

		orderID, err := db.Orders().CreateOrder(ctx.Context(), customerID, reqBody.VendorID, reqBody)
		if err != nil {
			return response.FromFiberError(ctx, err500)
		}

		return response.WriteResponse(ctx, fiber.StatusCreated, "Order placed successfully", fiber.Map{"order_id": orderID})
	}
}

func CancelOrder() fiber.Handler {
	err500 := fiber.NewError(fiber.StatusInternalServerError, "Couldn't cancel order, please try again")
	return func(ctx fiber.Ctx) error {
		reqBody := &request.CancelOrder{}

		err := ctx.Bind().Body(reqBody)
		if err != nil {
			errorBags := server.ValErrToBag(err)
			return response.FromFiberError(ctx, fiber.ErrBadRequest, errorBags)
		}

		customerID := int64(0)
		if userCtx, ok := ctx.Locals(constants.UserCtxKey).(request.UserCtx); ok {
			customerID = userCtx.ID
		}
		if customerID == 0 {
			return response.FromFiberError(ctx, fiber.ErrUnauthorized)
		}

		err = db.Orders().CancelOrder(ctx.Context(), customerID, reqBody.OrderID)
		if err != nil {
			return response.FromFiberError(ctx, err500)
		}

		return response.WriteResponse(ctx, fiber.StatusOK, "Order cancelled successfully")
	}
}
