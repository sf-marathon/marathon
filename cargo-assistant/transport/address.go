package transport

import (
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	svc "marathon/cargo-assistant/service"
	ca "marathon/cargo-assistant"
)

func MakeAddressHttpHandler(s svc.IAddressService,router *mux.Router, logger log.Logger) http.Handler {

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	getGroupEndpoint := ca.MakeAddressEndpoint(s)
	getGroupEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Nanosecond), 1))(getGroupEndpoint)
	getGroupEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getGroupEndpoint)

	router.Methods("GET").
		Path("/address").
		Handler(kithttp.NewServer(
		getGroupEndpoint,
		decodeGetGroupRequest,
		encodeResponse,
		options...,
	))
	return router
}


/*func decodeAddGroupRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req *dao.Group
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}*/

//func decodeAddOrderRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
//	var req *Order
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		return nil, err
//	}
//	return req, nil
//}

//
//func decodeGetGroupRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
//	return ca.GetGroupRequest{}, nil
//}
//
//func decodeGetStartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
//	vars := mux.Vars(r)
//	id, ok := vars["id"]
//	if !ok {
//		return nil,errors.New("param err")
//	}
//	return ca.GetStartRequest{Id: id}, nil
//}
//
//func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	return json.NewEncoder(w).Encode(response)
//}
