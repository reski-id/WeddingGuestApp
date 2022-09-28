package delivery

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"weddingguest_app/domain"
	"weddingguest_app/feature/common"

	"github.com/labstack/echo/v4"
)

type noteHandler struct {
	noteUsecase domain.NoteUseCase
}

func New(nu domain.NoteUseCase) domain.NoteHandler {
	return &noteHandler{
		noteUsecase: nu,
	}
}

func (nh *noteHandler) InsertNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp NoteInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		fmt.Println(tmp)

		var userid = common.ExtractData(c)
		data, err := nh.noteUsecase.AddNote(common.ExtractData(c), tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(userid)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (nh *noteHandler) UpdateNote() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp NoteInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		data, err := nh.noteUsecase.UpNote(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"data":    data,
		})
	}
}

func (nh *noteHandler) DeleteNote() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := nh.noteUsecase.DelNote(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete note",
		})
	}
}

func (nh *noteHandler) GetAllNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.noteUsecase.GetAllN()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all note",
			"users":   data,
		})
	}
}

func (nh *noteHandler) GetNoteID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNote := c.Param("id")
		id, _ := strconv.Atoi(idNote)
		data, err := nh.noteUsecase.GetSpecificNote(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my note",
			"users":   data,
		})
	}
}
