package service

import (
	"database/sql"
	"context"
	"errors"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"BackEnd/datastruct"
	//"github.com/google/uuid"
)

type UserService interface {
	//masukin modul
	SignUp(ctx context.Context, user datastruct.User) (*datastruct.User, error)
	CheckUsernameAvailability(ctx context.Context, username string) (bool, error)
	CheckEmailAvailability(ctx context.Context, email string)(bool, error)
	ValidateEmail(ctx context.Context, email string, otp uint32) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error)
	GetUserPassword(ctx context.Context, email string)(string, error)
	UpdatePassword(ctx context.Context, email string, password string) error
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

func (s *userService) SignUp(ctx context.Context, user datastruct.User) (*datastruct.User, error) {
 	logger := log.With(s.logger, "method", "CreateUser")

	// uuid, _ := uuid.NewUUID()
	// user.UserID = uuid
	user.Created_date = GetNow()

	if err := s.repo.SignUp(ctx, user); err != nil {
		level.Error(s.logger).Log("err", err)
		return nil, err
	}
	logger.Log("create user")
	return &user, nil
}

func (s *userService) CheckUsernameAvailability(ctx context.Context, username string) (bool, error){
	logger := log.With(s.logger, "method", "CheckUsername")

	isExist, err := s.repo.CheckUsername(ctx, username)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return false, err
	}

	if isExist && err == nil {
		return false, err
	}

	logger.Log("username checked")
	return true, nil
}

func (s *userService) CheckEmailAvailability(ctx context.Context, email string)(bool, error){
	logger := log.With(s.logger, "method", "CheckEmail")

	isExist, err := s.repo.CheckEmail(ctx, email)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return false, err
	}

	if isExist && err == nil {
		return false, err
	}

	logger.Log("email checked")
	return true, nil
}

func (s *userService) ValidateEmail(ctx context.Context, email string, otp uint32) (bool, error) {
	
	logger := log.With(s.logger, "method", "ValidateEmail")

	code, err := s.repo.ValidateEmail(ctx, email, otp)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return false, err
	}

	if code != otp {
		return false, err
	}
	
	err2 := s.repo.UpdateEmailVerified(ctx, email) 
	if err2 != nil {
		return false, err
	}

	logger.Log("email validated")
	return true, nil
}

func(s *userService) GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error){
	var user *datastruct.User
	var err error

	logger := log.With(s.logger, "method", "GetUserByEmail")
	user, err = s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		level.Error(s.logger).Log("err", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		level.Error(s.logger).Log("err", err)
		return nil, err
	}

	if user.Email_verified == false {
		return nil, err
	}

	logger.Log("email checked")
	return user, nil
}

func(s *userService) GetUserPassword(ctx context.Context, email string)(string, error){
	logger := log.With(s.logger, "method", "GetUserPassword")

	password, err := s.repo.GetUserPassword(ctx, email)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return " ", err
	}

	if err == sql.ErrNoRows {
		level.Error(s.logger).Log("err", err)
		return " ", err
	}

	logger.Log("GetUserPassword done")
	return password, nil
}

func(s *userService) UpdatePassword(ctx context.Context, email string, password string) error {
	logger := log.With(s.logger, "method", "UpdatePassword")

	err := s.repo.UpdatePassword(ctx, email, password)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return err
	}

	logger.Log("UpdatePassword done")
	return nil
}

func GetNow() time.Time {
	return time.Now().UTC()
}
