//go:build wireinject

package di

import (
	http "eatry/pkg/api"
	"eatry/pkg/api/handlers"
	config "eatry/pkg/config"
	db "eatry/pkg/db"
	repository "eatry/pkg/repository"
	usecase "eatry/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,

		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handlers.NewAdminHandler,

		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handlers.NewUserHandler,

		repository.NewCategoryRepository,
		usecase.NewCategoryUseCase,
		handlers.NewCategoryHandler,

		repository.NewProductRepository,
		usecase.NewProductUseCase,
		handlers.NewProductHandler,

		repository.NewCartRepository,
		usecase.NewCartUseCase,
		handlers.NewCartHandler,

		repository.NewCouponRepository,
		usecase.NewCouponUseCase,
		handlers.NewCouponHandler,

		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		handlers.NewOrderHandler,

		repository.NewPaymentRepository,
		usecase.NewPaymentUseCase,
		handlers.NewPaymentHandler,

		repository.NewWishlistRepository,
		usecase.NewWishlistUseCase,
		handlers.NewWishlistHandler,

		repository.NewWalletRepository,
		usecase.NewWalletUseCase,
		handlers.NewWalletHandler,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
