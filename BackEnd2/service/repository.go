package service

import (
	// "strconv"
	// "math/rand"
	// "time"
	"BackEnd/datastruct"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "SadhelEX <sadhelex@gmail.com>"
const CONFIG_AUTH_EMAIL = "oky.bagus.tif20@polban.ac.id"
const CONFIG_AUTH_PASSWORD = "sayacakep87"

var RepoErr = errors.New("Unable to handle Repo Request")

type Repository interface {
	SignUp(ctx context.Context, user datastruct.User) error
	CheckUsername(ctx context.Context, username string) (bool, error)
	CheckEmail(ctx context.Context, email string) (bool, error)
	ValidateEmail(ctx context.Context, email string, otp uint32) (uint32, error)
	UpdateEmailVerified(ctx context.Context, email string) error
	GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error)
	GetUserPassword(ctx context.Context, email string) (string, error)
	UpdatePassword(ctx context.Context, email string, password string) error
}

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}
}

func (r *repo) SignUp(ctx context.Context, user datastruct.User) error {
	token_hash := GenerateRandomString(15)
	user.Token_hash = token_hash

	min := 1000
	max := 9999
	rand.Seed(time.Now().UnixNano())
	rd := rand.Intn(max-min+1) + min
	s := strconv.Itoa(rd)
	to := []string{user.Email}
	cc := []string{user.Email}
	subject := "Email Verification"
	message := "You can also enter this verification code : " + s

	sendMail(to, cc, subject, message)
	user.OTP = rd
	hashPass, eror := PasswordHashing(user.Password)
	if eror != nil {
		return nil
	}
	user.Password = hashPass
	user.Email_verified = false
	user.Phonenumber = " "
	user.Identity_type = " "
	user.Identity_no = " "
	user.Address_ktp = " "
	user.Postal_code = " "
	user.Emergency_call = " "

	sql := `INSERT INTO users (username, name, password, created_date, email, token_hash, otp, email_verified,
				phonenumber, identity_type, identity_no, address_ktp, postal_code, emergency_call)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	if user.Name == "" || user.Password == "" || user.Username == "" || user.Email == "" {
		return RepoErr
	}
	_, err := r.db.ExecContext(
		ctx,
		sql,
		user.Username,
		user.Name,
		user.Password,
		user.Created_date,
		user.Email,
		user.Token_hash,
		user.OTP,
		user.Email_verified,
		user.Phonenumber,
		user.Identity_type,
		user.Identity_no,
		user.Address_ktp,
		user.Postal_code,
		user.Emergency_call)

	if err != nil {
		return err
	}

	return nil
}

func PasswordHashing(raw string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}

func (r *repo) CheckUsername(ctx context.Context, username string) (bool, error) {
	var exists bool

	level.Debug(r.logger).Log("msg", "start run CheckUsername")
	sql := `SELECT EXISTS(SELECT * FROM users WHERE username=$1);`

	err := r.db.QueryRow(sql, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	level.Debug(r.logger).Log("msg", "finish CheckUsername")
	return exists, nil
}

func (r *repo) CheckEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	level.Debug(r.logger).Log("msg", "start run CheckEmail")
	sql := `SELECT EXISTS(SELECT * FROM users WHERE email = $1);`

	err := r.db.QueryRow(sql, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	level.Debug(r.logger).Log("msg", "finish CheckEmail")
	return exists, nil
}

func (r *repo) ValidateEmail(ctx context.Context, email string, otp uint32) (uint32, error) {
	var otp_sys uint32
	level.Debug(r.logger).Log("msg", "start run ValidateEmail")

	sql := `SELECT otp FROM users WHERE email = $1 LIMIT 1;`
	err := r.db.QueryRow(sql, email).Scan(&otp_sys)

	if err != nil {
		return 0, err
	}

	level.Debug(r.logger).Log("msg", "finish ValidateEmail")
	return otp_sys, nil
}

func (r *repo) UpdateEmailVerified(ctx context.Context, email string) error {
	level.Debug(r.logger).Log("msg", "start run UpdateEmailVerified")

	email_verified := true
	sql := `UPDATE users SET email_verified = $1 WHERE email = $2;`
	_, err := r.db.ExecContext(ctx, sql, email_verified, email)

	if err != nil {
		return err
	}

	level.Debug(r.logger).Log("msg", "finish UpdateEmailVerified")
	return nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error) {
	level.Debug(r.logger).Log("msg", "start run GetUserByEmail")

	var user datastruct.User
	sql := `SELECT username, password, email, name, created_date, email_verified,
			phonenumber, identity_type, identity_no, address_ktp, postal_code, emergencycall
			FROM users WHERE email = $1;`
	err := r.db.QueryRowContext(ctx, sql, email).Scan(
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Created_date,
		&user.Email_verified,
		&user.Phonenumber,
		&user.Identity_type,
		&user.Identity_no,
		&user.Address_ktp,
		&user.Postal_code,
		&user.Emergency_call,
	)

	if err != nil {
		return nil, err
	}

	level.Debug(r.logger).Log("msg", "finish GetUserByEmail")

	return &user, nil
}

func (r *repo) GetUserPassword(ctx context.Context, email string) (string, error) {
	level.Debug(r.logger).Log("msg", "start run GetUserPassword")

	var password string

	sql := `SELECT password FROM users WHERE email = $1;`
	err := r.db.QueryRow(sql, email).Scan(&password)

	if err != nil {
		return " ", err
	}

	level.Debug(r.logger).Log("msg", "finish GetUserPassword")

	return password, nil
}

func (r *repo) UpdatePassword(ctx context.Context, email string, password string) error {
	level.Debug(r.logger).Log("msg", "start run UpdatePassword")

	sql := `UPDATE users SET password = $1 WHERE email = $2;`
	_, err := r.db.ExecContext(ctx, sql, password, email)

	if err != nil {
		return err
	}

	level.Debug(r.logger).Log("msg", "finish UpdatePassword")
	return nil
}
