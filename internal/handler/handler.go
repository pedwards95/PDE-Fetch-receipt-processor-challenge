package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/errorhandler"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/logger"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/points"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/receipts"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type ctxKey string

// ...
const (
	REQUEST_TAG = ctxKey("request_id")
)

// Handler ...
type Handler struct {
	logger         *logger.Logger
	pointmanager   *points.Manager
	receiptmanager *receipts.Manager
	router         *chi.Mux
}

// New new handler
func New(lg *logger.Logger, pm *points.Manager, rm *receipts.Manager) (http.Handler, error) {

	router := chi.NewRouter()

	hl := &Handler{
		logger:         lg,
		pointmanager:   pm,
		receiptmanager: rm,
		router:         router,
	}

	lg.Infof(context.Background(), "Application starting")

	//add a request_id to all incoming traffic
	router.Use(hl.addRequestID)

	// CORS
	router.Use(hl.setupCors().Handler)

	// routes
	router.Group(func(router chi.Router) {
		routes := []struct {
			httpVerb    string
			route       string
			handlerFunc func(http.ResponseWriter, *http.Request) error
		}{
			//GET
			{http.MethodGet, "/receipts/{id}/points", hl.HandleGetPoints},
			//POST
			{http.MethodPost, "/receipts/process", hl.HandleProcessReceipt},
		}
		for _, r := range routes {
			switch r.httpVerb {
			case http.MethodGet:
				router.Get(r.route, RoutingMiddleware(r.route, http.HandlerFunc(hl.ErrorCatch(r.handlerFunc))).ServeHTTP)
			case http.MethodPost:
				router.Post(r.route, RoutingMiddleware(r.route, http.HandlerFunc(hl.ErrorCatch(r.handlerFunc))).ServeHTTP)
			}
		}
	})

	return hl.router, nil
}

// ErrorReply ...
type ErrorReply struct {
	Message string
}

// ServeHTTP serves http request
func (hl *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hl.router.ServeHTTP(w, r)
}

// ErrorCatch wrapper for handler functions, catches errors and logs
func (hl *Handler) ErrorCatch(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			var caughtError *errorhandler.Error
			httpcode := http.StatusInternalServerError
			if cerr, ok := err.(*errorhandler.Error); ok {
				caughtError = cerr
				httpcode = caughtError.ErrorHTTPCode
			}
			reply := &ErrorReply{Message: err.Error()}
			render.Status(r, httpcode)
			render.JSON(w, r, reply)
			if caughtError != nil {
				hl.logger.Errorf(r.Context(), "Handler Error catch %+v", caughtError)
			} else {
				hl.logger.Errorf(r.Context(), "Handler Error catch %s", err.Error())
			}
		}
	}
}

// AddRequestID generates a new UUID and appends it to the context as "request_id"
func (hl *Handler) addRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, REQUEST_TAG, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// fastJSON use jsoniter for faster encode
func (hl *Handler) fastJSON(w http.ResponseWriter, r *http.Request, resp interface{}) {
	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buff.Bytes())
}

// RoutingMiddleware doesn't do much right now, but can be used for logging, metric, more
func RoutingMiddleware(route string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		// metrics here
	})
}
