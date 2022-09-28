package delivery

import "weddingguest_app/domain"

type NoteResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Note    string `json:"note"`
	UserID  int    `json:"user_id"`
}

func FromDomain(data domain.Note) NoteResponse {
	var res NoteResponse
	res.ID = int(data.ID)
	res.UserID = int(data.UserID)
	res.Name = data.Name
	res.Address = data.Address
	res.Phone = data.Phone
	res.Note = data.Note
	return res
}
