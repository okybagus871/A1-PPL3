package user

import (
	//"fmt"
   	"context"
	"github.com/go-kit/kit/endpoint"
	//"net/http"
	//"encoding/json"
)

type Endpoints struct {
	SignUp endpoint.Endpoint
}

func MakeEndpoints(s UserService) Endpoints {
	return Endpoints {
		SignUp: makeSignUpEndpoint(s),
	}
}
func makeSignUpEndpoint (s UserService) endpoint.Endpoint {
   return func(ctx context.Context, request interface{}) (interface{}, error) {
	   req := request.(SignUpReq)

	   _, err := s.SignUp(ctx, req.User)
	   return DefaultResponse{Status:true, Message: "Success creating user"}, err
   }
}

