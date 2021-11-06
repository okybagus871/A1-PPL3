package user

import (
	"context"
	"encoding/json"
	//"log"
	"net/http"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler{
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/signup").Handler(httptransport.NewServer(
		endpoints.SignUp,
		decodeSignUpReq,
		encodeResponse,
	))

	return r
}

func decodeSignUpReq(ctx context.Context, r *http.Request)(interface{}, error){
	var req SignUpReq
	err := json.NewDecoder(r.Body).Decode(&req.User)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}