package delivery

import (
	"weddingguest_app/config"
	"weddingguest_app/domain"
	"weddingguest_app/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteNote(e *echo.Echo, bc domain.NoteHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/note", bc.InsertNote(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/note/:id", bc.UpdateNote(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/note/:id", bc.DeleteNote(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/note", bc.GetAllNote())
	e.GET("/note/:id", bc.GetNoteID())
}
