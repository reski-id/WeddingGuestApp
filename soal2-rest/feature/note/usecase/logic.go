package usecase

import (
	"errors"
	"fmt"
	"weddingguest_app/domain"
)

type noteUseCase struct {
	noteData domain.NoteData
}

func New(model domain.NoteData) domain.NoteUseCase {
	return &noteUseCase{
		noteData: model,
	}
}

func (nu *noteUseCase) AddNote(IDUser int, newNote domain.Note) (domain.Note, error) {
	if IDUser == -1 {
		return domain.Note{}, errors.New("invalid user")
	}

	newNote.UserID = IDUser
	fmt.Println("note", newNote)
	res := nu.noteData.Insert(newNote)

	if res.ID == 0 {
		return domain.Note{}, errors.New("error insert note")
	}
	return res, nil
}

func (nu *noteUseCase) UpNote(IDNote int, updateData domain.Note) (domain.Note, error) {
	if IDNote == -1 {
		return domain.Note{}, errors.New("invalid note")
	}

	// updateData.UserID = IDNote
	res := nu.noteData.Update(IDNote, updateData)
	if res.ID == 0 {
		return domain.Note{}, errors.New("error update note")
	}

	return res, nil
}

func (nu *noteUseCase) DelNote(IDNote int) (bool, error) {
	res := nu.noteData.Delete(IDNote)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (nu *noteUseCase) GetAllN() ([]domain.Note, error) {
	res := nu.noteData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *noteUseCase) GetSpecificNote(noteID int) ([]domain.Note, error) {
	res := nu.noteData.GetNoteID(noteID)
	if noteID == -1 {
		return nil, errors.New("error update note")
	}

	return res, nil
}
