package user

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	//"github.com/google/uuid"
)

type UserService interface {
	//masukin modul
	SignUp(ctx context.Context, user userProfile) (*userProfile, error)
}

type userService struct{
	repo Repository
	logger log.Logger
}


func NewService(rep Repository, logger log.Logger) UserService {
	return &userService{
		repo: rep,
		logger: logger,
	}
}

var ErrEmpty = errors.New("empty string")

func (s *userService) SignUp(ctx context.Context, user userProfile) (*userProfile, error) {
 	logger := log.With(s.logger, "method", "CreateUser")

	// uuid, _ := uuid.NewUUID()
	// user.UserID = uuid
	user.Created_date = GetNow()
	
	//newUser := s.repo.CreateUser(ctx, user)

	if err := s.repo.CreateUser(ctx, user); err != nil {
		level.Error(s.logger).Log("err", err)
		return nil, err
	}
	logger.Log("create user")
	return &user, nil
}

func GetNow() time.Time {
	return time.Now().UTC()
}