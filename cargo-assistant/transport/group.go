package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
	"errors"
	svc "marathon/cargo-assistant/service"
	ca "marathon/cargo-assistant"
)

func MakeHttpHandler(s svc.IGroupService, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	router = router.PathPrefix("/ca").Subrouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}
	fieldKeys := []string{"method"}
    s=ca.NewInstrumentingService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "ca", //不能用中横线,引发panic,命名合法性reg:[a-zA-Z_:][a-zA-Z0-9_:]*
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "ca",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of request in microseconds.",
		}, fieldKeys),
		s)

	getGroupEndpoint := ca.MakeGetGroupEndpoint(s)
	getGroupEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Nanosecond), 1))(getGroupEndpoint)
	getGroupEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getGroupEndpoint)

	//getOrderEndpoint :=MakeStartEndpoint(s)
	//getOrderEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getOrderEndpoint)
	//getOrderEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getOrderEndpoint)
	//addOrderEndpoint := MakeAddOrderEndpoint(s)
	//addOrderEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(addOrderEndpoint)
	//addOrderEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addOrderEndpoint)

	router.Methods("GET").
		Path("/group").
		Handler(kithttp.NewServer(
		getGroupEndpoint,
		decodeGetGroupRequest,
		encodeResponse,
		options...,
	))




/*	router.Methods("POST").
		Path("/order").
		Handler(kithttp.NewServer(
		addOrderEndpoint,
		decodeAddOrderRequest,
		encodeResponse,
		options...,
	))*/

	//router.Methods("POST").
	//	Path("/order").
	//	Handler(kithttp.NewServer(
	//	addOrderEndpoint,
	//	decodeAddOrderRequest,
	//	encodeResponse,
	//	options...,
	//))


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


func decodeGetGroupRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return ca.GetGroupRequest{}, nil
}

func decodeGetStartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil,errors.New("param err")
	}
	return ca.GetStartRequest{Id: id}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
