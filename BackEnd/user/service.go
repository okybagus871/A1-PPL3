package user

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
	
	user := User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Name: userRequest.Name,
	}
	newUser, err := s.repository.CreateUser(user)
	return newUser, err
}