package user

import "strings"

type Service interface {
	GetUser() ([]User, error)
	CreateUser(userRequest UserRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUser() ([]User, error) {
	return s.repository.GetUser()
}

func (s *service) CreateUser(userRequest UserRequest) (User, error) {
	
	var first_name, last_name string
	nameParts := splitName(userRequest.Firstname)
	if len(nameParts) == 0 {
		first_name = nameParts[0]
	} else if len(nameParts) > 0 {
		var nm string
		for i := 0; i < len(nameParts) - 1; i++ {
			n := []string {nm, nameParts[i]}
			nm = strings.Join(n," ")
		}
		first_name = nm
		last_name = nameParts[len(nameParts)-1]
	}
	user := User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Firstname: first_name,
		Lastname: last_name,
	}
	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}