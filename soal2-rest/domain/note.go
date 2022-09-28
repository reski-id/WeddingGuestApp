package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Note struct {
	ID        int
	Name      string
	Address   string
	Phone     string
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
}

type NoteHandler interface {
	InsertNote() echo.HandlerFunc
	UpdateNote() echo.HandlerFunc
	DeleteNote() echo.HandlerFunc
	GetAllNote() echo.HandlerFunc
	GetNoteID() echo.HandlerFunc
}

type NoteUseCase interface {
	AddNote(IDUser int, useNote Note) (Note, error)
	UpNote(IDNote int, updateData Note) (Note, error)
	DelNote(IDNote int) (bool, error)
	GetAllN() ([]Note, error)
	GetSpecificNote(NoteID int) ([]Note, error)
}

type NoteData interface {
	Insert(insertNote Note) Note
	Update(IDNote int, updatedNote Note) Note
	Delete(IDNote int) bool
	GetAll() []Note
	GetNoteID(NoteID int) []Note
}
