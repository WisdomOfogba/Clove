package db

import (
	"github.com/chibx/vendor-pulse/internal/global"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	usersRepo    struct{ db *pgxpool.Pool }
	mealsRepo    struct{ db *pgxpool.Pool }
	ordersRepo   struct{ db *pgxpool.Pool }
	paymentsRepo struct{ db *pgxpool.Pool }
)

var (
	usersR    *usersRepo
	mealsR    *mealsRepo
	ordersR   *ordersRepo
	paymentsR *paymentsRepo
)

func Users() *usersRepo {
	if usersR == nil {
		usersR = &usersRepo{global.DB}
	}
	return usersR
}

func Meals() *mealsRepo {
	if mealsR == nil {
		mealsR = &mealsRepo{global.DB}
	}
	return mealsR
}

func Orders() *ordersRepo {
	if ordersR == nil {
		ordersR = &ordersRepo{global.DB}
	}
	return ordersR
}

func Payments() *paymentsRepo {
	if paymentsR == nil {
		paymentsR = &paymentsRepo{global.DB}
	}
	return paymentsR
}
