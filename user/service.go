package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegistUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository repository) *service {
	return &service{&repository}
}

func (s *service) RegisterUser(input RegistUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, nil
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, nil
	}
	return newUser, nil
}
