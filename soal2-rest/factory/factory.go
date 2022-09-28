package factory

import (
	ud "weddingguest_app/feature/user/data"
	userDelivery "weddingguest_app/feature/user/delivery"
	us "weddingguest_app/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)
}
