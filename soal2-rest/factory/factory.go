package factory

import (
	ud "weddingguest_app/feature/user/data"
	userDelivery "weddingguest_app/feature/user/delivery"
	us "weddingguest_app/feature/user/usecase"

	nd "weddingguest_app/feature/note/data"
	noteDelivery "weddingguest_app/feature/note/delivery"
	nu "weddingguest_app/feature/note/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	noteData := nd.New(db)
	noteCase := nu.New(noteData)
	noteHandler := noteDelivery.New(noteCase)
	noteDelivery.RouteNote(e, noteHandler)

}
