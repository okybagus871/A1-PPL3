package service

import (
	"fmt"
   	"context"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	//"BackEnd/datastruct"
)

type Endpoints struct {
	SignUp endpoint.Endpoint
	CheckUsernameAvailability endpoint.Endpoint
	CheckEmailAvailability endpoint.Endpoint
	ValidateEmail endpoint.Endpoint
}

func MakeEndpoints(s UserService) Endpoints {
	return Endpoints {
		SignUp: makeSignUpEndpoint(s),
		CheckUsernameAvailability: makeCheckUsernameEndpoint(s),
		CheckEmailAvailability: makeCheckEmailEndpoint(s),
		ValidateEmail: makeValidateEmailEndpoint(s),
	}
}

func makeSignUpEndpoint (s UserService) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	   req := request.(SignUpReq)

	   _, err := s.SignUp(ctx, req.UserReq)
	   if err != nil {
		   return DefaultResponse{Status:false, Message: "Some msg error"}, err
	   }
	   return DefaultResponse{Status:true, Message: "Success creating user"}, err
   }
}

func decodeSignUpReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req.UserReq)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func makeCheckUsernameEndpoint (s UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(CheckUsernameReq)

		ret, err := s.CheckUsernameAvailability(ctx,req.Username)
		if err != nil {
			return DefaultResponse{Status:false, Message: "Username error"}, err
		} else if ret == false &&  err == nil {
			return DefaultResponse{Status:false, Message: "User has already been taken"}, err
		}
		return DefaultResponse{Status:true, Message: "User available"}, err
	}
}

func decodeCheckUsernameReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req CheckUsernameReq
	params := mux.Vars(r)
	username := params["username"]

	req.Username = username
	fmt.Println(username)
	return req, nil
}

func makeCheckEmailEndpoint (s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(CheckEmailReq)

		ret, err := s.CheckEmailAvailability(ctx,req.Email)
		if err != nil {
			return DefaultResponse{Status:false, Message: "Email error"}, err
		} else if ret == false &&  err == nil {
			return DefaultResponse{Status:false, Message: "User has already been taken"}, err
		}
		return DefaultResponse{Status:true, Message: "User available"}, err
	}
}

func decodeCheckEmailReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req CheckEmailReq
	params := mux.Vars(r)
	email := params["email"]

	req.Email = email
	fmt.Println(email)
	return req, nil
}

func makeValidateEmailEndpoint (s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(ValidateEmailReq)

		ret, err := s.ValidateEmail(ctx,req.Email,req.OTP)
		if err != nil {
			return DefaultResponse{Status:false, Message: "OTP is not valid"}, err
		} else if ret == true && err == nil {
			return DefaultResponse{Status:true, Message: "Email verified"}, err
		}
		return DefaultResponse{Status:false, Message: "OTP is not valid"}, err
	}
}

func decodeValidateEmailReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req ValidateEmailReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

