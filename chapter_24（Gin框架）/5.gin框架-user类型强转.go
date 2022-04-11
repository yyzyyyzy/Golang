package main

type UserDto struct {
	Name  string `db:"name"`
	Phone string `db:"phone"`
}

func ToUserDto(user User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
