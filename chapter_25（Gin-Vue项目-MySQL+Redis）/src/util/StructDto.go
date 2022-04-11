package util

import "model"

type UserDto struct {
	UserName string `db:"username"`
	Phone    string `db:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		UserName: user.UserName,
		Phone:    user.Phone,
	}
}
