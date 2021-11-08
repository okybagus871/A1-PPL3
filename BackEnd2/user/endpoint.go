package user

import (
	//"fmt"
   	"context"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Endpoints struct {
	SignUp endpoint.Endpoint
	CheckUsernameAvailability endpoint.Endpoint
}

func MakeEndpoints(s UserService) Endpoints {
	return Endpoints {
		SignUp: makeSignUpEndpoint(s),
		CheckUsernameAvailability: makeCheckUsernameEndpoint(s),
	}
}
func makeSignUpEndpoint (s UserService) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	   req := request.(SignUpReq)

	   _, err := s.SignUp(ctx, req.UserReq)
	   if err != nil {
		   return DefaultResponse{Status:true, Message: "Some msg error"}, err
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

		_, err := s.CheckUsernameAvailability(ctx,req.Username)
		if err != nil {
			return DefaultResponse{Status:false, Message: "User has already been taken"}, err
		}
		return DefaultResponse{Status:true, Message: "User available"}, err
	}
}

func decodeCheckUsernameReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req CheckUsernameReq
	vars := mux.Vars(r)

	req = CheckUsernameReq{
		Username: vars["username"],
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

