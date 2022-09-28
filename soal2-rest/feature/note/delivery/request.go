package delivery

import (
	"time"
	"weddingguest_app/domain"
)

type NoteInsertRequest struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
	Note    string `json:"note" form:"note"`
}

func (ni *NoteInsertRequest) ToDomain() domain.Note {
	return domain.Note{
		Name:      ni.Name,
		Address:   ni.Address,
		Phone:     ni.Phone,
		Note:      ni.Note,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
