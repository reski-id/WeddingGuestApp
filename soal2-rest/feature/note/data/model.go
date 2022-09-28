package data

import (
	"time"
	"weddingguest_app/domain"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Name    string
	Address string
	Phone   string
	Note    string
	UserID  int
	User    User `gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
}

func (b *Note) ToDomain() domain.Note {
	return domain.Note{
		ID:        int(b.ID),
		Name:      b.Name,
		Address:   b.Address,
		Phone:     b.Phone,
		Note:      b.Note,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: time.Time{},
		UserID:    int(b.UserID),
	}
}

func ParseToArr(arr []Note) []domain.Note {
	var res []domain.Note

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Note) Note {
	var res Note
	res.ID = uint(data.ID)
	res.UserID = int(data.UserID)
	res.Name = data.Name
	res.Phone = data.Phone
	res.Address = data.Address
	res.Note = data.Note
	return res
}
