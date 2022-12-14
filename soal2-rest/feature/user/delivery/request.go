package delivery

import "weddingguest_app/domain"

type InsertFormat struct {
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Role     string `json:"role" form:"role"`
	Photo    string `json:"photo" form:"photo"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		UserName: i.UserName,
		Email:    i.Email,
		Password: i.Password,
		FullName: i.FullName,
		Role:     i.Role,
		Photo:    i.Photo,
	}
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		Email:    lf.Email,
		Password: lf.Password,
	}
}
