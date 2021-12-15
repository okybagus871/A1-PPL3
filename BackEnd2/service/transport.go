package service

import (
	"context"
	//"encoding/json"
	//"log"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/signup").Handler(httptransport.NewServer(
		endpoints.SignUp,
		decodeSignUpReq,
		encodeResponse,
	))
	
	r.Methods("GET").Path("/check-username").Handler(httptransport.NewServer(
		endpoints.CheckUsernameAvailability,
		decodeCheckUsernameReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/check-email").Handler(httptransport.NewServer(
		endpoints.CheckEmailAvailability,
		decodeCheckEmailReq,
		encodeResponse,
	))

	mailR := r.PathPrefix("/verify").Methods(http.MethodPost).Subrouter()
	mailR.Path("/mail").Handler(httptransport.NewServer(
		endpoints.ValidateEmail,
		decodeValidateEmailReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/get-user").Handler(httptransport.NewServer(
		endpoints.GetUserByEmail,
		decodeGetUserByEmailReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
