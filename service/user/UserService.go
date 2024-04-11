package user

import (
	"ginEssential/dao/user"
	"ginEssential/dto"
)

type Service struct {
	userDAO user.DAO
}

func (s *Service) GetAllUsers() ([]dto.UserDto, error) { //model.User类型的切片
	users, err := s.userDAO.GetAllUsers()
	if err != nil {
		return nil, err
	}
	userDtos := make([]dto.UserDto, len(users))
	for i, user := range users {
		userDtos[i] = dto.ToUserDto(&user)
	}
	return userDtos, nil
}
