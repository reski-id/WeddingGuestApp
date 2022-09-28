package data

import (
	"fmt"
	"log"
	"weddingguest_app/domain"

	"gorm.io/gorm"
)

type noteData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.NoteData {
	return &noteData{
		db: db,
	}
}

func (nd *noteData) Insert(newData domain.Note) domain.Note {
	// fmt.Println("data :", newData)
	cnv := ToLocal(newData)
	// fmt.Println("cnv", cnv)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Note{}
	}
	return cnv.ToDomain()
}

func (bd *noteData) Update(dataID int, updatedData domain.Note) domain.Note {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Note{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *noteData) Delete(dataID int) bool {
	err := nd.db.Where("ID = ?", dataID).Delete(&Note{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (nd *noteData) GetAll() []domain.Note {
	var data []Note
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *noteData) GetNoteID(dataID int) []domain.Note {
	var data []Note
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
