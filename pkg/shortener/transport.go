package shortener

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeFindEndpoint(svc RedirectService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(findRequest)
		url, err := svc.Find(req.Code)
		if err != nil {
			return findResponse{URL: url, Err: err.Error()}, nil
		}
		return findResponse{URL: url, Err: ""}, nil
	}
}

func MakeStoreEndpoint(svc RedirectService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(storeRequest)
		code, err := svc.Store(req.URL)
		if err != nil {
			return storeResponse{Code: code, Err: err.Error()}, nil
		}
		return storeResponse{Code: code, Err: ""}, nil
	}
}

func DecodeFindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request findRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeStoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request storeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type storeRequest struct {
	URL string `json:"url"`
}

type storeResponse struct {
	Code string `json:"code"`
	Err  string `json:"err,omitempty"`
}

type findRequest struct {
	Code string `json:"code"`
}

type findResponse struct {
	URL string `json:"url"`
	Err string `json:"err,omitempty"`
}
